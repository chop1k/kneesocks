package server

import (
	"fmt"
	"net"
	"socks/config"
	"socks/logger"
	v4 "socks/protocol/v4"
	"socks/utils"
	"time"
)

type V4Handler interface {
	HandleV4(request []byte, client net.Conn)
}

type BaseV4Handler struct {
	protocol      v4.Protocol
	parser        v4.Parser
	bindManager   BindManager
	config        config.SocksV4Config
	tcpConfig     config.TcpConfig
	streamHandler StreamHandler
	utils         utils.AddressUtils
	logger        logger.SocksV4Logger
}

func NewBaseV4Handler(
	protocol v4.Protocol,
	parser v4.Parser,
	bindManager BindManager,
	config config.SocksV4Config,
	tcpConfig config.TcpConfig,
	streamHandler StreamHandler,
	utils utils.AddressUtils,
	logger logger.SocksV4Logger,
) (BaseV4Handler, error) {
	return BaseV4Handler{
		protocol:      protocol,
		parser:        parser,
		bindManager:   bindManager,
		config:        config,
		tcpConfig:     tcpConfig,
		streamHandler: streamHandler,
		utils:         utils,
		logger:        logger,
	}, nil
}

func (b BaseV4Handler) HandleV4(request []byte, client net.Conn) {
	chunk, err := b.parser.ParseRequest(request)

	if err != nil {
		b.sendFailAndClose(client)

		return
	}

	if chunk.CommandCode == 1 {
		go b.logger.ConnectRequest(client.RemoteAddr().String(), chunk)

		if !b.config.IsConnectAllowed() {
			b.sendFailAndClose(client)

			go b.logger.ConnectNotAllowed(client.RemoteAddr().String(), chunk)

			return
		}

		b.handleConnect(chunk, client)
	} else if chunk.CommandCode == 2 {
		go b.logger.BindRequest(client.RemoteAddr().String(), chunk)

		if !b.config.IsBindAllowed() {
			b.sendFailAndClose(client)

			go b.logger.BindNotAllowed(client.RemoteAddr().String(), chunk)

			return
		}

		b.handleBind(chunk, client)
	} else {
		b.sendFailAndClose(client)
	}
}

func (b BaseV4Handler) sendFailAndClose(client net.Conn) {
	_ = b.protocol.ResponseWithCode(91, uint16(b.tcpConfig.GetBindPort()), net.IP{0, 0, 0, 0}, client)
	_ = client.Close()
}

func (b BaseV4Handler) handleConnect(chunk v4.RequestChunk, client net.Conn) {
	addr := fmt.Sprintf("%s:%d", chunk.DestinationIp, chunk.DestinationPort)

	deadline := time.Second * time.Duration(b.config.GetConnectDeadline())

	host, err := net.DialTimeout("tcp4", addr, deadline)

	if err != nil {
		b.sendFailAndClose(client)

		go b.logger.ConnectFailed(client.RemoteAddr().String(), chunk)

		return
	}

	b.connectSendSuccess(chunk, host, client)
}

func (b BaseV4Handler) connectSendSuccess(chunk v4.RequestChunk, host net.Conn, client net.Conn) {
	err := b.protocol.ResponseWithCode(90, uint16(b.tcpConfig.GetBindPort()), net.IP{0, 0, 0, 0}, client)

	if err != nil {
		b.sendFailAndClose(client)

		_ = host.Close()

		go b.logger.ConnectFailed(client.RemoteAddr().String(), chunk)

		return
	}

	go b.logger.ConnectSuccessful(client.RemoteAddr().String(), chunk)

	go b.streamHandler.ClientToHost(client, host)
	b.streamHandler.HostToClient(client, host)
}

func (b BaseV4Handler) handleBind(chunk v4.RequestChunk, client net.Conn) {
	address := fmt.Sprintf("%s:%d", chunk.DestinationIp, chunk.DestinationPort)

	err := b.bindManager.Bind(address)

	if err != nil {
		b.sendFailAndClose(client)

		go b.logger.BindFailed(client.RemoteAddr().String(), chunk)

		return
	}

	go b.logger.BindSuccessful(client.RemoteAddr().String(), chunk)

	b.bindSendFirstResponse(chunk, address, client)
}

func (b BaseV4Handler) bindSendFirstResponse(chunk v4.RequestChunk, address string, client net.Conn) {
	err := b.protocol.ResponseWithCode(90, 0, net.IP{0, 0, 0, 0}, client)

	if err != nil {
		b.sendFailAndClose(client)

		go b.logger.BindFailed(client.RemoteAddr().String(), chunk)

		return
	}

	b.bindWait(chunk, address, client)

	b.bindManager.Remove(address)
}

func (b BaseV4Handler) bindWait(chunk v4.RequestChunk, address string, client net.Conn) {
	deadline := time.Second * time.Duration(b.config.GetBindDeadline())

	host, err := b.bindManager.ReceiveHost(address, deadline)

	if err != nil {
		b.sendFailAndClose(client)

		go b.logger.BindFailed(client.RemoteAddr().String(), chunk)

		return
	}

	b.bindCheckAddress(chunk, address, host, client)
}

func (b BaseV4Handler) bindCheckAddress(chunk v4.RequestChunk, address string, host net.Conn, client net.Conn) {
	hostAddr, hostPort, parseErr := b.utils.ParseAddress(host.RemoteAddr().String())

	if parseErr != nil {
		b.sendFailAndClose(client)

		_ = host.Close()

		go b.logger.BindFailed(client.RemoteAddr().String(), chunk)

		return
	}

	addrType, determineErr := b.utils.DetermineAddressType(hostAddr)

	if determineErr != nil {
		b.sendFailAndClose(client)

		_ = host.Close()

		go b.logger.BindFailed(client.RemoteAddr().String(), chunk)

		return
	}

	if addrType != 1 {
		b.sendFailAndClose(client)

		_ = host.Close()

		go b.logger.BindFailed(client.RemoteAddr().String(), chunk)

		return
	}

	b.bindSendSecondResponse(chunk, address, hostAddr, uint16(hostPort), host, client)
}

func (b BaseV4Handler) bindSendSecondResponse(chunk v4.RequestChunk, address string, hostAddr string, hostPort uint16, host net.Conn, client net.Conn) {
	ip := net.ParseIP(hostAddr).To4()

	err := b.protocol.ResponseWithCode(90, hostPort, ip, client)

	if err != nil {
		b.sendFailAndClose(client)

		_ = host.Close()

		go b.logger.BindFailed(client.RemoteAddr().String(), chunk)

		return
	}

	err = b.bindManager.SendClient(address, client)

	if err != nil {
		_ = client.Close()
		_ = host.Close()

		go b.logger.BindFailed(client.RemoteAddr().String(), chunk)

		return
	}

	go b.logger.Bound(client.RemoteAddr().String(), host.RemoteAddr().String(), chunk)

	b.streamHandler.ClientToHost(client, host)
}
