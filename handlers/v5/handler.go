package v5

import (
	"fmt"
	"net"
	v52 "socks/config/v5"
	"socks/handlers/v5/helpers"
	v53 "socks/logger/v5"
	v5 "socks/protocol/v5"
)

type Handler interface {
	Handle(request []byte, client net.Conn)
}

type BaseHandler struct {
	parser                v5.Parser
	authenticationHandler AuthenticationHandler
	logger                v53.Logger
	connectHandler        ConnectHandler
	bindHandler           BindHandler
	udpAssociationHandler UdpAssociationHandler
	errorHandler          ErrorHandler
	sender                v5.Sender
	receiver              v5.Receiver
	validator             helpers.Validator
	cleaner               helpers.Cleaner
	replicator            v52.ConfigReplicator
}

func NewBaseHandler(
	parser v5.Parser,
	authenticationHandler AuthenticationHandler,
	logger v53.Logger,
	connectHandler ConnectHandler,
	bindHandler BindHandler,
	udpAssociationHandler UdpAssociationHandler,
	errorHandler ErrorHandler,
	sender v5.Sender,
	receiver v5.Receiver,
	validator helpers.Validator,
	cleaner helpers.Cleaner,
	replicator v52.ConfigReplicator,
) (BaseHandler, error) {
	return BaseHandler{
		parser:                parser,
		authenticationHandler: authenticationHandler,
		logger:                logger,
		connectHandler:        connectHandler,
		bindHandler:           bindHandler,
		udpAssociationHandler: udpAssociationHandler,
		errorHandler:          errorHandler,
		sender:                sender,
		receiver:              receiver,
		validator:             validator,
		cleaner:               cleaner,
		replicator:            replicator,
	}, nil
}

func (b BaseHandler) Handle(request []byte, client net.Conn) {
	configPointer := b.replicator.Copy()

	if configPointer == nil {
		// TODO: do
	}

	config := *configPointer

	methods, err := b.parser.ParseMethods(request)

	if err != nil {
		b.errorHandler.HandleParseMethodsError(config, err, client)

		return
	}

	b.handleAuthentication(config, methods, client)
}

func (b BaseHandler) handleAuthentication(config v52.Config, methods v5.MethodsChunk, client net.Conn) {
	name, err := b.authenticationHandler.HandleAuthentication(config, methods, client)

	if err != nil {
		_ = client.Close()

		b.logger.Auth.Failed(client.RemoteAddr().String())

		return
	}

	b.logger.Auth.Successful(client.RemoteAddr().String(), name)

	b.handleChunk(config, name, client)
}

func (b BaseHandler) handleChunk(config v52.Config, name string, client net.Conn) {
	chunk, err := b.receiver.ReceiveRequest(config, client)

	if err != nil {
		b.errorHandler.HandleReceiveRequestError(config, err, client)

		return
	}

	b.handleCommand(config, name, chunk, client)
}

func (b BaseHandler) handleCommand(config v52.Config, name string, chunk v5.RequestChunk, client net.Conn) {
	var address string

	if chunk.AddressType == 4 {
		address = fmt.Sprintf("[%s]:%d", chunk.Address, chunk.Port)
	} else if chunk.AddressType == 1 || chunk.AddressType == 3 {
		address = fmt.Sprintf("%s:%d", chunk.Address, chunk.Port)
	} else {
		b.errorHandler.HandleInvalidAddressTypeError(config, chunk.AddressType, chunk.Address, client)

		return
	}

	if !b.validator.ValidateRestrictions(config, chunk.CommandCode, name, chunk.AddressType, address, client) {
		return
	}

	if chunk.CommandCode == 1 {
		b.handleConnect(config, name, address, client)
	} else if chunk.CommandCode == 2 {
		b.handleBind(config, name, address, client)
	} else if chunk.CommandCode == 3 {
		b.handleUdpAssociate(config, name, address, client)
	} else {
		b.errorHandler.HandleUnknownCommandError(config, chunk.CommandCode, address, client)

		return
	}

	b.cleaner.Clean(name)
}

func (b BaseHandler) handleConnect(config v52.Config, name string, address string, client net.Conn) {
	b.logger.Connect.Request(client.RemoteAddr().String(), address)

	b.connectHandler.HandleConnect(config, name, address, client)
}

func (b BaseHandler) handleBind(config v52.Config, name string, address string, client net.Conn) {
	b.logger.Bind.Request(client.RemoteAddr().String(), address)

	b.bindHandler.HandleBind(config, name, address, client)
}

func (b BaseHandler) handleUdpAssociate(config v52.Config, name string, _ string, client net.Conn) {
	b.logger.Association.Request(client.RemoteAddr().String())

	b.udpAssociationHandler.HandleUdpAssociation(config, name, client)
}
