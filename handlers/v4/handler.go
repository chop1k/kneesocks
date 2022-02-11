package v4

import (
	"fmt"
	"net"
	"socks/handlers/v4/helpers"
	v43 "socks/logger/v4"
	v4 "socks/protocol/v4"
)

type Handler interface {
	Handle(request []byte, client net.Conn)
}

type BaseHandler struct {
	parser         v4.Parser
	logger         v43.Logger
	connectHandler ConnectHandler
	bindHandler    BindHandler
	sender         v4.Sender
	errorHandler   ErrorHandler
	validator      helpers.Validator
	cleaner        helpers.Cleaner
}

func NewBaseHandler(
	parser v4.Parser,
	logger v43.Logger,
	connectHandler ConnectHandler,
	bindHandler BindHandler,
	sender v4.Sender,
	errorHandler ErrorHandler,
	validator helpers.Validator,
	cleaner helpers.Cleaner,
) (BaseHandler, error) {
	return BaseHandler{
		parser:         parser,
		logger:         logger,
		connectHandler: connectHandler,
		bindHandler:    bindHandler,
		sender:         sender,
		errorHandler:   errorHandler,
		validator:      validator,
		cleaner:        cleaner,
	}, nil
}

func (b BaseHandler) Handle(request []byte, client net.Conn) {
	chunk, err := b.parser.ParseRequest(request)

	if err != nil {
		b.errorHandler.HandleChunkParseError(err, client)

		return
	}

	address := fmt.Sprintf("%s:%d", chunk.DestinationIp, chunk.DestinationPort)

	if !b.validator.ValidateRestrictions(chunk.CommandCode, address, client) {
		return
	}

	if chunk.CommandCode == 1 {
		b.handleConnect(address, client)
	} else if chunk.CommandCode == 2 {
		b.handleBind(address, client)
	} else {
		b.sender.SendFailAndClose(client)

		return
	}

	err = b.cleaner.Clean()

	if err != nil {
		panic(err)
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
