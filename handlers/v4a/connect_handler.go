package v4a

import (
	"net"
	"socks/handlers/v4a/helpers"
	v4a2 "socks/logger/v4a"
	"socks/protocol/v4a"
)

type ConnectHandler interface {
	HandleConnect(address string, client net.Conn)
}

type BaseConnectHandler struct {
	logger       v4a2.Logger
	sender       v4a.Sender
	errorHandler ErrorHandler
	dialer       helpers.Dialer
	transmitter  helpers.Transmitter
}

func NewBaseConnectHandler(
	logger v4a2.Logger,
	sender v4a.Sender,
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

	err = b.transmitter.TransferConnect(client, host)

	if err != nil {
		b.errorHandler.HandleTransferError(err, client, host)

		return
	}
}
