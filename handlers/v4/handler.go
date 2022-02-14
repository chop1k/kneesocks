package v4

import (
	"fmt"
	"net"
	v42 "socks/config/v4"
	"socks/handlers/v4/helpers"
	v43 "socks/logger/v4"
	v4 "socks/protocol/v4"
)

type Handler struct {
	parser         v4.Parser
	logger         v43.Logger
	connectHandler ConnectHandler
	bindHandler    BindHandler
	sender         v4.Sender
	errorHandler   ErrorHandler
	validator      helpers.Validator
	cleaner        helpers.Cleaner
	replicator     v42.ConfigReplicator
}

func NewHandler(
	parser v4.Parser,
	logger v43.Logger,
	connectHandler ConnectHandler,
	bindHandler BindHandler,
	sender v4.Sender,
	errorHandler ErrorHandler,
	validator helpers.Validator,
	cleaner helpers.Cleaner,
	replicator v42.ConfigReplicator,
) (Handler, error) {
	return Handler{
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

func (b Handler) Handle(request []byte, client net.Conn) {
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

	address := fmt.Sprintf("%s:%d", chunk.DestinationIp, chunk.DestinationPort)

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

func (b Handler) handleConnect(config v42.Config, address string, client net.Conn) {
	b.logger.Connect.Request(client.RemoteAddr().String(), address)

	b.connectHandler.HandleConnect(config, address, client)
}

func (b Handler) handleBind(config v42.Config, address string, client net.Conn) {
	b.logger.Bind.Request(client.RemoteAddr().String(), address)

	b.bindHandler.HandleBind(config, address, client)
}
