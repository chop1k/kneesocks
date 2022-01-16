package server

import (
	"fmt"
	"net"
	"os"
	"socks/config"
	"socks/logger"
	v5 "socks/protocol/v5"
	"socks/utils"
	"syscall"
	"time"
)

type V5Handler interface {
	HandleV5(request []byte, client net.Conn)
}

type BaseV5Handler struct {
	protocol              v5.Protocol
	parser                v5.Parser
	bindManager           BindManager
	config                config.SocksV5Config
	streamHandler         StreamHandler
	utils                 utils.AddressUtils
	udpAssociationManager UdpAssociationManager
	authenticationHandler AuthenticationHandler
	logger                logger.SocksV5Logger
	tcpConfig             config.TcpConfig
	udpConfig             config.UdpConfig
}

func NewBaseV5Handler(
	protocol v5.Protocol,
	parser v5.Parser,
	bindManager BindManager,
	config config.SocksV5Config,
	streamHandler StreamHandler,
	utils utils.AddressUtils,
	udpAssociationManager UdpAssociationManager,
	authenticationHandler AuthenticationHandler,
	logger logger.SocksV5Logger,
	tcpConfig config.TcpConfig,
	udpConfig config.UdpConfig,
) (BaseV5Handler, error) {
	return BaseV5Handler{
		protocol:              protocol,
		parser:                parser,
		bindManager:           bindManager,
		config:                config,
		streamHandler:         streamHandler,
		utils:                 utils,
		udpAssociationManager: udpAssociationManager,
		authenticationHandler: authenticationHandler,
		logger:                logger,
		tcpConfig:             tcpConfig,
		udpConfig:             udpConfig,
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
	_ = b.protocol.ResponseWithCode(1, 1, "0.0.0.0", uint16(b.tcpConfig.GetBindPort()), client)
	_ = client.Close()
}

func (b BaseV5Handler) sendCommandNotSupportedAndClose(client net.Conn) {
	_ = b.protocol.ResponseWithCode(7, 1, "0.0.0.0", uint16(b.tcpConfig.GetBindPort()), client)
	_ = client.Close()
}

func (b BaseV5Handler) sendAddressNotSupportedAndClose(client net.Conn) {
	_ = b.protocol.ResponseWithCode(8, 1, "0.0.0.0", uint16(b.tcpConfig.GetBindPort()), client)
	_ = client.Close()
}

func (b BaseV5Handler) sendConnectionRefusedAndClose(client net.Conn) {
	_ = b.protocol.ResponseWithCode(5, 1, "0.0.0.0", uint16(b.tcpConfig.GetBindPort()), client)
	_ = client.Close()
}

func (b BaseV5Handler) sendNetworkUnreachableAndClose(client net.Conn) {
	_ = b.protocol.ResponseWithCode(3, 1, "0.0.0.0", uint16(b.tcpConfig.GetBindPort()), client)
	_ = client.Close()
}

func (b BaseV5Handler) sendHostUnreachableAndClose(client net.Conn) {
	_ = b.protocol.ResponseWithCode(4, 1, "0.0.0.0", uint16(b.tcpConfig.GetBindPort()), client)
	_ = client.Close()
}

func (b BaseV5Handler) sendConnectionNotAllowedAndClose(client net.Conn) {
	_ = b.protocol.ResponseWithCode(2, 1, "0.0.0.0", uint16(b.tcpConfig.GetBindPort()), client)
	_ = client.Close()
}

func (b BaseV5Handler) handleAuthentication(methods v5.MethodsChunk, client net.Conn) {
	authenticated := b.authenticationHandler.HandleAuthentication(methods, client)

	if !authenticated {
		_ = client.Close()

		go b.logger.AuthenticationFailed(client.RemoteAddr().String())

		return
	}

	go b.logger.AuthenticationSuccessful(client.RemoteAddr().String())

	b.handleRequest(client)
}

func (b BaseV5Handler) handleRequest(client net.Conn) {
	request := make([]byte, 1024)

	i, err := client.Read(request)

	if err != nil {
		_ = client.Close()

		return
	}

	b.handleCommand(request[:i], client)
}

func (b BaseV5Handler) handleCommand(request []byte, client net.Conn) {
	chunk, parseErr := b.parser.ParseRequest(request)

	if parseErr != nil {
		_ = client.Close()

		return
	}

	if chunk.CommandCode == 1 {
		go b.logger.ConnectRequest(client.RemoteAddr().String(), chunk)

		if !b.config.IsConnectAllowed() {
			b.sendConnectionNotAllowedAndClose(client)

			go b.logger.ConnectNowAllowed(client.RemoteAddr().String(), chunk)

			return
		}

		b.handleConnect(chunk, client)
	} else if chunk.CommandCode == 2 {
		go b.logger.BindRequest(client.RemoteAddr().String(), chunk)

		if !b.config.IsBindAllowed() {
			b.sendConnectionNotAllowedAndClose(client)

			go b.logger.BindNotAllowed(client.RemoteAddr().String(), chunk)

			return
		}

		b.handleBind(chunk, client)
	} else if chunk.CommandCode == 3 {
		go b.logger.UdpAssociationRequest(client.RemoteAddr().String(), chunk)

		if !b.config.IsUdpAssociationAllowed() {
			b.sendConnectionNotAllowedAndClose(client)

			go b.logger.UdpAssociationNotAllowed(client.RemoteAddr().String(), chunk)

			return
		}

		b.handleUdpAssociate(chunk, client)
	} else {
		b.sendCommandNotSupportedAndClose(client)

		return
	}
}

func (b BaseV5Handler) handleConnect(chunk v5.RequestChunk, client net.Conn) {
	var host net.Conn
	var err error

	deadline := time.Second * time.Duration(b.config.GetConnectDeadline())

	if chunk.AddressType == 1 {
		host, err = net.DialTimeout("tcp4", fmt.Sprintf("%s:%d", chunk.Address, chunk.Port), deadline)
	} else if chunk.AddressType == 3 {
		host, err = net.DialTimeout("tcp", fmt.Sprintf("%s:%d", chunk.Address, chunk.Port), deadline)
	} else if chunk.AddressType == 4 {
		host, err = net.DialTimeout("tcp6", fmt.Sprintf("[%s]:%d", chunk.Address, chunk.Port), deadline)
	} else {
		b.sendAddressNotSupportedAndClose(client)

		go b.logger.ConnectFailed(client.RemoteAddr().String(), chunk)

		return
	}

	if err != nil {
		if b.checkConnectionRefusedError(err) {
			b.sendConnectionRefusedAndClose(client)

			go b.logger.ConnectFailed(client.RemoteAddr().String(), chunk)

			return
		}

		if b.checkNetworkUnreachableError(err) {
			b.sendNetworkUnreachableAndClose(client)

			go b.logger.ConnectFailed(client.RemoteAddr().String(), chunk)

			return
		}

		if b.checkHostUnreachableError(err) {
			b.sendHostUnreachableAndClose(client)

			go b.logger.ConnectFailed(client.RemoteAddr().String(), chunk)

			return
		}

		b.sendFailAndClose(client)

		go b.logger.ConnectFailed(client.RemoteAddr().String(), chunk)

		return
	}

	b.connectSendResponse(chunk, host, client)
}

func (b BaseV5Handler) connectSendResponse(chunk v5.RequestChunk, host, client net.Conn) {
	addr, port, parseErr := b.utils.ParseAddress(host.RemoteAddr().String())

	if parseErr != nil {
		b.sendFailAndClose(client)

		_ = host.Close()

		return
	}

	addrType, determineErr := b.utils.DetermineAddressType(addr)

	if determineErr != nil {
		b.sendFailAndClose(client)

		_ = host.Close()

		return
	}

	responseErr := b.protocol.ResponseWithCode(0, addrType, addr, uint16(port), client)

	if responseErr != nil {
		b.sendFailAndClose(client)

		_ = host.Close()

		return
	}

	go b.logger.ConnectSuccessful(client.RemoteAddr().String(), chunk)

	go b.streamHandler.ClientToHost(client, host)
	b.streamHandler.HostToClient(client, host)
}

func (b BaseV5Handler) errorToErrno(err error) int {
	opErr, ok := err.(*net.OpError)

	if !ok {
		return -1
	}

	sysErr, ko := opErr.Err.(*os.SyscallError)

	if !ko {
		return -1
	}

	errno, oo := sysErr.Err.(syscall.Errno)

	if !oo {
		return -1
	}

	return int(errno)
}

func (b BaseV5Handler) checkConnectionRefusedError(err error) bool {
	return b.errorToErrno(err) == 111
}

func (b BaseV5Handler) checkNetworkUnreachableError(err error) bool {
	return b.errorToErrno(err) == 101
}

func (b BaseV5Handler) checkHostUnreachableError(err error) bool {
	errno := b.errorToErrno(err)

	return errno == 113 || errno == 112
}
func (b BaseV5Handler) handleBind(chunk v5.RequestChunk, client net.Conn) {
	address := fmt.Sprintf("%s:%d", chunk.Address, chunk.Port)

	err := b.bindManager.Bind(address)

	if err != nil {
		b.sendConnectionNotAllowedAndClose(client)

		go b.logger.BindFailed(client.RemoteAddr().String(), chunk)

		return
	}

	go b.logger.BindSuccessful(client.RemoteAddr().String(), chunk)

	b.bindSendFirstResponse(chunk, address, client)
}

func (b BaseV5Handler) bindSendFirstResponse(chunk v5.RequestChunk, address string, client net.Conn) {
	err := b.protocol.ResponseWithCode(0, 1, "0.0.0.0", uint16(b.tcpConfig.GetBindPort()), client)

	if err != nil {
		_ = client.Close()

		go b.logger.BindFailed(client.RemoteAddr().String(), chunk)

		return
	}

	b.bindWait(chunk, address, client)

	b.bindManager.Remove(address)
}

func (b BaseV5Handler) bindWait(chunk v5.RequestChunk, address string, client net.Conn) {
	deadline := time.Second * time.Duration(b.config.GetBindDeadline())

	host, err := b.bindManager.ReceiveHost(address, deadline)

	if err != nil {
		_ = client.Close()

		go b.logger.BindFailed(client.RemoteAddr().String(), chunk)

		return
	}

	b.bindCheckAddress(chunk, address, host, client)
}

func (b BaseV5Handler) bindCheckAddress(chunk v5.RequestChunk, address string, host, client net.Conn) {
	hostAddr, hostPort, parseErr := b.utils.ParseAddress(host.RemoteAddr().String())

	if parseErr != nil {
		b.sendFailAndClose(client)

		_ = host.Close()

		go b.logger.BindFailed(client.RemoteAddr().String(), chunk)

		return
	}

	addrType, determineErr := b.utils.DetermineAddressType(hostAddr)

	if determineErr != nil {
		b.sendFailAndClose(client)

		_ = host.Close()

		go b.logger.BindFailed(client.RemoteAddr().String(), chunk)

		return
	}

	b.sendSecondResponse(chunk, address, addrType, hostAddr, uint16(hostPort), host, client)
}

func (b BaseV5Handler) sendSecondResponse(chunk v5.RequestChunk, address string, addrType byte, hostAddress string, hostPort uint16, host, client net.Conn) {
	err := b.protocol.ResponseWithCode(0, addrType, hostAddress, hostPort, client)

	if err != nil {
		b.sendFailAndClose(client)

		_ = host.Close()

		go b.logger.BindFailed(client.RemoteAddr().String(), chunk)

		return
	}

	err = b.bindManager.SendClient(address, client)

	if err != nil {
		b.sendFailAndClose(client)

		_ = host.Close()

		go b.logger.BindFailed(client.RemoteAddr().String(), chunk)

		return
	}

	go b.logger.Bound(client.RemoteAddr().String(), host.RemoteAddr().String(), chunk)

	b.streamHandler.ClientToHost(host, client)
}

func (b BaseV5Handler) handleUdpAssociate(chunk v5.RequestChunk, client net.Conn) {
	address, _, err := b.utils.ParseAddress(client.RemoteAddr().String())

	if err != nil {
		b.sendFailAndClose(client)

		go b.logger.UdpAssociationFailed(client.RemoteAddr().String(), chunk)

		return
	}

	b.udpAssociationManager.Set(address, client.RemoteAddr())

	b.udpSendResponse(chunk, address, client)
}

func (b BaseV5Handler) udpSendResponse(chunk v5.RequestChunk, address string, client net.Conn) {
	err := b.protocol.ResponseWithCode(0, 1, "0.0.0.0", uint16(b.udpConfig.GetBindPort()), client)

	if err != nil {
		b.sendFailAndClose(client)

		go b.logger.UdpAssociationFailed(client.RemoteAddr().String(), chunk)

		return
	}

	go b.logger.UdpAssociationSuccessful(client.RemoteAddr().String(), chunk)

	b.udpWaitForClose(address, client)
}

func (b BaseV5Handler) udpWaitForClose(address string, client net.Conn) {
	for {
		buffer := make([]byte, 512)

		_, err := client.Read(buffer)

		_ = buffer

		if err != nil {
			_ = client.Close()

			break
		}
	}

	b.udpAssociationManager.Remove(address)
}
