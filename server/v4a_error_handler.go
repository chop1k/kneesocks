package server

import (
	"net"
	"socks/logger"
	"socks/utils"
)

type V4aErrorHandler interface {
	HandleV4aDialError(err error, address string, client net.Conn)
	HandleV4aConnectIOError(err error, address string, client net.Conn)
	HandleV4aConnectIOErrorWithHost(err error, address string, client net.Conn, host net.Conn)
	HandleV4aBindIOError(err error, address string, client net.Conn)
	HandleV4aBindIOErrorWithHost(err error, address string, client net.Conn, host net.Conn)
}

type BaseV4aErrorHandler struct {
	logger logger.SocksV4aLogger
	sender V4aSender
	errors utils.ErrorUtils
}

func NewBaseV4aErrorHandler(
	logger logger.SocksV4aLogger,
	sender V4aSender,
	errors utils.ErrorUtils,
) (BaseV4aErrorHandler, error) {
	return BaseV4aErrorHandler{
		logger: logger,
		sender: sender,
		errors: errors,
	}, nil
}

func (b BaseV4aErrorHandler) HandleV4aDialError(err error, address string, client net.Conn) {
	if b.errors.IsConnectionRefusedError(err) {
		b.logger.ConnectRefused(client.RemoteAddr().String(), address)
	} else if b.errors.IsNetworkUnreachableError(err) {
		b.logger.ConnectNetworkUnreachable(client.RemoteAddr().String(), address)
	} else if b.errors.IsHostUnreachableError(err) {
		b.logger.ConnectHostUnreachable(client.RemoteAddr().String(), address)
	}

	b.sender.SendFailAndClose(client)

	b.logger.UnknownError(client.RemoteAddr().String(), address, err)
	b.logger.ConnectFailed(client.RemoteAddr().String(), address)
}

func (b BaseV4aErrorHandler) HandleV4aConnectIOError(err error, address string, client net.Conn) {
	b.sender.SendFailAndClose(client)

	b.logger.UnknownError(client.RemoteAddr().String(), address, err)
	b.logger.ConnectFailed(client.RemoteAddr().String(), address)
}

func (b BaseV4aErrorHandler) HandleV4aConnectIOErrorWithHost(err error, address string, client net.Conn, host net.Conn) {
	b.sender.SendFailAndClose(client)

	_ = host.Close()

	b.logger.UnknownError(client.RemoteAddr().String(), address, err)
	b.logger.ConnectFailed(client.RemoteAddr().String(), address)
}

func (b BaseV4aErrorHandler) HandleV4aBindIOError(err error, address string, client net.Conn) {
	b.sender.SendFailAndClose(client)

	b.logger.UnknownError(client.RemoteAddr().String(), address, err)
	b.logger.BindFailed(client.RemoteAddr().String(), address)
}

func (b BaseV4aErrorHandler) HandleV4aBindIOErrorWithHost(err error, address string, client net.Conn, host net.Conn) {
	b.sender.SendFailAndClose(client)

	_ = host.Close()

	b.logger.UnknownError(client.RemoteAddr().String(), address, err)
	b.logger.BindFailed(client.RemoteAddr().String(), address)
}
