package v4

import (
	"net"
	v43 "socks/config/v4"
	"socks/handlers/v4/helpers"
	v42 "socks/logger/v4"
	"socks/protocol/v4"
)

type ConnectHandler interface {
	HandleConnect(config v43.Config, address string, client net.Conn)
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

func (b BaseConnectHandler) HandleConnect(config v43.Config, address string, client net.Conn) {
	host, err := b.dialer.Dial(config, address)

	if err != nil {
		b.errorHandler.HandleDialError(config, err, address, client)

		return
	}

	b.connectSendSuccess(config, address, host, client)
}

func (b BaseConnectHandler) connectSendSuccess(config v43.Config, address string, host net.Conn, client net.Conn) {
	err := b.sender.SendSuccess(config, client)

	if err != nil {
		b.errorHandler.HandleConnectIOErrorWithHost(config, err, address, client, host)

		return
	}

	b.logger.Connect.Successful(client.RemoteAddr().String(), address)

	b.transmitter.TransferConnect(config, client, host)
}
