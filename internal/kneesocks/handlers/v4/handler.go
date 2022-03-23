package v4

import (
	"fmt"
	"net"
	v42 "socks/internal/kneesocks/config/v4"
	"socks/internal/kneesocks/handlers/v4/helpers"
	v43 "socks/internal/kneesocks/logger/v4"
	v44 "socks/pkg/protocol/v4"
)

type Handler struct {
	parser         v44.Parser
	logger         v43.Logger
	connectHandler *ConnectHandler
	bindHandler    *BindHandler
	sender         v44.Sender
	errorHandler   ErrorHandler
	validator      helpers.Validator
	cleaner        helpers.Cleaner
	replicator     v42.ConfigReplicator
}

func NewHandler(
	parser v44.Parser,
	logger v43.Logger,
	connectHandler *ConnectHandler,
	bindHandler *BindHandler,
	sender v44.Sender,
	errorHandler ErrorHandler,
	validator helpers.Validator,
	cleaner helpers.Cleaner,
	replicator v42.ConfigReplicator,
) (*Handler, error) {
	return &Handler{
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
	config := b.replicator.Copy()

	chunk, err := b.parser.ParseRequest(request)

	if err != nil {
		b.errorHandler.HandleChunkParseError(config, err, client)

		return
	}

	address := fmt.Sprintf("%s:%d", chunk.DestinationIp, chunk.DestinationPort)

	if chunk.CommandCode == 1 && b.connectHandler == nil {
		b.sender.SendFailAndClose(config, client)

		b.logger.Restrictions.NotAllowed(client.RemoteAddr().String(), address)

		return
	}

	if chunk.CommandCode == 2 && b.bindHandler == nil {
		b.sender.SendFailAndClose(config, client)

		b.logger.Restrictions.NotAllowed(client.RemoteAddr().String(), address)

		return
	}

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
