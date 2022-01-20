package server

import (
	"net"
	"socks/logger"
	"socks/utils"
)

type V4ErrorHandler interface {
	HandleV4DialError(err error, address string, client net.Conn)
	HandleV4ConnectIOError(err error, address string, client net.Conn)
	HandleV4ConnectIOErrorWithHost(err error, address string, client net.Conn, host net.Conn)
	HandleV4BindIOError(err error, address string, client net.Conn)
	HandleV4BindIOErrorWithHost(err error, address string, client net.Conn, host net.Conn)
	HandleV4AddressParsingError(err error, address string, client net.Conn, host net.Conn)
	HandleV4AddressDeterminationError(err error, address string, client net.Conn, host net.Conn)
	HandleV4InvalidAddressTypeError(address string, client net.Conn, host net.Conn)
	HandleV4BindManagerBindError(err error, address string, client net.Conn)
	HandleV4BindManagerReceiveHostError(err error, address string, client net.Conn)
	HandleV4BindManagerSendClientError(err error, address string, client net.Conn, host net.Conn)
	HandleV4ChunkParseError(err error, client net.Conn)
}

type BaseV4ErrorHandler struct {
	logger logger.SocksV4Logger
	sender V4Sender
	errors utils.ErrorUtils
}

func NewBaseV4ErrorHandler(
	logger logger.SocksV4Logger,
	sender V4Sender,
	errors utils.ErrorUtils,
) (BaseV4ErrorHandler, error) {
	return BaseV4ErrorHandler{
		logger: logger,
		sender: sender,
		errors: errors,
	}, nil
}

func (b BaseV4ErrorHandler) HandleV4DialError(err error, address string, client net.Conn) {
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

func (b BaseV4ErrorHandler) HandleV4ConnectIOError(err error, address string, client net.Conn) {
	b.sender.SendFailAndClose(client)

	b.logger.UnknownError(client.RemoteAddr().String(), address, err)
	b.logger.ConnectFailed(client.RemoteAddr().String(), address)
}

func (b BaseV4ErrorHandler) HandleV4ConnectIOErrorWithHost(err error, address string, client net.Conn, host net.Conn) {
	b.sender.SendFailAndClose(client)

	_ = host.Close()

	b.logger.UnknownError(client.RemoteAddr().String(), address, err)
	b.logger.ConnectFailed(client.RemoteAddr().String(), address)
}

func (b BaseV4ErrorHandler) HandleV4BindIOError(err error, address string, client net.Conn) {
	b.sender.SendFailAndClose(client)

	b.logger.UnknownError(client.RemoteAddr().String(), address, err)
	b.logger.BindFailed(client.RemoteAddr().String(), address)
}

func (b BaseV4ErrorHandler) HandleV4BindIOErrorWithHost(err error, address string, client net.Conn, host net.Conn) {
	b.sender.SendFailAndClose(client)

	_ = host.Close()

	b.logger.UnknownError(client.RemoteAddr().String(), address, err)
	b.logger.BindFailed(client.RemoteAddr().String(), address)
}

func (b BaseV4ErrorHandler) HandleV4AddressParsingError(err error, address string, client net.Conn, host net.Conn) {
	b.sender.SendFailAndClose(client)

	_ = host.Close()

	b.logger.AddressParsingError(client.RemoteAddr().String(), host.RemoteAddr().String(), address, err)
}

func (b BaseV4ErrorHandler) HandleV4AddressDeterminationError(err error, address string, client net.Conn, host net.Conn) {
	b.sender.SendFailAndClose(client)

	_ = host.Close()

	b.logger.AddressDeterminationError(client.RemoteAddr().String(), host.RemoteAddr().String(), address, err)
}

func (b BaseV4ErrorHandler) HandleV4InvalidAddressTypeError(address string, client net.Conn, host net.Conn) {
	b.sender.SendFailAndClose(client)

	_ = host.Close()

	b.logger.InvalidAddressTypeError(client.RemoteAddr().String(), host.RemoteAddr().String(), address)
}

func (b BaseV4ErrorHandler) HandleV4BindManagerBindError(err error, address string, client net.Conn) {
	b.sender.SendFailAndClose(client)

	b.logger.BindError(client.RemoteAddr().String(), address, err)
}

func (b BaseV4ErrorHandler) HandleV4BindManagerReceiveHostError(err error, address string, client net.Conn) {
	b.sender.SendFailAndClose(client)

	if err == TimeoutError {
		b.logger.BindTimeout(client.RemoteAddr().String(), address)
	} else {
		b.logger.ReceiveHostError(client.RemoteAddr().String(), address, err)
	}
}

func (b BaseV4ErrorHandler) HandleV4BindManagerSendClientError(err error, address string, client net.Conn, host net.Conn) {
	_ = client.Close()
	_ = host.Close()

	b.logger.SendClientError(client.RemoteAddr().String(), host.RemoteAddr().String(), address, err)
}

func (b BaseV4ErrorHandler) HandleV4ChunkParseError(err error, client net.Conn) {
	b.sender.SendFailAndClose(client)

	b.logger.ParseError(client.RemoteAddr().String(), err)
}
