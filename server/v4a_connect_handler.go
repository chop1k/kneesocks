package server

import (
	"net"
	"socks/config"
	"socks/logger"
	"socks/protocol/v4a"
	"time"
)

type V4aConnectHandler interface {
	HandleV4aConnect(address string, client net.Conn)
}

type BaseV4aConnectHandler struct {
	config        config.SocksV4aConfig
	tcpConfig     config.TcpConfig
	streamHandler StreamHandler
	logger        logger.SocksV4aLogger
	protocol      v4a.Protocol
}

func NewBaseV4aConnectHandler(
	config config.SocksV4aConfig,
	tcpConfig config.TcpConfig,
	streamHandler StreamHandler,
	logger logger.SocksV4aLogger,
	protocol v4a.Protocol,
) (BaseV4aConnectHandler, error) {
	return BaseV4aConnectHandler{
		config:        config,
		tcpConfig:     tcpConfig,
		streamHandler: streamHandler,
		logger:        logger,
		protocol:      protocol,
	}, nil
}

func (b BaseV4aConnectHandler) sendFailAndClose(client net.Conn) {
	_ = b.protocol.ResponseWithFail(uint16(b.tcpConfig.GetBindPort()), net.IP{0, 0, 0, 0}, client)
	_ = client.Close()
}

func (b BaseV4aConnectHandler) HandleV4aConnect(address string, client net.Conn) {
	deadline := time.Second * time.Duration(b.config.GetConnectDeadline())

	host, err := net.DialTimeout("tcp", address, deadline)

	if err != nil {
		b.sendFailAndClose(client)

		b.logger.ConnectFailed(client.RemoteAddr().String(), address)

		return
	}

	b.connectSendSuccess(address, host, client)
}

func (b BaseV4aConnectHandler) connectSendSuccess(address string, host net.Conn, client net.Conn) {
	err := b.protocol.ResponseWithSuccess(uint16(b.tcpConfig.GetBindPort()), net.IP{0, 0, 0, 0}, client)

	if err != nil {
		b.sendFailAndClose(client)

		_ = host.Close()

		b.logger.ConnectFailed(client.RemoteAddr().String(), address)

		return
	}

	b.logger.ConnectSuccessful(client.RemoteAddr().String(), address)

	go b.streamHandler.ClientToHost(client, host)
	b.streamHandler.HostToClient(client, host)
}
