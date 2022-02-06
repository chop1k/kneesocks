package handlers

import (
	"net"
	"socks/handlers/v4"
	"socks/handlers/v4a"
	"socks/handlers/v5"
	tcp2 "socks/logger/tcp"
	"socks/protocol"
)

type ConnectionHandler interface {
	HandleConnection(client net.Conn)
}

type BaseConnectionHandler struct {
	v4Handler   v4.Handler
	v4aHandler  v4a.Handler
	v5Handler   v5.Handler
	logger      tcp2.Logger
	receiver    protocol.Receiver
	bindHandler BindHandler
}

func NewBaseConnectionHandler(
	v4Handler v4.Handler,
	v4aHandler v4a.Handler,
	v5Handler v5.Handler,
	logger tcp2.Logger,
	receiver protocol.Receiver,
	bindHandler BindHandler,
) (BaseConnectionHandler, error) {
	return BaseConnectionHandler{
		v5Handler:   v5Handler,
		v4aHandler:  v4aHandler,
		v4Handler:   v4Handler,
		logger:      logger,
		receiver:    receiver,
		bindHandler: bindHandler,
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
		b.bindHandler.Handle(request, client)

		return
	}

	if request[0] == 4 {
		b.checkV4(request, client)
	} else if request[0] == 5 {
		b.checkV5(request, client)
	} else {
		b.bindHandler.Handle(request, client)
	}
}

func (b BaseConnectionHandler) checkV4(request []byte, client net.Conn) {
	if len(request) < 9 {
		b.bindHandler.Handle(request, client)

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
		b.bindHandler.Handle(request, client)

		return
	}

	if int(request[1])+2 != len(request) {
		b.bindHandler.Handle(request, client)

		return
	}

	b.logger.Connection.ProtocolDetermined(client.RemoteAddr().String(), "socksV5")

	b.v5Handler.Handle(request, client)
}
