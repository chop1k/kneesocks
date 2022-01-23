package server

import (
	"fmt"
	"net"
	"socks/config"
	"socks/logger"
	"socks/utils"
	"time"
)

type ConnectionHandler interface {
	HandleConnection(client net.Conn)
}

type BaseConnectionHandler struct {
	streamHandler StreamHandler
	v4Handler     V4Handler
	v4aHandler    V4aHandler
	v5Handler     V5Handler
	bindManager   BindManager
	utils         utils.AddressUtils
	logger        logger.TcpLogger
	config        config.TcpConfig
}

func NewBaseConnectionHandler(
	streamHandler StreamHandler,
	v4Handler V4Handler,
	v4aHandler V4aHandler,
	v5Handler V5Handler,
	bindManager BindManager,
	utils utils.AddressUtils,
	logger logger.TcpLogger,
	config config.TcpConfig,
) (BaseConnectionHandler, error) {
	return BaseConnectionHandler{
		streamHandler: streamHandler,
		v5Handler:     v5Handler,
		v4aHandler:    v4aHandler,
		v4Handler:     v4Handler,
		bindManager:   bindManager,
		utils:         utils,
		logger:        logger,
		config:        config,
	}, nil
}

func (b BaseConnectionHandler) HandleConnection(client net.Conn) {
	buffer := make([]byte, 512)

	i, err := client.Read(buffer)

	if err != nil {
		_ = client.Close()

		return
	}

	b.checkProtocol(buffer[:i], client)
}

func (b BaseConnectionHandler) checkProtocol(request []byte, client net.Conn) {
	if len(request) < 3 {
		b.checkBound(request, client)

		return
	}

	if request[0] == 4 {
		b.checkV4(request, client)
	} else if request[0] == 5 {
		b.checkV5(request, client)
	} else {
		b.checkBound(request, client)
	}
}

func (b BaseConnectionHandler) checkV4(request []byte, client net.Conn) {
	if len(request) < 9 {
		b.checkBound(request, client)

		return
	}

	if request[len(request)-1] != 0 {
		b.checkBound(request, client)

		return
	}

	if request[4] == 0 && request[5] == 0 && request[6] == 0 && request[7] != 0 {
		b.logger.ConnectionProtocolDetermined(client.RemoteAddr().String(), "socksV4a")

		b.v4aHandler.HandleV4a(request, client)
	} else {
		b.logger.ConnectionProtocolDetermined(client.RemoteAddr().String(), "socksV4")

		b.v4Handler.HandleV4(request, client)
	}
}

func (b BaseConnectionHandler) checkV5(request []byte, client net.Conn) {
	if len(request) < 3 {
		b.checkBound(request, client)

		return
	}

	if int(request[1])+2 != len(request) {
		b.checkBound(request, client)

		return
	}

	b.logger.ConnectionProtocolDetermined(client.RemoteAddr().String(), "socksV5")

	b.v5Handler.HandleV5(request, client)
}

func (b BaseConnectionHandler) checkBound(request []byte, client net.Conn) {
	addr := client.RemoteAddr().String()

	if b.bindManager.IsBound(addr) {
		b.exchange(request, addr, client)
	} else {
		b.checkDomain(request, addr, client)
	}
}

func (b BaseConnectionHandler) checkDomain(request []byte, addr string, client net.Conn) {
	hostAddr, hostPort, parseErr := b.utils.ParseAddress(addr)

	if parseErr != nil {
		_ = client.Close()

		b.logger.AddressParseError(addr, parseErr)

		return
	}

	addresses, err := net.LookupAddr(hostAddr)

	if err != nil {
		_ = client.Close()

		b.logger.LookupError(hostAddr, err)

		return
	}

	for _, address := range addresses {
		address = fmt.Sprintf("%s:%d", address, hostPort)

		if b.bindManager.IsBound(address) {
			b.exchange(request, address, client)

			return
		}
	}

	_ = client.Close()

	b.logger.ConnectionDenied(addr)
}

func (b BaseConnectionHandler) exchange(request []byte, addr string, client net.Conn) {
	err := b.bindManager.SendHost(addr, client)

	if err != nil {
		_ = client.Close()

		b.logger.SendHostError(addr, err)

		return
	}

	deadline := time.Second * time.Duration(b.config.GetExchangeDeadline())

	host, receiveErr := b.bindManager.ReceiveClient(addr, deadline)

	if receiveErr != nil {
		_ = client.Close()

		b.logger.ReceiveClientError(addr, receiveErr)

		return
	}

	b.logger.ConnectionBound(client.RemoteAddr().String(), host.RemoteAddr().String())

	_, err = host.Write(request)

	if err != nil {
		_ = client.Close()
		_ = host.Close()

		b.logger.WriteRequestError(client.RemoteAddr().String(), host.RemoteAddr().String(), err)

		return
	}

	b.streamHandler.HostToClient(client, host)
}
