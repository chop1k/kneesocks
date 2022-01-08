package server

import (
	"fmt"
	"net"
	config2 "socks/config/tree"
	"socks/logger"
	"socks/protocol/v4a"
	"socks/utils"
)

type V4aHandler interface {
	HandleV4a(request []byte, client net.Conn)
}

type BaseV4aHandler struct {
	protocol      v4a.Protocol
	parser        v4a.Parser
	bindManager   BindManager
	config        config2.Config
	streamHandler StreamHandler
	utils         utils.AddressUtils
	logger        logger.SocksV4aLogger
}

func NewBaseV4aHandler(
	protocol v4a.Protocol,
	parser v4a.Parser,
	bindManager BindManager,
	config config2.Config,
	streamHandler StreamHandler,
	utils utils.AddressUtils,
	logger logger.SocksV4aLogger,
) BaseV4aHandler {
	return BaseV4aHandler{
		protocol:      protocol,
		parser:        parser,
		bindManager:   bindManager,
		config:        config,
		streamHandler: streamHandler,
		utils:         utils,
		logger:        logger,
	}
}

func (b BaseV4aHandler) HandleV4a(request []byte, client net.Conn) {
	chunk, err := b.parser.ParseRequest(request)

	if err != nil {
		b.sendFailAndClose(client)

		return
	}

	if chunk.CommandCode == 1 {
		go b.logger.ConnectRequest(client.RemoteAddr().String(), chunk)

		b.handleConnect(chunk, client)
	} else if chunk.CommandCode == 2 {
		go b.logger.BindRequest(client.RemoteAddr().String(), chunk)

		b.handleBind(chunk, client)
	} else {
		b.sendFailAndClose(client)
	}
}

func (b BaseV4aHandler) sendFailAndClose(client net.Conn) {
	_ = b.protocol.ResponseWithCode(91, uint16(b.config.Tcp.BindPort), net.IP{0, 0, 0, 0}, client)
	_ = client.Close()
}

func (b BaseV4aHandler) handleConnect(chunk v4a.RequestChunk, client net.Conn) {
	addr := fmt.Sprintf("%s:%d", chunk.Domain, chunk.DestinationPort)

	host, err := net.Dial("tcp", addr)

	if err != nil {
		b.sendFailAndClose(client)

		go b.logger.ConnectFailed(client.RemoteAddr().String(), chunk)

		return
	}

	b.connectSendSuccess(chunk, host, client)
}

func (b BaseV4aHandler) connectSendSuccess(chunk v4a.RequestChunk, host net.Conn, client net.Conn) {
	err := b.protocol.ResponseWithCode(90, uint16(b.config.Tcp.BindPort), net.IP{0, 0, 0, 0}, client)

	if err != nil {
		b.sendFailAndClose(client)

		_ = host.Close()

		go b.logger.ConnectFailed(client.RemoteAddr().String(), chunk)

		return
	}

	go b.logger.ConnectSuccessful(client.RemoteAddr().String(), chunk)

	go b.streamHandler.ClientToHost(client, host)
	b.streamHandler.HostToClient(client, host)
}

func (b BaseV4aHandler) handleBind(chunk v4a.RequestChunk, client net.Conn) {
	address := fmt.Sprintf("%s:%d", chunk.Domain, chunk.DestinationPort)

	err := b.bindManager.Bind(address)

	if err != nil {
		b.sendFailAndClose(client)

		go b.logger.BindFailed(client.RemoteAddr().String(), chunk)

		return
	}

	go b.logger.BindSuccessful(client.RemoteAddr().String(), chunk)

	b.bindSendFirstResponse(chunk, address, client)
}

func (b BaseV4aHandler) bindSendFirstResponse(chunk v4a.RequestChunk, address string, client net.Conn) {
	err := b.protocol.ResponseWithCode(90, 0, net.IP{0, 0, 0, 0}, client)

	if err != nil {
		b.sendFailAndClose(client)

		go b.logger.BindFailed(client.RemoteAddr().String(), chunk)

		return
	}

	b.bindWait(chunk, address, client)

	b.bindManager.Remove(address)
}

func (b BaseV4aHandler) bindWait(chunk v4a.RequestChunk, address string, client net.Conn) {
	host, waitErr := b.bindManager.ReceiveHost(address)

	if waitErr != nil {
		b.sendFailAndClose(client)

		go b.logger.BindFailed(client.RemoteAddr().String(), chunk)

		return
	}

	b.bindCheckAddress(chunk, address, host, client)
}

func (b BaseV4aHandler) bindCheckAddress(chunk v4a.RequestChunk, address string, host net.Conn, client net.Conn) {
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

	if addrType != 1 {
		b.sendFailAndClose(client)

		_ = host.Close()

		go b.logger.BindFailed(client.RemoteAddr().String(), chunk)

		return
	}

	b.bindSendSecondResponse(chunk, address, hostAddr, uint16(hostPort), host, client)
}

func (b BaseV4aHandler) bindSendSecondResponse(chunk v4a.RequestChunk, address string, hostAddr string, hostPort uint16, host net.Conn, client net.Conn) {
	ip := net.ParseIP(hostAddr).To4()

	err := b.protocol.ResponseWithCode(90, hostPort, ip, client)

	if err != nil {
		b.sendFailAndClose(client)

		_ = host.Close()

		go b.logger.BindFailed(client.RemoteAddr().String(), chunk)

		return
	}

	err = b.bindManager.SendClient(address, client)

	if err != nil {
		_ = client.Close()
		_ = host.Close()

		go b.logger.BindFailed(client.RemoteAddr().String(), chunk)

		return
	}

	go b.logger.BindSuccessful(client.RemoteAddr().String(), chunk)

	b.streamHandler.ClientToHost(host, client)
}
