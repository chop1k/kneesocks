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
	HandleV4aAddressParsingError(err error, address string, client net.Conn, host net.Conn)
	HandleV4aAddressDeterminationError(err error, address string, client net.Conn, host net.Conn)
	HandleV4aInvalidAddressTypeError(address string, client net.Conn, host net.Conn)
	HandleV4aBindManagerBindError(err error, address string, client net.Conn)
	HandleV4aBindManagerReceiveHostError(err error, address string, client net.Conn)
	HandleV4aBindManagerSendClientError(err error, address string, client net.Conn, host net.Conn)
	HandleV4aChunkParseError(err error, client net.Conn)
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

func (b BaseV4aErrorHandler) HandleV4aAddressParsingError(err error, address string, client net.Conn, host net.Conn) {
	b.sender.SendFailAndClose(client)

	_ = host.Close()

	b.logger.AddressParsingError(client.RemoteAddr().String(), host.RemoteAddr().String(), address, err)
}

func (b BaseV4aErrorHandler) HandleV4aAddressDeterminationError(err error, address string, client net.Conn, host net.Conn) {
	b.sender.SendFailAndClose(client)

	_ = host.Close()

	b.logger.AddressDeterminationError(client.RemoteAddr().String(), host.RemoteAddr().String(), address, err)
}

func (b BaseV4aErrorHandler) HandleV4aInvalidAddressTypeError(address string, client net.Conn, host net.Conn) {
	b.sender.SendFailAndClose(client)

	_ = host.Close()

	b.logger.InvalidAddressTypeError(client.RemoteAddr().String(), host.RemoteAddr().String(), address)
}

func (b BaseV4aErrorHandler) HandleV4aBindManagerBindError(err error, address string, client net.Conn) {
	b.sender.SendFailAndClose(client)

	b.logger.BindError(client.RemoteAddr().String(), address, err)
}

func (b BaseV4aErrorHandler) HandleV4aBindManagerReceiveHostError(err error, address string, client net.Conn) {
	b.sender.SendFailAndClose(client)

	if err == TimeoutError {
		b.logger.BindTimeout(client.RemoteAddr().String(), address)
	} else {
		b.logger.ReceiveHostError(client.RemoteAddr().String(), address, err)
	}
}

func (b BaseV4aErrorHandler) HandleV4aBindManagerSendClientError(err error, address string, client net.Conn, host net.Conn) {
	_ = client.Close()
	_ = host.Close()

	b.logger.SendClientError(client.RemoteAddr().String(), host.RemoteAddr().String(), address, err)
}

func (b BaseV4aErrorHandler) HandleV4aChunkParseError(err error, client net.Conn) {
	b.sender.SendFailAndClose(client)

	b.logger.ParseError(client.RemoteAddr().String(), err)
}
