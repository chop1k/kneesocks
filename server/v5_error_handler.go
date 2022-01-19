package server

import (
	"net"
	"socks/logger"
	"socks/utils"
)

type V5ErrorHandler interface {
	HandleV5DialError(err error, address string, client net.Conn)
	HandleV5ConnectIOError(err error, address string, client net.Conn)
	HandleV5ConnectIOErrorWithHost(err error, address string, client net.Conn, host net.Conn)
	HandleV5BindIOError(err error, address string, client net.Conn)
	HandleV5BindIOErrorWithHost(err error, address string, client net.Conn, host net.Conn)
	HandleV5UdpAssociationError(err error, address string, client net.Conn)
}

type BaseV5ErrorHandler struct {
	logger logger.SocksV5Logger
	sender V5Sender
	errors utils.ErrorUtils
}

func NewBaseV5ErrorHandler(
	logger logger.SocksV5Logger,
	sender V5Sender,
	errors utils.ErrorUtils,
) (BaseV5ErrorHandler, error) {
	return BaseV5ErrorHandler{
		logger: logger,
		sender: sender,
		errors: errors,
	}, nil
}

func (b BaseV5ErrorHandler) HandleV5DialError(err error, address string, client net.Conn) {
	if b.errors.IsConnectionRefusedError(err) {
		b.sender.SendConnectionRefusedAndClose(client)

		b.logger.ConnectRefused(client.RemoteAddr().String(), address)

		return
	}

	if b.errors.IsNetworkUnreachableError(err) {
		b.sender.SendNetworkUnreachableAndClose(client)

		b.logger.ConnectNetworkUnreachable(client.RemoteAddr().String(), address)

		return
	}

	if b.errors.IsHostUnreachableError(err) {
		b.sender.SendHostUnreachableAndClose(client)

		b.logger.ConnectHostUnreachable(client.RemoteAddr().String(), address)

		return
	}

	b.sender.SendFailAndClose(client)

	b.logger.UnknownError(client.RemoteAddr().String(), address, err)
	b.logger.ConnectFailed(client.RemoteAddr().String(), address)
}

func (b BaseV5ErrorHandler) HandleV5ConnectIOError(err error, address string, client net.Conn) {
	b.sender.SendFailAndClose(client)

	b.logger.UnknownError(client.RemoteAddr().String(), address, err)
	b.logger.ConnectFailed(client.RemoteAddr().String(), address)
}

func (b BaseV5ErrorHandler) HandleV5ConnectIOErrorWithHost(err error, address string, client net.Conn, host net.Conn) {
	b.sender.SendFailAndClose(client)

	_ = host.Close()

	b.logger.UnknownError(client.RemoteAddr().String(), address, err)
	b.logger.ConnectFailed(client.RemoteAddr().String(), address)
}

func (b BaseV5ErrorHandler) HandleV5BindIOError(err error, address string, client net.Conn) {
	b.sender.SendFailAndClose(client)

	b.logger.UnknownError(client.RemoteAddr().String(), address, err)
	b.logger.BindFailed(client.RemoteAddr().String(), address)
}

func (b BaseV5ErrorHandler) HandleV5BindIOErrorWithHost(err error, address string, client net.Conn, host net.Conn) {
	b.sender.SendFailAndClose(client)

	_ = host.Close()

	b.logger.UnknownError(client.RemoteAddr().String(), address, err)
	b.logger.BindFailed(client.RemoteAddr().String(), address)
}

func (b BaseV5ErrorHandler) HandleV5UdpAssociationError(err error, address string, client net.Conn) {
	b.sender.SendFailAndClose(client)

	b.logger.UnknownError(client.RemoteAddr().String(), address, err)
	b.logger.UdpAssociationFailed(client.RemoteAddr().String(), address)
}
