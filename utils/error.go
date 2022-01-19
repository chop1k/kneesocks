package utils

import (
	"net"
	"os"
	"syscall"
)

type ErrorUtils struct {
}

func NewErrorUtils() (ErrorUtils, error) {
	return ErrorUtils{}, nil
}

func (u ErrorUtils) errorToErrno(err error) int {
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

func (u ErrorUtils) IsConnectionRefusedError(err error) bool {
	return u.errorToErrno(err) == 111
}

func (u ErrorUtils) IsNetworkUnreachableError(err error) bool {
	return u.errorToErrno(err) == 101
}

func (u ErrorUtils) IsHostUnreachableError(err error) bool {
	errno := u.errorToErrno(err)

	return errno == 113 || errno == 112
}
