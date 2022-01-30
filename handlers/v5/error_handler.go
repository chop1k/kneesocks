package v5

import (
	"net"
	"socks/logger"
	"socks/utils"
)

type ErrorHandler interface {
	HandleDialError(err error, address string, client net.Conn)
	HandleConnectIOError(err error, address string, client net.Conn)
	HandleConnectIOErrorWithHost(err error, address string, client net.Conn, host net.Conn)
	HandleBindIOError(err error, address string, client net.Conn)
	HandleBindIOErrorWithHost(err error, address string, client net.Conn, host net.Conn)
	HandleUdpAssociationError(err error, address string, client net.Conn)
	HandleAddressParsingError(err error, address string, client net.Conn, host net.Conn)
	HandleAddressDeterminationError(err error, address string, client net.Conn, host net.Conn)
	HandleInvalidAddressTypeError(addressType byte, address string, client net.Conn)
	HandleBindManagerBindError(err error, address string, client net.Conn)
	HandleBindManagerReceiveHostError(err error, address string, client net.Conn)
	HandleBindManagerSendClientError(err error, address string, client net.Conn, host net.Conn)
	HandleReceiveRequestError(err error, client net.Conn)
	HandlePasswordReceiveRequestError(err error, client net.Conn)
	HandleUnknownCommandError(command byte, address string, client net.Conn)
	HandleParseMethodsError(err error, client net.Conn)
	HandleSelectMethodsError(err error, client net.Conn)
	HandlePasswordResponseError(err error, user string, client net.Conn)
	HandleUdpAddressParsingError(err error, client net.Conn)
	HandleIPv4AddressNotAllowed(address string, client net.Conn)
	HandleDomainAddressNotAllowed(address string, client net.Conn)
	HandleIPv6AddressNotAllowed(address string, client net.Conn)
}

type BaseErrorHandler struct {
	logger logger.SocksV5Logger
	sender Sender
	errors utils.ErrorUtils
}

func NewBaseErrorHandler(
	logger logger.SocksV5Logger,
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

func (b BaseErrorHandler) HandleUdpAssociationError(err error, address string, client net.Conn) {
	b.sender.SendFailAndClose(client)

	b.logger.UnknownError(client.RemoteAddr().String(), address, err)
	b.logger.UdpAssociationFailed(client.RemoteAddr().String(), address)
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

func (b BaseErrorHandler) HandleInvalidAddressTypeError(addressType byte, address string, client net.Conn) {
	b.sender.SendAddressNotSupportedAndClose(client)

	b.logger.InvalidAddressTypeError(client.RemoteAddr().String(), addressType, address)
}

func (b BaseErrorHandler) HandleBindManagerBindError(err error, address string, client net.Conn) {
	b.sender.SendFailAndClose(client)

	b.logger.BindError(client.RemoteAddr().String(), address, err)
}

func (b BaseErrorHandler) HandleBindManagerReceiveHostError(err error, address string, client net.Conn) {
	b.sender.SendFailAndClose(client)

	b.logger.ReceiveHostError(client.RemoteAddr().String(), address, err)
}

func (b BaseErrorHandler) HandleBindManagerSendClientError(err error, address string, client net.Conn, host net.Conn) {
	b.sender.SendFailAndClose(client)

	b.logger.SendClientError(client.RemoteAddr().String(), host.RemoteAddr().String(), address, err)
}

func (b BaseErrorHandler) HandleReceiveRequestError(err error, client net.Conn) {
	b.sender.SendFailAndClose(client)

	b.logger.ReceiveRequestError(client.RemoteAddr().String(), err)
}

func (b BaseErrorHandler) HandlePasswordReceiveRequestError(err error, client net.Conn) {
	_ = client.Close()

	b.logger.ReceiveRequestError(client.RemoteAddr().String(), err)
}

func (b BaseErrorHandler) HandleUnknownCommandError(command byte, address string, client net.Conn) {
	b.sender.SendCommandNotSupportedAndClose(client)

	b.logger.UnknownCommandError(client.RemoteAddr().String(), command, address)
}

func (b BaseErrorHandler) HandleParseMethodsError(err error, client net.Conn) {
	b.sender.SendFailAndClose(client)

	b.logger.ParseMethodsError(client.RemoteAddr().String(), err)
}

func (b BaseErrorHandler) HandleSelectMethodsError(err error, client net.Conn) {
	b.sender.SendFailAndClose(client)

	b.logger.SelectMethodsError(client.RemoteAddr().String(), err)
}

func (b BaseErrorHandler) HandlePasswordResponseError(err error, user string, client net.Conn) {
	b.sender.SendFailAndClose(client)

	b.logger.PasswordResponseError(client.RemoteAddr().String(), user, err)
}

func (b BaseErrorHandler) HandleUdpAddressParsingError(err error, client net.Conn) {
	b.sender.SendFailAndClose(client)

	b.logger.UdpAddressParsingError(client.RemoteAddr().String(), err)
}

func (b BaseErrorHandler) HandleIPv4AddressNotAllowed(address string, client net.Conn) {
	b.sender.SendAddressNotSupportedAndClose(client)

	b.logger.IPv4AddressNotAllowed(client.RemoteAddr().String(), address)
}

func (b BaseErrorHandler) HandleDomainAddressNotAllowed(address string, client net.Conn) {
	b.sender.SendAddressNotSupportedAndClose(client)

	b.logger.DomainAddressNotAllowed(client.RemoteAddr().String(), address)
}

func (b BaseErrorHandler) HandleIPv6AddressNotAllowed(address string, client net.Conn) {
	b.sender.SendAddressNotSupportedAndClose(client)

	b.logger.IPv6AddressNotAllowed(client.RemoteAddr().String(), address)
}
