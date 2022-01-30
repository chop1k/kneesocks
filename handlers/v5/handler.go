package v5

import (
	"fmt"
	"net"
	"socks/config"
	"socks/logger"
	v5 "socks/protocol/v5"
)

type Handler interface {
	Handle(request []byte, client net.Conn)
}

type BaseHandler struct {
	protocol              v5.Protocol
	parser                v5.Parser
	config                config.SocksV5Config
	authenticationHandler AuthenticationHandler
	logger                logger.SocksV5Logger
	connectHandler        ConnectHandler
	bindHandler           BindHandler
	udpAssociationHandler UdpAssociationHandler
	sender                Sender
	errorHandler          ErrorHandler
}

func NewBaseHandler(
	protocol v5.Protocol,
	parser v5.Parser,
	config config.SocksV5Config,
	authenticationHandler AuthenticationHandler,
	logger logger.SocksV5Logger,
	connectHandler ConnectHandler,
	bindHandler BindHandler,
	udpAssociationHandler UdpAssociationHandler,
	sender Sender,
	errorHandler ErrorHandler,
) (BaseHandler, error) {
	return BaseHandler{
		protocol:              protocol,
		parser:                parser,
		config:                config,
		authenticationHandler: authenticationHandler,
		logger:                logger,
		connectHandler:        connectHandler,
		bindHandler:           bindHandler,
		udpAssociationHandler: udpAssociationHandler,
		sender:                sender,
		errorHandler:          errorHandler,
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

		b.logger.AuthenticationFailed(client.RemoteAddr().String())

		return
	}

	b.logger.AuthenticationSuccessful(client.RemoteAddr().String(), name)

	b.handleChunk(name, client)
}

func (b BaseHandler) handleChunk(name string, client net.Conn) {
	chunk, err := b.protocol.ReceiveRequest(client)

	if err != nil {
		b.errorHandler.HandleReceiveRequestError(err, client)

		return
	}

	b.handleAddress(name, chunk, client)
}

func (b BaseHandler) handleAddress(name string, chunk v5.RequestChunk, client net.Conn) {
	if chunk.AddressType == 1 {
		b.handleIPv4(name, chunk, client)
	} else if chunk.AddressType == 3 {
		b.handleDomain(name, chunk, client)
	} else if chunk.AddressType == 4 {
		b.handleIPv6(name, chunk, client)
	} else {
		b.errorHandler.HandleInvalidAddressTypeError(chunk.AddressType, chunk.Address, client)

		return
	}
}

func (b BaseHandler) handleIPv4(name string, chunk v5.RequestChunk, client net.Conn) {
	address := fmt.Sprintf("%s:%d", chunk.Address, chunk.Port)

	if !b.config.IsIPv4Allowed() {
		b.errorHandler.HandleIPv4AddressNotAllowed(address, client)

		return
	}

	b.handleCommand(name, chunk.CommandCode, address, client)
}

func (b BaseHandler) handleDomain(name string, chunk v5.RequestChunk, client net.Conn) {
	address := fmt.Sprintf("%s:%d", chunk.Address, chunk.Port)

	if !b.config.IsDomainAllowed() {
		b.errorHandler.HandleDomainAddressNotAllowed(address, client)

		return
	}

	b.handleCommand(name, chunk.CommandCode, address, client)
}

func (b BaseHandler) handleIPv6(name string, chunk v5.RequestChunk, client net.Conn) {
	address := fmt.Sprintf("[%s]:%d", chunk.Address, chunk.Port)

	if !b.config.IsIPv6Allowed() {
		b.errorHandler.HandleIPv6AddressNotAllowed(address, client)

		return
	}

	b.handleCommand(name, chunk.CommandCode, address, client)
}

func (b BaseHandler) handleCommand(name string, command byte, address string, client net.Conn) {
	if command == 1 {
		b.handleConnect(name, address, client)
	} else if command == 2 {
		b.handleBind(name, address, client)
	} else if command == 3 {
		b.handleUdpAssociate(name, client)
	} else {
		b.errorHandler.HandleUnknownCommandError(command, address, client)

		return
	}
}

func (b BaseHandler) handleConnect(name string, address string, client net.Conn) {
	b.logger.ConnectRequest(client.RemoteAddr().String(), address)

	if !b.config.IsConnectAllowed() {
		b.sender.SendConnectionNotAllowedAndClose(client)

		b.logger.ConnectNotAllowed(client.RemoteAddr().String(), address)

		return
	}

	b.connectHandler.HandleConnect(name, address, client)
}

func (b BaseHandler) handleBind(name string, address string, client net.Conn) {
	b.logger.BindRequest(client.RemoteAddr().String(), address)

	if !b.config.IsBindAllowed() {
		b.sender.SendConnectionNotAllowedAndClose(client)

		b.logger.BindNotAllowed(client.RemoteAddr().String(), address)

		return
	}

	b.bindHandler.HandleBind(name, address, client)
}

func (b BaseHandler) handleUdpAssociate(name string, client net.Conn) {
	b.logger.UdpAssociationRequest(client.RemoteAddr().String())

	if !b.config.IsUdpAssociationAllowed() {
		b.sender.SendConnectionNotAllowedAndClose(client)

		b.logger.UdpAssociationNotAllowed(client.RemoteAddr().String())

		return
	}

	b.udpAssociationHandler.HandleUdpAssociation(name, client)
}
