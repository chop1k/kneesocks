package v5

import (
	"net"
	v53 "socks/config/v5"
	"socks/logger/v5"
	v52 "socks/protocol/v5"
	"socks/utils"
)

type ErrorHandler struct {
	logger v5.Logger
	sender v52.Sender
	errors utils.ErrorUtils
}

func NewErrorHandler(
	logger v5.Logger,
	sender v52.Sender,
	errors utils.ErrorUtils,
) (ErrorHandler, error) {
	return ErrorHandler{
		logger: logger,
		sender: sender,
		errors: errors,
	}, nil
}

func (b ErrorHandler) HandleDialError(config v53.Config, err error, address string, client net.Conn) {
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

func (b ErrorHandler) HandleConnectIOErrorWithHost(config v53.Config, err error, address string, client net.Conn, host net.Conn) {
	b.sender.SendFailAndClose(config, client)

	_ = host.Close()

	b.logger.Errors.UnknownError(client.RemoteAddr().String(), address, err)
	b.logger.Connect.Failed(client.RemoteAddr().String(), address)
}

func (b ErrorHandler) HandleBindIOError(config v53.Config, err error, address string, client net.Conn) {
	b.sender.SendFailAndClose(config, client)

	b.logger.Errors.UnknownError(client.RemoteAddr().String(), address, err)
	b.logger.Bind.Failed(client.RemoteAddr().String(), address)
}

func (b ErrorHandler) HandleBindIOErrorWithHost(config v53.Config, err error, address string, client net.Conn, host net.Conn) {
	b.sender.SendFailAndClose(config, client)

	_ = host.Close()

	b.logger.Errors.UnknownError(client.RemoteAddr().String(), address, err)
	b.logger.Bind.Failed(client.RemoteAddr().String(), address)
}

func (b ErrorHandler) HandleUdpAssociationError(config v53.Config, err error, address string, client net.Conn) {
	b.sender.SendFailAndClose(config, client)

	b.logger.Errors.UnknownError(client.RemoteAddr().String(), address, err)
	b.logger.Association.Failed(client.RemoteAddr().String(), address)
}

func (b ErrorHandler) HandleAddressParsingError(config v53.Config, err error, address string, client net.Conn, host net.Conn) {
	b.sender.SendFailAndClose(config, client)

	_ = host.Close()

	b.logger.Errors.AddressParsingError(client.RemoteAddr().String(), host.RemoteAddr().String(), address, err)
}

func (b ErrorHandler) HandleAddressDeterminationError(config v53.Config, err error, address string, client net.Conn, host net.Conn) {
	b.sender.SendFailAndClose(config, client)

	_ = host.Close()

	b.logger.Errors.AddressDeterminationError(client.RemoteAddr().String(), host.RemoteAddr().String(), address, err)
}

func (b ErrorHandler) HandleInvalidAddressTypeError(config v53.Config, addressType byte, address string, client net.Conn) {
	b.sender.SendAddressNotSupportedAndClose(config, client)

	b.logger.Errors.InvalidAddressTypeError(client.RemoteAddr().String(), addressType, address)
}

func (b ErrorHandler) HandleBindManagerBindError(config v53.Config, err error, address string, client net.Conn) {
	b.sender.SendFailAndClose(config, client)

	b.logger.Errors.BindError(client.RemoteAddr().String(), address, err)
}

func (b ErrorHandler) HandleBindManagerReceiveHostError(config v53.Config, err error, address string, client net.Conn) {
	b.sender.SendFailAndClose(config, client)

	b.logger.Errors.ReceiveHostError(client.RemoteAddr().String(), address, err)
}

func (b ErrorHandler) HandleBindManagerSendClientError(config v53.Config, err error, address string, client net.Conn, host net.Conn) {
	b.sender.SendFailAndClose(config, client)

	b.logger.Errors.SendClientError(client.RemoteAddr().String(), host.RemoteAddr().String(), address, err)
}

func (b ErrorHandler) HandleReceiveRequestError(config v53.Config, err error, client net.Conn) {
	b.sender.SendFailAndClose(config, client)

	b.logger.Errors.ReceiveRequestError(client.RemoteAddr().String(), err)
}

func (b ErrorHandler) HandlePasswordReceiveRequestError(err error, client net.Conn) {
	_ = client.Close()

	b.logger.Errors.ReceiveRequestError(client.RemoteAddr().String(), err)
}

func (b ErrorHandler) HandleUnknownCommandError(config v53.Config, command byte, address string, client net.Conn) {
	b.sender.SendCommandNotSupportedAndClose(config, client)

	b.logger.Errors.UnknownCommandError(client.RemoteAddr().String(), command, address)
}

func (b ErrorHandler) HandleParseMethodsError(config v53.Config, err error, client net.Conn) {
	b.sender.SendFailAndClose(config, client)

	b.logger.Errors.ParseMethodsError(client.RemoteAddr().String(), err)
}

func (b ErrorHandler) HandleMethodSelectionError(config v53.Config, err error, client net.Conn) {
	b.sender.SendFailAndClose(config, client)

	b.logger.Errors.SelectMethodsError(client.RemoteAddr().String(), err)
}

func (b ErrorHandler) HandlePasswordResponseError(err error, user string, client net.Conn) {
	_ = client.Close()

	b.logger.Errors.PasswordResponseError(client.RemoteAddr().String(), user, err)
}

func (b ErrorHandler) HandleUdpAddressParsingError(config v53.Config, err error, client net.Conn) {
	b.sender.SendFailAndClose(config, client)

	b.logger.Errors.UdpAddressParsingError(client.RemoteAddr().String(), err)
}

func (b ErrorHandler) HandleTransferError(err error, client net.Conn, host net.Conn) {
	_ = client.Close()
	_ = host.Close()

	b.logger.Errors.ConfigError(client.RemoteAddr().String(), err)
}
