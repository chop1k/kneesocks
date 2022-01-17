package server

import (
	"fmt"
	"net"
	"socks/config"
	"socks/logger"
	v4 "socks/protocol/v4"
)

type V4Handler interface {
	HandleV4(request []byte, client net.Conn)
}

type BaseV4Handler struct {
	protocol       v4.Protocol
	parser         v4.Parser
	config         config.SocksV4Config
	tcpConfig      config.TcpConfig
	logger         logger.SocksV4Logger
	connectHandler V4ConnectHandler
	bindHandler    V4BindHandler
}

func NewBaseV4Handler(
	protocol v4.Protocol,
	parser v4.Parser,
	config config.SocksV4Config,
	tcpConfig config.TcpConfig,
	logger logger.SocksV4Logger,
	connectHandler V4ConnectHandler,
	bindHandler V4BindHandler,
) (BaseV4Handler, error) {
	return BaseV4Handler{
		protocol:       protocol,
		parser:         parser,
		config:         config,
		tcpConfig:      tcpConfig,
		logger:         logger,
		connectHandler: connectHandler,
		bindHandler:    bindHandler,
	}, nil
}

func (b BaseV4Handler) HandleV4(request []byte, client net.Conn) {
	chunk, err := b.parser.ParseRequest(request)

	if err != nil {
		b.sendFailAndClose(client)

		return
	}

	address := fmt.Sprintf("%s:%d", chunk.DestinationIp, chunk.DestinationPort)

	if chunk.CommandCode == 1 {
		b.handleConnect(address, client)
	} else if chunk.CommandCode == 2 {
		b.handleBind(address, client)
	} else {
		b.sendFailAndClose(client)
	}
}

func (b BaseV4Handler) sendFailAndClose(client net.Conn) {
	_ = b.protocol.ResponseWithFail(uint16(b.tcpConfig.GetBindPort()), net.IP{0, 0, 0, 0}, client)
	_ = client.Close()
}

func (b BaseV4Handler) handleConnect(address string, client net.Conn) {
	b.logger.ConnectRequest(client.RemoteAddr().String(), address)

	if !b.config.IsConnectAllowed() {
		b.sendFailAndClose(client)

		b.logger.ConnectNotAllowed(client.RemoteAddr().String(), address)

		return
	}

	b.connectHandler.HandleV4Connect(address, client)
}
func (b BaseV4Handler) handleBind(address string, client net.Conn) {
	b.logger.BindRequest(client.RemoteAddr().String(), address)

	if !b.config.IsBindAllowed() {
		b.sendFailAndClose(client)

		b.logger.BindNotAllowed(client.RemoteAddr().String(), address)

		return
	}

	b.bindHandler.HandleV4Bind(address, client)
}
