package v4a

import (
	"net"
	v4a3 "socks/config/v4a"
	"socks/handlers/v4a/helpers"
	v4a2 "socks/logger/v4a"
	"socks/protocol/v4a"
)

type ConnectHandler interface {
	HandleConnect(config v4a3.Config, address string, client net.Conn)
}

type BaseConnectHandler struct {
	logger       v4a2.Logger
	sender       v4a.Sender
	errorHandler ErrorHandler
	transmitter  helpers.Transmitter
}

func NewBaseConnectHandler(
	logger v4a2.Logger,
	sender v4a.Sender,
	errorHandler ErrorHandler,
	transmitter helpers.Transmitter,
) (BaseConnectHandler, error) {
	return BaseConnectHandler{
		logger:       logger,
		sender:       sender,
		errorHandler: errorHandler,
		transmitter:  transmitter,
	}, nil
}

func (b BaseConnectHandler) HandleConnect(config v4a3.Config, address string, client net.Conn) {
	host, err := net.DialTimeout("tcp4", address, config.Deadline.Connect)

	if err != nil {
		b.errorHandler.HandleDialError(config, err, address, client)

		return
	}

	b.connectSendSuccess(config, address, host, client)
}

func (b BaseConnectHandler) connectSendSuccess(config v4a3.Config, address string, host net.Conn, client net.Conn) {
	err := b.sender.SendSuccess(config, client)

	if err != nil {
		b.errorHandler.HandleConnectIOErrorWithHost(config, err, address, client, host)

		return
	}

	b.logger.Connect.Successful(client.RemoteAddr().String(), address)

	b.transmitter.TransferConnect(config, client, host)
}
