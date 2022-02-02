package v4a

import (
	"net"
	"socks/handlers/v4a/helpers"
	v4a2 "socks/logger/v4a"
	"socks/transfer"
)

type ConnectHandler interface {
	HandleConnect(address string, client net.Conn)
}

type BaseConnectHandler struct {
	streamHandler transfer.StreamHandler
	logger        v4a2.Logger
	sender        helpers.Sender
	errorHandler  ErrorHandler
	dialer        helpers.Dialer
}

func NewBaseConnectHandler(
	streamHandler transfer.StreamHandler,
	logger v4a2.Logger,
	sender helpers.Sender,
	errorHandler ErrorHandler,
	dialer helpers.Dialer,
) (BaseConnectHandler, error) {
	return BaseConnectHandler{
		streamHandler: streamHandler,
		logger:        logger,
		sender:        sender,
		errorHandler:  errorHandler,
		dialer:        dialer,
	}, nil
}

func (b BaseConnectHandler) HandleConnect(address string, client net.Conn) {
	host, err := b.dialer.Dial(address)

	if err != nil {
		b.errorHandler.HandleDialError(err, address, client)

		return
	}

	b.connectSendSuccess(address, host, client)
}

func (b BaseConnectHandler) connectSendSuccess(address string, host net.Conn, client net.Conn) {
	err := b.sender.SendSuccess(client)

	if err != nil {
		b.errorHandler.HandleConnectIOErrorWithHost(err, address, client, host)

		return
	}

	b.logger.Connect.Successful(client.RemoteAddr().String(), address)

	go b.streamHandler.ClientToHost(client, host)
	b.streamHandler.HostToClient(client, host)
}
