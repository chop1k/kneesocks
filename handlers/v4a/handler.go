package v4a

import (
	"fmt"
	"net"
	"socks/handlers/v4a/helpers"
	v4a3 "socks/logger/v4a"
	"socks/protocol/v4a"
)

type Handler interface {
	Handle(request []byte, client net.Conn)
}

type BaseHandler struct {
	parser         v4a.Parser
	logger         v4a3.Logger
	connectHandler ConnectHandler
	bindHandler    BindHandler
	sender         helpers.Sender
	errorHandler   ErrorHandler
	validator      helpers.Validator
}

func NewBaseHandler(
	parser v4a.Parser,
	logger v4a3.Logger,
	connectHandler ConnectHandler,
	bindHandler BindHandler,
	sender helpers.Sender,
	errorHandler ErrorHandler,
	validator helpers.Validator,
) (BaseHandler, error) {
	return BaseHandler{
		parser:         parser,
		logger:         logger,
		connectHandler: connectHandler,
		bindHandler:    bindHandler,
		sender:         sender,
		errorHandler:   errorHandler,
		validator:      validator,
	}, nil
}

func (b BaseHandler) Handle(request []byte, client net.Conn) {
	chunk, err := b.parser.ParseRequest(request)

	if err != nil {
		b.errorHandler.HandleChunkParseError(err, client)

		return
	}

	address := fmt.Sprintf("%s:%d", chunk.Domain, chunk.DestinationPort)

	if !b.validator.ValidateRestrictions(chunk.CommandCode, address, client) {
		return
	}

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

	b.connectHandler.HandleConnect(address, client)
}

func (b BaseHandler) handleBind(address string, client net.Conn) {
	b.logger.Bind.Request(client.RemoteAddr().String(), address)

	b.bindHandler.HandleBind(address, client)
}
