package v4a

import (
	"fmt"
	"net"
	v4a2 "socks/internal/kneesocks/config/v4a"
	"socks/internal/kneesocks/handlers/v4a/helpers"
	v4a3 "socks/internal/kneesocks/logger/v4a"
	v4a4 "socks/pkg/protocol/v4a"
)

type Handler struct {
	parser         v4a4.Parser
	logger         v4a3.Logger
	connectHandler ConnectHandler
	bindHandler    BindHandler
	sender         v4a4.Sender
	errorHandler   ErrorHandler
	validator      helpers.Validator
	cleaner        helpers.Cleaner
	replicator     v4a2.ConfigReplicator
}

func NewHandler(
	parser v4a4.Parser,
	logger v4a3.Logger,
	connectHandler ConnectHandler,
	bindHandler BindHandler,
	sender v4a4.Sender,
	errorHandler ErrorHandler,
	validator helpers.Validator,
	cleaner helpers.Cleaner,
	replicator v4a2.ConfigReplicator,
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
		_ = client.Close()

		return
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

func (b Handler) handleConnect(config v4a2.Config, address string, client net.Conn) {
	b.logger.Connect.Request(client.RemoteAddr().String(), address)

	b.connectHandler.HandleConnect(config, address, client)
}

func (b Handler) handleBind(config v4a2.Config, address string, client net.Conn) {
	b.logger.Bind.Request(client.RemoteAddr().String(), address)

	b.bindHandler.HandleBind(config, address, client)
}
