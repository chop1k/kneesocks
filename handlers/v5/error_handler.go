package v5

import (
	"net"
	v53 "socks/config/v5"
	"socks/logger/v5"
	v52 "socks/protocol/v5"
	"socks/utils"
)

type ErrorHandler interface {
	HandleDialError(config v53.Config, err error, address string, client net.Conn)
	HandleConnectIOErrorWithHost(config v53.Config, err error, address string, client net.Conn, host net.Conn)
	HandleBindIOError(config v53.Config, err error, address string, client net.Conn)
	HandleBindIOErrorWithHost(config v53.Config, err error, address string, client net.Conn, host net.Conn)
	HandleUdpAssociationError(config v53.Config, err error, address string, client net.Conn)
	HandleAddressParsingError(config v53.Config, err error, address string, client net.Conn, host net.Conn)
	HandleAddressDeterminationError(config v53.Config, err error, address string, client net.Conn, host net.Conn)
	HandleInvalidAddressTypeError(config v53.Config, addressType byte, address string, client net.Conn)
	HandleBindManagerBindError(config v53.Config, err error, address string, client net.Conn)
	HandleBindManagerReceiveHostError(config v53.Config, err error, address string, client net.Conn)
	HandleBindManagerSendClientError(config v53.Config, err error, address string, client net.Conn, host net.Conn)
	HandleReceiveRequestError(config v53.Config, err error, client net.Conn)
	HandlePasswordReceiveRequestError(err error, client net.Conn)
	HandleUnknownCommandError(config v53.Config, command byte, address string, client net.Conn)
	HandleParseMethodsError(config v53.Config, err error, client net.Conn)
	HandleMethodSelectionError(config v53.Config, err error, client net.Conn)
	HandlePasswordResponseError(err error, user string, client net.Conn)
	HandleUdpAddressParsingError(config v53.Config, err error, client net.Conn)
	HandleTransferError(err error, client net.Conn, host net.Conn)
}

type BaseErrorHandler struct {
	logger v5.Logger
	sender v52.Sender
	errors utils.ErrorUtils
}

func NewBaseErrorHandler(
	logger v5.Logger,
	sender v52.Sender,
	errors utils.ErrorUtils,
) (BaseErrorHandler, error) {
	return BaseErrorHandler{
		logger: logger,
		sender: sender,
		errors: errors,
	}, nil
}

func (b BaseErrorHandler) HandleDialError(config v53.Config, err error, address string, client net.Conn) {
	if b.errors.IsConnectionRefusedError(err) {
		b.sender.SendConnectionRefusedAndClose(config, client)

		b.logger.Connect.Refused(client.RemoteAddr().String(), address)

		return
	}

	if b.errors.IsNetworkUnreachableError(err) {
		b.sender.SendNetworkUnreachableAndClose(config, client)

		b.logger.Connect.NetworkUnreachable(client.RemoteAddr().String(), address)

		return
	}

	if b.errors.IsHostUnreachableError(err) {
		b.sender.SendHostUnreachableAndClose(config, client)

		b.logger.Connect.HostUnreachable(client.RemoteAddr().String(), address)

		return
	}

	b.sender.SendFailAndClose(config, client)

	b.logger.Errors.UnknownError(client.RemoteAddr().String(), address, err)
	b.logger.Connect.Failed(client.RemoteAddr().String(), address)
}

func (b BaseErrorHandler) HandleConnectIOErrorWithHost(config v53.Config, err error, address string, client net.Conn, host net.Conn) {
	b.sender.SendFailAndClose(config, client)

	_ = host.Close()

	b.logger.Errors.UnknownError(client.RemoteAddr().String(), address, err)
	b.logger.Connect.Failed(client.RemoteAddr().String(), address)
}

func (b BaseErrorHandler) HandleBindIOError(config v53.Config, err error, address string, client net.Conn) {
	b.sender.SendFailAndClose(config, client)

	b.logger.Errors.UnknownError(client.RemoteAddr().String(), address, err)
	b.logger.Bind.Failed(client.RemoteAddr().String(), address)
}

func (b BaseErrorHandler) HandleBindIOErrorWithHost(config v53.Config, err error, address string, client net.Conn, host net.Conn) {
	b.sender.SendFailAndClose(config, client)

	_ = host.Close()

	b.logger.Errors.UnknownError(client.RemoteAddr().String(), address, err)
	b.logger.Bind.Failed(client.RemoteAddr().String(), address)
}

