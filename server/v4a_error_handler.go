package server

import (
	"net"
	"os"
	"socks/logger"
	"syscall"
)

type V4aErrorHandler interface {
	HandleV4aNetworkError(err error, address string, client net.Conn)
}

type BaseV4aErrorHandler struct {
	logger logger.SocksV4aLogger
	sender V4aSender
}

func NewBaseV4aErrorHandler(logger logger.SocksV4aLogger, sender V4aSender) (BaseV4aErrorHandler, error) {
	return BaseV4aErrorHandler{
		logger: logger,
		sender: sender,
	}, nil
}

func (b BaseV4aErrorHandler) HandleV4aNetworkError(err error, address string, client net.Conn) {
	if b.checkConnectionRefusedError(err) {
		b.logger.ConnectRefused(client.RemoteAddr().String(), address)
	} else if b.checkNetworkUnreachableError(err) {
		b.logger.ConnectNetworkUnreachable(client.RemoteAddr().String(), address)
	} else if b.checkHostUnreachableError(err) {
		b.logger.ConnectHostUnreachable(client.RemoteAddr().String(), address)
	}

	b.sender.SendFailAndClose(client)

	//b.errors.UnknownConnectError(client.RemoteAddr().String(), err)
	b.logger.ConnectFailed(client.RemoteAddr().String(), address)
}

func (b BaseV4aErrorHandler) errorToErrno(err error) int {
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

func (b BaseV4aErrorHandler) checkConnectionRefusedError(err error) bool {
	return b.errorToErrno(err) == 111
}

func (b BaseV4aErrorHandler) checkNetworkUnreachableError(err error) bool {
	return b.errorToErrno(err) == 101
}

func (b BaseV4aErrorHandler) checkHostUnreachableError(err error) bool {
	errno := b.errorToErrno(err)

	return errno == 113 || errno == 112
}
