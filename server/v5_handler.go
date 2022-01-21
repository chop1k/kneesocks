package server

import (
	"fmt"
	"net"
	"socks/config"
	"socks/logger"
	v5 "socks/protocol/v5"
)

type V5Handler interface {
	HandleV5(request []byte, client net.Conn)
}

type BaseV5Handler struct {
	protocol              v5.Protocol
	parser                v5.Parser
	config                config.SocksV5Config
	authenticationHandler V5AuthenticationHandler
	logger                logger.SocksV5Logger
	connectHandler        V5ConnectHandler
	bindHandler           V5BindHandler
	udpAssociationHandler V5UdpAssociationHandler
	sender                V5Sender
	errorHandler          V5ErrorHandler
}

func NewBaseV5Handler(
	protocol v5.Protocol,
	parser v5.Parser,
	config config.SocksV5Config,
	authenticationHandler V5AuthenticationHandler,
	logger logger.SocksV5Logger,
	connectHandler V5ConnectHandler,
	bindHandler V5BindHandler,
	udpAssociationHandler V5UdpAssociationHandler,
	sender V5Sender,
	errorHandler V5ErrorHandler,
) (BaseV5Handler, error) {
	return BaseV5Handler{
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

func (b BaseV5Handler) HandleV5(request []byte, client net.Conn) {
	methods, err := b.parser.ParseMethods(request)

	if err != nil {
		b.errorHandler.HandleV5ParseMethodsError(err, client)

		return
	}

	b.handleAuthentication(methods, client)
}

func (b BaseV5Handler) handleAuthentication(methods v5.MethodsChunk, client net.Conn) {
	name, err := b.authenticationHandler.HandleAuthentication(methods, client)

	if err != nil {
		_ = client.Close()

		b.logger.AuthenticationFailed(client.RemoteAddr().String())

		return
	}

	b.logger.AuthenticationSuccessful(client.RemoteAddr().String(), name)

	b.handleChunk(client)
}

func (b BaseV5Handler) handleChunk(client net.Conn) {
	chunk, err := b.protocol.ReceiveRequest(client)

	if err != nil {
		b.errorHandler.HandleV5ReceiveRequestError(err, client)

		return
	}

	b.handleAddress(chunk, client)
}

func (b BaseV5Handler) handleAddress(chunk v5.RequestChunk, client net.Conn) {
	if chunk.AddressType == 1 {
		b.handleIPv4(chunk, client)
	} else if chunk.AddressType == 3 {
		b.handleDomain(chunk, client)
	} else if chunk.AddressType == 4 {
		b.handleIPv6(chunk, client)
	} else {
		b.errorHandler.HandleV5InvalidAddressTypeError(chunk.AddressType, chunk.Address, client)

		return
	}
}

func (b BaseV5Handler) handleIPv4(chunk v5.RequestChunk, client net.Conn) {
	address := fmt.Sprintf("%s:%d", chunk.Address, chunk.Port)

	if !b.config.IsIPv4Allowed() {
		b.errorHandler.HandleV5IPv4AddressNotAllowed(address, client)

		return
	}

	b.handleCommand(chunk.CommandCode, address, client)
}

func (b BaseV5Handler) handleDomain(chunk v5.RequestChunk, client net.Conn) {
	address := fmt.Sprintf("%s:%d", chunk.Address, chunk.Port)

	if !b.config.IsDomainAllowed() {
		b.errorHandler.HandleV5DomainAddressNotAllowed(address, client)

		return
	}

	b.handleCommand(chunk.CommandCode, address, client)
}

func (b BaseV5Handler) handleIPv6(chunk v5.RequestChunk, client net.Conn) {
	address := fmt.Sprintf("[%s]:%d", chunk.Address, chunk.Port)

	if !b.config.IsIPv6Allowed() {
		b.errorHandler.HandleV5IPv6AddressNotAllowed(address, client)

		return
	}

	b.handleCommand(chunk.CommandCode, address, client)
}

func (b BaseV5Handler) handleCommand(command byte, address string, client net.Conn) {
	if command == 1 {
		b.handleConnect(address, client)
	} else if command == 2 {
		b.handleBind(address, client)
	} else if command == 3 {
		b.handleUdpAssociate(client)
	} else {
		b.errorHandler.HandleV5UnknownCommandError(command, address, client)

		return
	}
}

func (b BaseV5Handler) handleConnect(address string, client net.Conn) {
	b.logger.ConnectRequest(client.RemoteAddr().String(), address)

	if !b.config.IsConnectAllowed() {
		b.sender.SendConnectionNotAllowedAndClose(client)

		b.logger.ConnectNotAllowed(client.RemoteAddr().String(), address)

		return
	}

	b.connectHandler.HandleV5Connect(address, client)
}

func (b BaseV5Handler) handleBind(address string, client net.Conn) {
	b.logger.BindRequest(client.RemoteAddr().String(), address)

	if !b.config.IsBindAllowed() {
		b.sender.SendConnectionNotAllowedAndClose(client)

		b.logger.BindNotAllowed(client.RemoteAddr().String(), address)

		return
	}

	b.bindHandler.HandleV5Bind(address, client)
}

func (b BaseV5Handler) handleUdpAssociate(client net.Conn) {
	b.logger.UdpAssociationRequest(client.RemoteAddr().String())

	if !b.config.IsUdpAssociationAllowed() {
		b.sender.SendConnectionNotAllowedAndClose(client)

		b.logger.UdpAssociationNotAllowed(client.RemoteAddr().String())

		return
	}

	b.udpAssociationHandler.HandleV5UdpAssociation(client)
}
