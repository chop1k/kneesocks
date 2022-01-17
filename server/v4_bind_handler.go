package server

import (
	"net"
	"socks/config"
	"socks/logger"
	v4 "socks/protocol/v4"
	"socks/utils"
	"time"
)

type V4BindHandler interface {
	HandleV4Bind(address string, client net.Conn)
}

type BaseV4BindHandler struct {
	config        config.SocksV4Config
	tcpConfig     config.TcpConfig
	logger        logger.SocksV4Logger
	streamHandler StreamHandler
	protocol      v4.Protocol
	bindManager   BindManager
	utils         utils.AddressUtils
}

func NewBaseV4BindHandler(config config.SocksV4Config, tcpConfig config.TcpConfig, logger logger.SocksV4Logger, streamHandler StreamHandler, protocol v4.Protocol, bindManager BindManager, utils utils.AddressUtils) (BaseV4BindHandler, error) {
	return BaseV4BindHandler{
		config:        config,
		tcpConfig:     tcpConfig,
		logger:        logger,
		streamHandler: streamHandler,
		protocol:      protocol,
		bindManager:   bindManager,
		utils:         utils,
	}, nil
}

func (b BaseV4BindHandler) sendFailAndClose(client net.Conn) {
	_ = b.protocol.ResponseWithFail(uint16(b.tcpConfig.GetBindPort()), net.IP{0, 0, 0, 0}, client)
	_ = client.Close()
}

func (b BaseV4BindHandler) HandleV4Bind(address string, client net.Conn) {
	err := b.bindManager.Bind(address)

	if err != nil {
		b.sendFailAndClose(client)

		b.logger.BindFailed(client.RemoteAddr().String(), address)

		return
	}

	b.logger.BindSuccessful(client.RemoteAddr().String(), address)

	b.bindSendFirstResponse(address, client)
}

func (b BaseV4BindHandler) bindSendFirstResponse(address string, client net.Conn) {
	err := b.protocol.ResponseWithSuccess(0, net.IP{0, 0, 0, 0}, client)

	if err != nil {
		b.sendFailAndClose(client)

		b.logger.BindFailed(client.RemoteAddr().String(), address)

		return
	}

	b.bindWait(address, client)

	b.bindManager.Remove(address)
}

func (b BaseV4BindHandler) bindWait(address string, client net.Conn) {
	deadline := time.Second * time.Duration(b.config.GetBindDeadline())

	host, err := b.bindManager.ReceiveHost(address, deadline)

	if err != nil {
		b.sendFailAndClose(client)

		if err == TimeoutError {
			b.logger.BindTimeout(client.RemoteAddr().String(), address)
		} else {
			b.logger.BindFailed(client.RemoteAddr().String(), address)
		}

		return
	}

	b.bindCheckAddress(address, host, client)
}

func (b BaseV4BindHandler) bindCheckAddress(address string, host net.Conn, client net.Conn) {
	hostAddr, hostPort, parseErr := b.utils.ParseAddress(host.RemoteAddr().String())

	if parseErr != nil {
		b.sendFailAndClose(client)

		_ = host.Close()

		b.logger.BindFailed(client.RemoteAddr().String(), address)

		return
	}

	addrType, determineErr := b.utils.DetermineAddressType(hostAddr)

	if determineErr != nil {
		b.sendFailAndClose(client)

		_ = host.Close()

		b.logger.BindFailed(client.RemoteAddr().String(), address)

		return
	}

	if addrType != 1 {
		b.sendFailAndClose(client)

		_ = host.Close()

		b.logger.BindFailed(client.RemoteAddr().String(), address)

		return
	}

	b.bindSendSecondResponse(address, hostAddr, uint16(hostPort), host, client)
}

func (b BaseV4BindHandler) bindSendSecondResponse(address string, hostAddr string, hostPort uint16, host net.Conn, client net.Conn) {
	ip := net.ParseIP(hostAddr).To4()

	err := b.protocol.ResponseWithSuccess(hostPort, ip, client)

	if err != nil {
		b.sendFailAndClose(client)

		_ = host.Close()

		b.logger.BindFailed(client.RemoteAddr().String(), address)

		return
	}

	err = b.bindManager.SendClient(address, client)

	if err != nil {
		_ = client.Close()
		_ = host.Close()

		b.logger.BindFailed(client.RemoteAddr().String(), address)

		return
	}

	b.logger.Bound(client.RemoteAddr().String(), host.RemoteAddr().String())

	b.streamHandler.ClientToHost(client, host)
}
