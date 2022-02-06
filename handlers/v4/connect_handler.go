package v4

import (
	"net"
	"socks/handlers/v4/helpers"
	v42 "socks/logger/v4"
	"socks/protocol/v4"
)

type ConnectHandler interface {
	HandleConnect(address string, client net.Conn)
}

type BaseConnectHandler struct {
	logger       v42.Logger
	sender       v4.Sender
	errorHandler ErrorHandler
	dialer       helpers.Dialer
	transmitter  helpers.Transmitter
}

func NewBaseConnectHandler(
	logger v42.Logger,
	sender v4.Sender,
	errorHandler ErrorHandler,
	dialer helpers.Dialer,
	transmitter helpers.Transmitter,
) (BaseConnectHandler, error) {
	return BaseConnectHandler{
		logger:       logger,
		sender:       sender,
		errorHandler: errorHandler,
		dialer:       dialer,
		transmitter:  transmitter,
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

	b.transmitter.TransferConnect(client, host)
}
