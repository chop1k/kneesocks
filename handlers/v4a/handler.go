package v4a

import (
	"fmt"
	"net"
	v4a2 "socks/config/v4a"
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
	sender         v4a.Sender
	errorHandler   ErrorHandler
	validator      helpers.Validator
	cleaner        helpers.Cleaner
	replicator     v4a2.ConfigReplicator
}

func NewBaseHandler(
	parser v4a.Parser,
	logger v4a3.Logger,
	connectHandler ConnectHandler,
	bindHandler BindHandler,
	sender v4a.Sender,
	errorHandler ErrorHandler,
	validator helpers.Validator,
	cleaner helpers.Cleaner,
	replicator v4a2.ConfigReplicator,
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
		replicator:     replicator,
	}, nil
}

func (b BaseHandler) Handle(request []byte, client net.Conn) {
	configPointer := b.replicator.Copy()

	if configPointer == nil {
		// TODO: do
	}

	config := *configPointer

	chunk, err := b.parser.ParseRequest(request)

	if err != nil {
		b.errorHandler.HandleChunkParseError(config, err, client)

		return
	}

	address := fmt.Sprintf("%s:%d", chunk.Domain, chunk.DestinationPort)

	if !b.validator.ValidateRestrictions(config, chunk.CommandCode, address, client) {
		return
	}

	if chunk.CommandCode == 1 {
		b.handleConnect(config, address, client)
	} else if chunk.CommandCode == 2 {
		b.handleBind(config, address, client)
	} else {
		b.sender.SendFailAndClose(config, client)

		return
	}

	b.cleaner.Clean()
}

func (b BaseHandler) handleConnect(config v4a2.Config, address string, client net.Conn) {
	b.logger.Connect.Request(client.RemoteAddr().String(), address)

	b.connectHandler.HandleConnect(config, address, client)
}

func (b BaseHandler) handleBind(config v4a2.Config, address string, client net.Conn) {
	b.logger.Bind.Request(client.RemoteAddr().String(), address)

	b.bindHandler.HandleBind(config, address, client)
}
