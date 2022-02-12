package v5

import (
	"fmt"
	"net"
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
	}, nil
}

func (b BaseHandler) Handle(request []byte, client net.Conn) {
	methods, err := b.parser.ParseMethods(request)

	if err != nil {
		b.errorHandler.HandleParseMethodsError(err, client)

		return
	}

	b.handleAuthentication(methods, client)
}

func (b BaseHandler) handleAuthentication(methods v5.MethodsChunk, client net.Conn) {
	name, err := b.authenticationHandler.HandleAuthentication(methods, client)

	if err != nil {
		_ = client.Close()

		b.logger.Auth.Failed(client.RemoteAddr().String())

		return
	}

	b.logger.Auth.Successful(client.RemoteAddr().String(), name)

	b.handleChunk(name, client)
}

func (b BaseHandler) handleChunk(name string, client net.Conn) {
	chunk, err := b.receiver.ReceiveRequest(client)

	if err != nil {
		b.errorHandler.HandleReceiveRequestError(err, client)

		return
	}

	b.handleCommand(name, chunk, client)
}

func (b BaseHandler) handleCommand(name string, chunk v5.RequestChunk, client net.Conn) {
	var address string

	if chunk.AddressType == 4 {
		address = fmt.Sprintf("[%s]:%d", chunk.Address, chunk.Port)
	} else if chunk.AddressType == 1 || chunk.AddressType == 3 {
		address = fmt.Sprintf("%s:%d", chunk.Address, chunk.Port)
	} else {
		b.errorHandler.HandleInvalidAddressTypeError(chunk.AddressType, chunk.Address, client)

		return
	}

	if !b.validator.ValidateRestrictions(chunk.CommandCode, name, chunk.AddressType, address, client) {
		return
	}

	if chunk.CommandCode == 1 {
		b.handleConnect(name, address, client)
	} else if chunk.CommandCode == 2 {
		b.handleBind(name, address, client)
	} else if chunk.CommandCode == 3 {
		b.handleUdpAssociate(name, address, client)
	} else {
		b.errorHandler.HandleUnknownCommandError(chunk.CommandCode, address, client)

		return
	}

	b.cleaner.Clean(name)
}

func (b BaseHandler) handleConnect(name string, address string, client net.Conn) {
	b.logger.Connect.Request(client.RemoteAddr().String(), address)

	b.connectHandler.HandleConnect(name, address, client)
}

func (b BaseHandler) handleBind(name string, address string, client net.Conn) {
	b.logger.Bind.Request(client.RemoteAddr().String(), address)

	b.bindHandler.HandleBind(name, address, client)
}

func (b BaseHandler) handleUdpAssociate(name string, _ string, client net.Conn) {
	b.logger.Association.Request(client.RemoteAddr().String())

	b.udpAssociationHandler.HandleUdpAssociation(name, client)
}
