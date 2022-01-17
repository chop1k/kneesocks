package server

import (
	"fmt"
	"net"
	"socks/config"
	"socks/logger"
	"socks/protocol/v4a"
)

type V4aHandler interface {
	HandleV4a(request []byte, client net.Conn)
}

type BaseV4aHandler struct {
	protocol       v4a.Protocol
	parser         v4a.Parser
	config         config.SocksV4aConfig
	logger         logger.SocksV4aLogger
	tcpConfig      config.TcpConfig
	connectHandler V4aConnectHandler
	bindHandler    V4aBindHandler
}

func NewBaseV4aHandler(
	protocol v4a.Protocol,
	parser v4a.Parser,
	config config.SocksV4aConfig,
	logger logger.SocksV4aLogger,
	tcpConfig config.TcpConfig,
	connectHandler V4aConnectHandler,
	bindHandler V4aBindHandler,
) (BaseV4aHandler, error) {
	return BaseV4aHandler{
		protocol:       protocol,
		parser:         parser,
		config:         config,
		logger:         logger,
		tcpConfig:      tcpConfig,
		connectHandler: connectHandler,
		bindHandler:    bindHandler,
	}, nil
}

func (b BaseV4aHandler) HandleV4a(request []byte, client net.Conn) {
	chunk, err := b.parser.ParseRequest(request)

	if err != nil {
		b.sendFailAndClose(client)

		return
	}

	address := fmt.Sprintf("%s:%d", chunk.Domain, chunk.DestinationPort)

	if chunk.CommandCode == 1 {
		b.handleConnect(address, client)
	} else if chunk.CommandCode == 2 {
		b.handleBind(address, client)
	} else {
		b.sendFailAndClose(client)
	}
}

func (b BaseV4aHandler) sendFailAndClose(client net.Conn) {
	_ = b.protocol.ResponseWithFail(uint16(b.tcpConfig.GetBindPort()), net.IP{0, 0, 0, 0}, client)
	_ = client.Close()
}

func (b BaseV4aHandler) handleConnect(address string, client net.Conn) {
	b.logger.ConnectRequest(client.RemoteAddr().String(), address)

	if !b.config.IsConnectAllowed() {
		b.sendFailAndClose(client)

		b.logger.ConnectNotAllowed(client.RemoteAddr().String(), address)

		return
	}

	b.connectHandler.HandleV4aConnect(address, client)
}

func (b BaseV4aHandler) handleBind(address string, client net.Conn) {
	b.logger.BindRequest(client.RemoteAddr().String(), address)

	if !b.config.IsBindAllowed() {
		b.sendFailAndClose(client)

		b.logger.BindNotAllowed(client.RemoteAddr().String(), address)

		return
	}

	b.bindHandler.HandleV4aBind(address, client)
}
