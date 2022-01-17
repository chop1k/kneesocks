package server

import (
	"net"
	"socks/config"
	"socks/logger"
	"socks/protocol/v4a"
	"socks/utils"
	"time"
)

type V4aBindHandler interface {
	HandleV4aBind(address string, client net.Conn)
}

type BaseV4aBindHandler struct {
	config        config.SocksV4aConfig
	tcpConfig     config.TcpConfig
	logger        logger.SocksV4aLogger
	streamHandler StreamHandler
	protocol      v4a.Protocol
	bindManager   BindManager
	utils         utils.AddressUtils
}

func NewBaseV4aBindHandler(config config.SocksV4Config, tcpConfig config.TcpConfig, logger logger.SocksV4aLogger, streamHandler StreamHandler, protocol v4a.Protocol, bindManager BindManager, utils utils.AddressUtils) (BaseV4aBindHandler, error) {
	return BaseV4aBindHandler{
		config:        config,
		tcpConfig:     tcpConfig,
		logger:        logger,
		streamHandler: streamHandler,
		protocol:      protocol,
		bindManager:   bindManager,
		utils:         utils,
	}, nil
}

func (b BaseV4aBindHandler) sendFailAndClose(client net.Conn) {
	_ = b.protocol.ResponseWithFail(uint16(b.tcpConfig.GetBindPort()), net.IP{0, 0, 0, 0}, client)
	_ = client.Close()
}

func (b BaseV4aBindHandler) HandleV4aBind(address string, client net.Conn) {
	err := b.bindManager.Bind(address)

	if err != nil {
		b.sendFailAndClose(client)

		b.logger.BindFailed(client.RemoteAddr().String(), address)

		return
	}

	b.logger.BindSuccessful(client.RemoteAddr().String(), address)

	b.bindSendFirstResponse(address, client)
}

func (b BaseV4aBindHandler) bindSendFirstResponse(address string, client net.Conn) {
	err := b.protocol.ResponseWithSuccess(0, net.IP{0, 0, 0, 0}, client)

	if err != nil {
		b.sendFailAndClose(client)

		b.logger.BindFailed(client.RemoteAddr().String(), address)

		return
	}

	b.bindWait(address, client)

	b.bindManager.Remove(address)
}

func (b BaseV4aBindHandler) bindWait(address string, client net.Conn) {
	deadline := time.Second * time.Duration(b.config.GetBindDeadline())

	host, err := b.bindManager.ReceiveHost(address, deadline)

	if err != nil {
		b.sendFailAndClose(client)

		b.logger.BindFailed(client.RemoteAddr().String(), address)

		return
	}

	b.bindCheckAddress(address, host, client)
}

func (b BaseV4aBindHandler) bindCheckAddress(address string, host net.Conn, client net.Conn) {
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

func (b BaseV4aBindHandler) bindSendSecondResponse(address string, hostAddr string, hostPort uint16, host net.Conn, client net.Conn) {
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

	b.logger.BindSuccessful(client.RemoteAddr().String(), address)

	b.streamHandler.ClientToHost(host, client)
}
