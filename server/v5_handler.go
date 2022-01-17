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
	authenticationHandler AuthenticationHandler
	logger                logger.SocksV5Logger
	connectHandler        V5ConnectHandler
	bindHandler           V5BindHandler
	udpAssociationHandler V5UdpAssociationHandler
	sender                V5Sender
}

func NewBaseV5Handler(
	protocol v5.Protocol,
	parser v5.Parser,
	config config.SocksV5Config,
	authenticationHandler AuthenticationHandler,
	logger logger.SocksV5Logger,
	connectHandler V5ConnectHandler,
	bindHandler V5BindHandler,
	udpAssociationHandler V5UdpAssociationHandler,
	sender V5Sender,
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
	}, nil
}

func (b BaseV5Handler) HandleV5(request []byte, client net.Conn) {
	methods, err := b.parser.ParseMethods(request)

	if err != nil {
		b.sender.SendFailAndClose(client)

		return
	}

	b.handleAuthentication(methods, client)
}

func (b BaseV5Handler) handleAuthentication(methods v5.MethodsChunk, client net.Conn) {
	authenticated := b.authenticationHandler.HandleAuthentication(methods, client)

	if !authenticated {
		_ = client.Close()

		b.logger.AuthenticationFailed(client.RemoteAddr().String())

		return
	}

	b.logger.AuthenticationSuccessful(client.RemoteAddr().String())

	b.handleCommand(client)
}

func (b BaseV5Handler) handleCommand(client net.Conn) {
	chunk, parseErr := b.protocol.ReceiveRequest(client)

	if parseErr != nil {
		_ = client.Close()

		return
	}

	var address string

	if chunk.AddressType == 1 {
		address = fmt.Sprintf("%s:%d", chunk.Address, chunk.Port)
	} else if chunk.AddressType == 3 {
		address = fmt.Sprintf("%s:%d", chunk.Address, chunk.Port)
	} else if chunk.AddressType == 4 {
		address = fmt.Sprintf("[%s]:%d", chunk.Address, chunk.Port)
	} else {
		b.sender.SendAddressNotSupportedAndClose(client)

		return
	}

	if chunk.CommandCode == 1 {
		b.handleConnect(address, client)
	} else if chunk.CommandCode == 2 {
		b.handleBind(address, client)
	} else if chunk.CommandCode == 3 {
		b.handleUdpAssociate(client)
	} else {
		b.sender.SendCommandNotSupportedAndClose(client)

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
