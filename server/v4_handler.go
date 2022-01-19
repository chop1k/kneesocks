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
	parser         v4.Parser
	config         config.SocksV4Config
	logger         logger.SocksV4Logger
	connectHandler V4ConnectHandler
	bindHandler    V4BindHandler
	sender         V4Sender
}

func NewBaseV4Handler(
	parser v4.Parser,
	config config.SocksV4Config,
	logger logger.SocksV4Logger,
	connectHandler V4ConnectHandler,
	bindHandler V4BindHandler,
	sender V4Sender,
) (BaseV4Handler, error) {
	return BaseV4Handler{
		parser:         parser,
		config:         config,
		logger:         logger,
		connectHandler: connectHandler,
		bindHandler:    bindHandler,
		sender:         sender,
	}, nil
}

func (b BaseV4Handler) HandleV4(request []byte, client net.Conn) {
	chunk, err := b.parser.ParseRequest(request)

	if err != nil {
		b.sender.SendFailAndClose(client)

		b.logger.ParseError(client.RemoteAddr().String(), err)

		return
	}

	address := fmt.Sprintf("%s:%d", chunk.DestinationIp, chunk.DestinationPort)

	if chunk.CommandCode == 1 {
		b.handleConnect(address, client)
	} else if chunk.CommandCode == 2 {
		b.handleBind(address, client)
	} else {
		b.sender.SendFailAndClose(client)
	}
}

func (b BaseV4Handler) handleConnect(address string, client net.Conn) {
	b.logger.ConnectRequest(client.RemoteAddr().String(), address)

	if !b.config.IsConnectAllowed() {
		b.sender.SendFailAndClose(client)

		b.logger.ConnectNotAllowed(client.RemoteAddr().String(), address)

		return
	}

	b.connectHandler.HandleV4Connect(address, client)
}
func (b BaseV4Handler) handleBind(address string, client net.Conn) {
	b.logger.BindRequest(client.RemoteAddr().String(), address)

	if !b.config.IsBindAllowed() {
		b.sender.SendFailAndClose(client)

		b.logger.BindNotAllowed(client.RemoteAddr().String(), address)

		return
	}

	b.bindHandler.HandleV4Bind(address, client)
}
