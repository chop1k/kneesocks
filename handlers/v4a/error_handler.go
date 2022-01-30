package v4a

import (
	"net"
	"socks/logger"
	"socks/managers"
	"socks/utils"
)

type ErrorHandler interface {
	HandleDialError(err error, address string, client net.Conn)
	HandleConnectIOError(err error, address string, client net.Conn)
	HandleConnectIOErrorWithHost(err error, address string, client net.Conn, host net.Conn)
	HandleBindIOError(err error, address string, client net.Conn)
	HandleBindIOErrorWithHost(err error, address string, client net.Conn, host net.Conn)
	HandleAddressParsingError(err error, address string, client net.Conn, host net.Conn)
	HandleAddressDeterminationError(err error, address string, client net.Conn, host net.Conn)
	HandleInvalidAddressTypeError(address string, client net.Conn, host net.Conn)
	HandleBindManagerBindError(err error, address string, client net.Conn)
	HandleBindManagerReceiveHostError(err error, address string, client net.Conn)
	HandleBindManagerSendClientError(err error, address string, client net.Conn, host net.Conn)
	HandleChunkParseError(err error, client net.Conn)
}

type BaseErrorHandler struct {
	logger logger.SocksV4aLogger
	sender Sender
	errors utils.ErrorUtils
}

func NewBaseErrorHandler(
	logger logger.SocksV4aLogger,
	sender Sender,
	errors utils.ErrorUtils,
) (BaseErrorHandler, error) {
	return BaseErrorHandler{
		logger: logger,
		sender: sender,
		errors: errors,
	}, nil
}

func (b BaseErrorHandler) HandleDialError(err error, address string, client net.Conn) {
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

func (b BaseErrorHandler) HandleConnectIOError(err error, address string, client net.Conn) {
	b.sender.SendFailAndClose(client)

	b.logger.UnknownError(client.RemoteAddr().String(), address, err)
	b.logger.ConnectFailed(client.RemoteAddr().String(), address)
}

func (b BaseErrorHandler) HandleConnectIOErrorWithHost(err error, address string, client net.Conn, host net.Conn) {
	b.sender.SendFailAndClose(client)

	_ = host.Close()

	b.logger.UnknownError(client.RemoteAddr().String(), address, err)
	b.logger.ConnectFailed(client.RemoteAddr().String(), address)
}

func (b BaseErrorHandler) HandleBindIOError(err error, address string, client net.Conn) {
	b.sender.SendFailAndClose(client)

	b.logger.UnknownError(client.RemoteAddr().String(), address, err)
	b.logger.BindFailed(client.RemoteAddr().String(), address)
}

func (b BaseErrorHandler) HandleBindIOErrorWithHost(err error, address string, client net.Conn, host net.Conn) {
	b.sender.SendFailAndClose(client)

	_ = host.Close()

	b.logger.UnknownError(client.RemoteAddr().String(), address, err)
	b.logger.BindFailed(client.RemoteAddr().String(), address)
}

func (b BaseErrorHandler) HandleAddressParsingError(err error, address string, client net.Conn, host net.Conn) {
	b.sender.SendFailAndClose(client)

	_ = host.Close()

	b.logger.AddressParsingError(client.RemoteAddr().String(), host.RemoteAddr().String(), address, err)
}

func (b BaseErrorHandler) HandleAddressDeterminationError(err error, address string, client net.Conn, host net.Conn) {
	b.sender.SendFailAndClose(client)

	_ = host.Close()

	b.logger.AddressDeterminationError(client.RemoteAddr().String(), host.RemoteAddr().String(), address, err)
}

func (b BaseErrorHandler) HandleInvalidAddressTypeError(address string, client net.Conn, host net.Conn) {
	b.sender.SendFailAndClose(client)

	_ = host.Close()

	b.logger.InvalidAddressTypeError(client.RemoteAddr().String(), host.RemoteAddr().String(), address)
}

func (b BaseErrorHandler) HandleBindManagerBindError(err error, address string, client net.Conn) {
	b.sender.SendFailAndClose(client)

	b.logger.BindError(client.RemoteAddr().String(), address, err)
}

func (b BaseErrorHandler) HandleBindManagerReceiveHostError(err error, address string, client net.Conn) {
	b.sender.SendFailAndClose(client)

	if err == managers.TimeoutError {
		b.logger.BindTimeout(client.RemoteAddr().String(), address)
	} else {
		b.logger.ReceiveHostError(client.RemoteAddr().String(), address, err)
	}
}

func (b BaseErrorHandler) HandleBindManagerSendClientError(err error, address string, client net.Conn, host net.Conn) {
	_ = client.Close()
	_ = host.Close()

	b.logger.SendClientError(client.RemoteAddr().String(), host.RemoteAddr().String(), address, err)
}

func (b BaseErrorHandler) HandleChunkParseError(err error, client net.Conn) {
	b.sender.SendFailAndClose(client)

	b.logger.ParseError(client.RemoteAddr().String(), err)
}
