package v4

import (
	"fmt"
	"net"
	v42 "socks/config/v4"
	v43 "socks/logger/v4"
	v4 "socks/protocol/v4"
)

type Handler interface {
	Handle(request []byte, client net.Conn)
}

type BaseHandler struct {
	parser         v4.Parser
	config         v42.Config
	logger         v43.Logger
	connectHandler ConnectHandler
	bindHandler    BindHandler
	sender         Sender
	errorHandler   ErrorHandler
}

func NewBaseHandler(
	parser v4.Parser,
	config v42.Config,
	logger v43.Logger,
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

	address := fmt.Sprintf("%s:%d", chunk.DestinationIp, chunk.DestinationPort)

	if chunk.CommandCode == 1 {
		b.handleConnect(address, client)
	} else if chunk.CommandCode == 2 {
		b.handleBind(address, client)
	} else {
		b.sender.SendFailAndClose(client)
	}
}

func (b BaseHandler) handleConnect(address string, client net.Conn) {
	b.logger.Connect.Request(client.RemoteAddr().String(), address)

	if !b.config.IsConnectAllowed() {
		b.sender.SendFailAndClose(client)

		b.logger.Restrictions.NotAllowed(client.RemoteAddr().String(), address)

		return
	}

	b.connectHandler.HandleConnect(address, client)
}
func (b BaseHandler) handleBind(address string, client net.Conn) {
	b.logger.Bind.Request(client.RemoteAddr().String(), address)

	if !b.config.IsBindAllowed() {
		b.sender.SendFailAndClose(client)

		b.logger.Restrictions.NotAllowed(client.RemoteAddr().String(), address)

		return
	}

	b.bindHandler.HandleBind(address, client)
}