func (b BaseErrorHandler) HandleUdpAssociationError(config v53.Config, err error, address string, client net.Conn) {
	b.sender.SendFailAndClose(config, client)

	b.logger.Errors.UnknownError(client.RemoteAddr().String(), address, err)
	b.logger.Association.Failed(client.RemoteAddr().String(), address)
}

func (b BaseErrorHandler) HandleAddressParsingError(config v53.Config, err error, address string, client net.Conn, host net.Conn) {
	b.sender.SendFailAndClose(config, client)

	_ = host.Close()

	b.logger.Errors.AddressParsingError(client.RemoteAddr().String(), host.RemoteAddr().String(), address, err)
}

func (b BaseErrorHandler) HandleAddressDeterminationError(config v53.Config, err error, address string, client net.Conn, host net.Conn) {
	b.sender.SendFailAndClose(config, client)

	_ = host.Close()

	b.logger.Errors.AddressDeterminationError(client.RemoteAddr().String(), host.RemoteAddr().String(), address, err)
}

func (b BaseErrorHandler) HandleInvalidAddressTypeError(config v53.Config, addressType byte, address string, client net.Conn) {
	b.sender.SendAddressNotSupportedAndClose(config, client)

	b.logger.Errors.InvalidAddressTypeError(client.RemoteAddr().String(), addressType, address)
}

func (b BaseErrorHandler) HandleBindManagerBindError(config v53.Config, err error, address string, client net.Conn) {
	b.sender.SendFailAndClose(config, client)

	b.logger.Errors.BindError(client.RemoteAddr().String(), address, err)
}

func (b BaseErrorHandler) HandleBindManagerReceiveHostError(config v53.Config, err error, address string, client net.Conn) {
	b.sender.SendFailAndClose(config, client)

	b.logger.Errors.ReceiveHostError(client.RemoteAddr().String(), address, err)
}

func (b BaseErrorHandler) HandleBindManagerSendClientError(config v53.Config, err error, address string, client net.Conn, host net.Conn) {
	b.sender.SendFailAndClose(config, client)

	b.logger.Errors.SendClientError(client.RemoteAddr().String(), host.RemoteAddr().String(), address, err)
}

func (b BaseErrorHandler) HandleReceiveRequestError(config v53.Config, err error, client net.Conn) {
	b.sender.SendFailAndClose(config, client)

	b.logger.Errors.ReceiveRequestError(client.RemoteAddr().String(), err)
}

func (b BaseErrorHandler) HandlePasswordReceiveRequestError(err error, client net.Conn) {
	_ = client.Close()

	b.logger.Errors.ReceiveRequestError(client.RemoteAddr().String(), err)
}

func (b BaseErrorHandler) HandleUnknownCommandError(config v53.Config, command byte, address string, client net.Conn) {
	b.sender.SendCommandNotSupportedAndClose(config, client)

	b.logger.Errors.UnknownCommandError(client.RemoteAddr().String(), command, address)
}

func (b BaseErrorHandler) HandleParseMethodsError(config v53.Config, err error, client net.Conn) {
	b.sender.SendFailAndClose(config, client)

	b.logger.Errors.ParseMethodsError(client.RemoteAddr().String(), err)
}

func (b BaseErrorHandler) HandleMethodSelectionError(config v53.Config, err error, client net.Conn) {
	b.sender.SendFailAndClose(config, client)

	b.logger.Errors.SelectMethodsError(client.RemoteAddr().String(), err)
}

func (b BaseErrorHandler) HandlePasswordResponseError(err error, user string, client net.Conn) {
	_ = client.Close()

	b.logger.Errors.PasswordResponseError(client.RemoteAddr().String(), user, err)
}

func (b BaseErrorHandler) HandleUdpAddressParsingError(config v53.Config, err error, client net.Conn) {
	b.sender.SendFailAndClose(config, client)

	b.logger.Errors.UdpAddressParsingError(client.RemoteAddr().String(), err)
}

func (b BaseErrorHandler) HandleTransferError(err error, client net.Conn, host net.Conn) {
	_ = client.Close()
	_ = host.Close()

	b.logger.Errors.ConfigError(client.RemoteAddr().String(), err)
}
