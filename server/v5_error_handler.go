package server

import (
	"net"
	"os"
	"socks/logger"
	"syscall"
)

type V5ErrorHandler interface {
	HandleV5NetworkError(err error, address string, client net.Conn)
}

type BaseV5ErrorHandler struct {
	logger logger.SocksV5Logger
	sender V5Sender
}

func NewBaseV5ErrorHandler(logger logger.SocksV5Logger, sender V5Sender) (BaseV5ErrorHandler, error) {
	return BaseV5ErrorHandler{
		logger: logger,
		sender: sender,
	}, nil
}

func (b BaseV5ErrorHandler) HandleV5NetworkError(err error, address string, client net.Conn) {
	if b.checkConnectionRefusedError(err) {
		b.sender.SendConnectionRefusedAndClose(client)

		b.logger.ConnectRefused(client.RemoteAddr().String(), address)

		return
	}

	if b.checkNetworkUnreachableError(err) {
		b.sender.SendNetworkUnreachableAndClose(client)

		b.logger.ConnectNetworkUnreachable(client.RemoteAddr().String(), address)

		return
	}

	if b.checkHostUnreachableError(err) {
		b.sender.SendHostUnreachableAndClose(client)

		b.logger.ConnectHostUnreachable(client.RemoteAddr().String(), address)

		return
	}

	b.sender.SendFailAndClose(client)

	//b.errors.UnknownConnectError(client.RemoteAddr().String(), err)
	b.logger.ConnectFailed(client.RemoteAddr().String(), address)
}

func (b BaseV5ErrorHandler) errorToErrno(err error) int {
	opErr, ok := err.(*net.OpError)

	if !ok {
		return -1
	}

	sysErr, ko := opErr.Err.(*os.SyscallError)

	if !ko {
		return -1
	}

	errno, oo := sysErr.Err.(syscall.Errno)

	if !oo {
		return -1
	}

	return int(errno)
}

func (b BaseV5ErrorHandler) checkConnectionRefusedError(err error) bool {
	return b.errorToErrno(err) == 111
}

func (b BaseV5ErrorHandler) checkNetworkUnreachableError(err error) bool {
	return b.errorToErrno(err) == 101
}

func (b BaseV5ErrorHandler) checkHostUnreachableError(err error) bool {
	errno := b.errorToErrno(err)

	return errno == 113 || errno == 112
}
