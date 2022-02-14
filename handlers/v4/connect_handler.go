package v4

import (
	"net"
	v43 "socks/config/v4"
	"socks/handlers/v4/helpers"
	v42 "socks/logger/v4"
	"socks/protocol/v4"
)

type ConnectHandler struct {
	logger       v42.Logger
	sender       v4.Sender
	errorHandler ErrorHandler
	transmitter  helpers.Transmitter
}

func NewConnectHandler(
	logger v42.Logger,
	sender v4.Sender,
	errorHandler ErrorHandler,
	transmitter helpers.Transmitter,
) (ConnectHandler, error) {
	return ConnectHandler{
		logger:       logger,
		sender:       sender,
		errorHandler: errorHandler,
		transmitter:  transmitter,
	}, nil
}

func (b ConnectHandler) HandleConnect(config v43.Config, address string, client net.Conn) {
	var host net.Conn
	var err error

	if config.Deadline.Connect > 0 {
		host, err = net.DialTimeout("tcp4", address, config.Deadline.Connect)
	} else {
		host, err = net.Dial("tcp4", address)
	}

	if err != nil {
		b.errorHandler.HandleDialError(config, err, address, client)

		return
	}

	b.connectSendSuccess(config, address, host, client)
}

func (b ConnectHandler) connectSendSuccess(config v43.Config, address string, host net.Conn, client net.Conn) {
	err := b.sender.SendSuccess(config, client)

	if err != nil {
		b.errorHandler.HandleConnectIOErrorWithHost(config, err, address, client, host)

		return
	}

	b.logger.Connect.Successful(client.RemoteAddr().String(), address)

	b.transmitter.TransferConnect(config, client, host)
}
