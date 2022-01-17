package server

import (
	"net"
	"socks/config"
	"socks/logger"
	v4 "socks/protocol/v4"
	"time"
)

type V4ConnectHandler interface {
	HandleV4Connect(address string, client net.Conn)
}

type BaseV4ConnectHandler struct {
	config        config.SocksV4Config
	tcpConfig     config.TcpConfig
	streamHandler StreamHandler
	logger        logger.SocksV4Logger
	protocol      v4.Protocol
}

func NewBaseV4ConnectHandler(
	config config.SocksV4Config,
	tcpConfig config.TcpConfig,
	streamHandler StreamHandler,
	logger logger.SocksV4Logger,
	protocol v4.Protocol,
) (BaseV4ConnectHandler, error) {
	return BaseV4ConnectHandler{
		config:        config,
		tcpConfig:     tcpConfig,
		streamHandler: streamHandler,
		logger:        logger,
		protocol:      protocol,
	}, nil
}

func (b BaseV4ConnectHandler) sendFailAndClose(client net.Conn) {
	_ = b.protocol.ResponseWithFail(uint16(b.tcpConfig.GetBindPort()), net.IP{0, 0, 0, 0}, client)
	_ = client.Close()
}

func (b BaseV4ConnectHandler) HandleV4Connect(address string, client net.Conn) {
	deadline := time.Second * time.Duration(b.config.GetConnectDeadline())

	host, err := net.DialTimeout("tcp4", address, deadline)

	if err != nil {
		b.sendFailAndClose(client)

		b.logger.ConnectFailed(client.RemoteAddr().String(), address)

		return
	}

	b.connectSendSuccess(address, host, client)
}

func (b BaseV4ConnectHandler) connectSendSuccess(address string, host net.Conn, client net.Conn) {
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
