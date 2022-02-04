package handlers

import (
	"fmt"
	"net"
	"socks/handlers/helpers"
	"socks/handlers/v4"
	"socks/handlers/v4a"
	"socks/handlers/v5"
	tcp2 "socks/logger/tcp"
	"socks/protocol"
	"socks/transfer"
	"socks/utils"
)

type ConnectionHandler interface {
	HandleConnection(client net.Conn)
}

type BaseConnectionHandler struct {
	streamHandler transfer.StreamHandler
	v4Handler     v4.Handler
	v4aHandler    v4a.Handler
	v5Handler     v5.Handler
	utils         utils.AddressUtils
	logger        tcp2.Logger
	receiver      protocol.Receiver
	binder        helpers.Binder
}

func NewBaseConnectionHandler(
	streamHandler transfer.StreamHandler,
	v4Handler v4.Handler,
	v4aHandler v4a.Handler,
	v5Handler v5.Handler,
	utils utils.AddressUtils,
	logger tcp2.Logger,
	receiver protocol.Receiver,
	binder helpers.Binder,
) (BaseConnectionHandler, error) {
	return BaseConnectionHandler{
		streamHandler: streamHandler,
		v5Handler:     v5Handler,
		v4aHandler:    v4aHandler,
		v4Handler:     v4Handler,
		utils:         utils,
		logger:        logger,
		receiver:      receiver,
		binder:        binder,
	}, nil
}

func (b BaseConnectionHandler) HandleConnection(client net.Conn) {
	buffer, err := b.receiver.ReceiveWelcome(client)

	if err != nil {
		_ = client.Close()

		b.logger.Errors.ReceiveWelcomeError(client.RemoteAddr().String(), err)

		return
	}

	b.checkProtocol(buffer, client)
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

	if request[4] == 0 && request[5] == 0 && request[6] == 0 && request[7] != 0 {
		b.logger.Connection.ProtocolDetermined(client.RemoteAddr().String(), "socksV4a")

		b.v4aHandler.Handle(request, client)
	} else {
		b.logger.Connection.ProtocolDetermined(client.RemoteAddr().String(), "socksV4")

		b.v4Handler.Handle(request, client)
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

	b.logger.Connection.ProtocolDetermined(client.RemoteAddr().String(), "socksV5")

	b.v5Handler.Handle(request, client)
}

func (b BaseConnectionHandler) checkBound(request []byte, client net.Conn) {
	addr := client.RemoteAddr().String()

	if b.binder.IsBound(addr) {
		b.exchange(request, addr, client)
	} else {
		b.checkDomain(request, addr, client)
	}
}

func (b BaseConnectionHandler) checkDomain(request []byte, address string, client net.Conn) {
	hostAddr, hostPort, parseErr := b.utils.ParseAddress(address)

	if parseErr != nil {
		_ = client.Close()

		b.logger.Errors.AddressParseError(address, parseErr)

		return
	}

	addresses, err := net.LookupAddr(hostAddr)

	if err != nil {
		_ = client.Close()

		b.logger.Errors.LookupError(hostAddr, err)

		return
	}

	for _, address := range addresses {
		address = fmt.Sprintf("%s:%d", address, hostPort)

		if b.binder.IsBound(address) {
			b.exchange(request, address, client)

			return
		}
	}

	_ = client.Close()

	b.logger.Connection.Denied(address)
}

func (b BaseConnectionHandler) exchange(request []byte, address string, client net.Conn) {
	err := b.binder.Send(address, client)

	if err != nil {
		_ = client.Close()

		b.logger.Errors.SendHostError(address, err)

		return
	}

	host, receiveErr := b.binder.Receive(address)

	if receiveErr != nil {
		_ = client.Close()

		b.logger.Errors.ReceiveClientError(address, receiveErr)

		return
	}

	b.logger.Connection.Bound(client.RemoteAddr().String(), host.RemoteAddr().String())

	_, err = host.Write(request)

	if err != nil {
		_ = client.Close()
		_ = host.Close()

		b.logger.Errors.WriteRequestError(client.RemoteAddr().String(), host.RemoteAddr().String(), err)

		return
	}

	b.streamHandler.ClientToHost(client, host)
}
