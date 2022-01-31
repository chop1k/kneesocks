package v4a

import (
	"fmt"
	"net"
	v4a2 "socks/config/v4a"
	v4a3 "socks/logger/v4a"
	"socks/protocol/v4a"
)

type Handler interface {
	Handle(request []byte, client net.Conn)
}

type BaseHandler struct {
	parser         v4a.Parser
	config         v4a2.Config
	logger         v4a3.Logger
	connectHandler ConnectHandler
	bindHandler    BindHandler
	sender         Sender
	errorHandler   ErrorHandler
}

func NewBaseHandler(
	parser v4a.Parser,
	config v4a2.Config,
	logger v4a3.Logger,
	connectHandler ConnectHandler,
	bindHandler BindHandler,
	sender Sender,
	errorHandler ErrorHandler,
) (BaseHandler, error) {
	return BaseHandler{
		parser:         parser,
		config:         config,
		logger:         logger,
		connectHandler: connectHandler,
		bindHandler:    bindHandler,
		sender:         sender,
		errorHandler:   errorHandler,
	}, nil
}

func (b BaseHandler) Handle(request []byte, client net.Conn) {
	chunk, err := b.parser.ParseRequest(request)

	if err != nil {
		b.errorHandler.HandleChunkParseError(err, client)

		return
	}

	address := fmt.Sprintf("%s:%d", chunk.Domain, chunk.DestinationPort)

	if chunk.CommandCode == 1 {
		b.handleConnect(address, client)
	} else if chunk.CommandCode == 2 {
		b.handleBind(address, client)
	} else {
		b.sender.SendFailAndClose(client)
	}
}

func (b BaseHandler) handleConnect(address string, client net.Conn) {
	b.logger.Connect.ConnectRequest(client.RemoteAddr().String(), address)

	if !b.config.IsConnectAllowed() {
		b.sender.SendFailAndClose(client)

		b.logger.Connect.ConnectNotAllowed(client.RemoteAddr().String(), address)

		return
	}

	b.connectHandler.HandleConnect(address, client)
}

func (b BaseHandler) handleBind(address string, client net.Conn) {
	b.logger.Bind.BindRequest(client.RemoteAddr().String(), address)

	if !b.config.IsBindAllowed() {
		b.sender.SendFailAndClose(client)

		b.logger.Bind.BindNotAllowed(client.RemoteAddr().String(), address)

		return
	}

	b.bindHandler.HandleBind(address, client)
}
