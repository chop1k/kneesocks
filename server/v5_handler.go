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
	tcpConfig             config.TcpConfig
	connectHandler        V5ConnectHandler
	bindHandler           V5BindHandler
	udpAssociationHandler V5UdpAssociationHandler
}

func NewBaseV5Handler(
	protocol v5.Protocol,
	parser v5.Parser,
	config config.SocksV5Config,
	authenticationHandler AuthenticationHandler,
	logger logger.SocksV5Logger,
	tcpConfig config.TcpConfig,
	connectHandler V5ConnectHandler,
	bindHandler V5BindHandler,
	udpAssociationHandler V5UdpAssociationHandler,
) (BaseV5Handler, error) {
	return BaseV5Handler{
		protocol:              protocol,
		parser:                parser,
		config:                config,
		authenticationHandler: authenticationHandler,
		logger:                logger,
		tcpConfig:             tcpConfig,
		connectHandler:        connectHandler,
		bindHandler:           bindHandler,
		udpAssociationHandler: udpAssociationHandler,
	}, nil
}

func (b BaseV5Handler) HandleV5(request []byte, client net.Conn) {
	methods, err := b.parser.ParseMethods(request)

	if err != nil {
		b.sendFailAndClose(client)

		return
	}

	b.handleAuthentication(methods, client)
}

func (b BaseV5Handler) sendFailAndClose(client net.Conn) {
	_ = b.protocol.ResponseWithFail(1, "0.0.0.0", uint16(b.tcpConfig.GetBindPort()), client)
	_ = client.Close()
}

func (b BaseV5Handler) sendCommandNotSupportedAndClose(client net.Conn) {
	_ = b.protocol.ResponseWithCommandNotSupported(1, "0.0.0.0", uint16(b.tcpConfig.GetBindPort()), client)
	_ = client.Close()
}

func (b BaseV5Handler) sendConnectionNotAllowedAndClose(client net.Conn) {
	_ = b.protocol.ResponseWithNotAllowed(1, "0.0.0.0", uint16(b.tcpConfig.GetBindPort()), client)
	_ = client.Close()
}

func (b BaseV5Handler) sendAddressNotSupportedAndClose(client net.Conn) {
	_ = b.protocol.ResponseWithAddressNotSupported(1, "0.0.0.0", uint16(b.tcpConfig.GetBindPort()), client)
	_ = client.Close()
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
		b.sendAddressNotSupportedAndClose(client)

		return
	}

	if chunk.CommandCode == 1 {
		b.handleConnect(address, client)
	} else if chunk.CommandCode == 2 {
		b.handleBind(address, client)
	} else if chunk.CommandCode == 3 {
		b.handleUdpAssociate(client)
	} else {
		b.sendCommandNotSupportedAndClose(client)

		return
	}
}

func (b BaseV5Handler) handleConnect(address string, client net.Conn) {
	b.logger.ConnectRequest(client.RemoteAddr().String(), address)

	if !b.config.IsConnectAllowed() {
		b.sendConnectionNotAllowedAndClose(client)

		b.logger.ConnectNotAllowed(client.RemoteAddr().String(), address)

		return
	}

	b.connectHandler.HandleV5Connect(address, client)
}

func (b BaseV5Handler) handleBind(address string, client net.Conn) {
	b.logger.BindRequest(client.RemoteAddr().String(), address)

	if !b.config.IsBindAllowed() {
		b.sendConnectionNotAllowedAndClose(client)

		b.logger.BindNotAllowed(client.RemoteAddr().String(), address)

		return
	}

	b.bindHandler.HandleV5Bind(address, client)
}

func (b BaseV5Handler) handleUdpAssociate(client net.Conn) {
	b.logger.UdpAssociationRequest(client.RemoteAddr().String())

	if !b.config.IsUdpAssociationAllowed() {
		b.sendConnectionNotAllowedAndClose(client)

		b.logger.UdpAssociationNotAllowed(client.RemoteAddr().String())

		return
	}

	b.udpAssociationHandler.HandleV5UdpAssociation(client)
}
