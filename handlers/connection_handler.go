package handlers

import (
	"net"
	"socks/config/tcp"
	"socks/handlers/v4"
	"socks/handlers/v4a"
	"socks/handlers/v5"
	tcp2 "socks/logger/tcp"
	"socks/protocol"
)

type ConnectionHandler struct {
	v4Handler   v4.Handler
	v4aHandler  v4a.Handler
	v5Handler   v5.Handler
	logger      tcp2.Logger
	receiver    protocol.Receiver
	bindHandler BindHandler
	replicator  tcp.ConfigReplicator
}

func NewConnectionHandler(
	v4Handler v4.Handler,
	v4aHandler v4a.Handler,
	v5Handler v5.Handler,
	logger tcp2.Logger,
	receiver protocol.Receiver,
	bindHandler BindHandler,
	replicator tcp.ConfigReplicator,
) (ConnectionHandler, error) {
	return ConnectionHandler{
		v5Handler:   v5Handler,
		v4aHandler:  v4aHandler,
		v4Handler:   v4Handler,
		logger:      logger,
		receiver:    receiver,
		bindHandler: bindHandler,
		replicator:  replicator,
	}, nil
}

func (b ConnectionHandler) HandleConnection(client net.Conn) {
	config := b.replicator.CopyDeadline()

	buffer, err := b.receiver.ReceiveWelcome(config, client)

	if err != nil {
		_ = client.Close()

		b.logger.Errors.ReceiveWelcomeError(client.RemoteAddr().String(), err)

		return
	}

	b.checkProtocol(config, buffer, client)
}

func (b ConnectionHandler) checkProtocol(config tcp.DeadlineConfig, request []byte, client net.Conn) {
	if len(request) < 3 {
		b.bindHandler.Handle(config, request, client)

		return
	}

	if request[0] == 4 {
		b.checkV4(config, request, client)
	} else if request[0] == 5 {
		b.checkV5(config, request, client)
	} else {
		b.bindHandler.Handle(config, request, client)
	}
}

func (b ConnectionHandler) checkV4(config tcp.DeadlineConfig, request []byte, client net.Conn) {
	if len(request) < 9 {
		b.bindHandler.Handle(config, request, client)

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

func (b ConnectionHandler) checkV5(config tcp.DeadlineConfig, request []byte, client net.Conn) {
	if len(request) < 3 {
		b.bindHandler.Handle(config, request, client)

		return
	}

	if int(request[1])+2 != len(request) {
		b.bindHandler.Handle(config, request, client)

		return
	}

	b.logger.Connection.ProtocolDetermined(client.RemoteAddr().String(), "socksV5")

	b.v5Handler.Handle(request, client)
}
