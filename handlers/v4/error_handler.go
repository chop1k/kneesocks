package v4

import (
	"net"
	v43 "socks/config/v4"
	"socks/logger/v4"
	"socks/managers"
	v42 "socks/protocol/v4"
	"socks/utils"
)

type ErrorHandler struct {
	logger v4.Logger
	sender v42.Sender
	errors utils.ErrorUtils
}

func NewErrorHandler(
	logger v4.Logger,
	sender v42.Sender,
	errors utils.ErrorUtils,
) (ErrorHandler, error) {
	return ErrorHandler{
		logger: logger,
		sender: sender,
		errors: errors,
	}, nil
}

func (b ErrorHandler) HandleDialError(config v43.Config, err error, address string, client net.Conn) {
	if b.errors.IsConnectionRefusedError(err) {
		b.logger.Connect.Refused(client.RemoteAddr().String(), address)
	} else if b.errors.IsNetworkUnreachableError(err) {
		b.logger.Connect.NetworkUnreachable(client.RemoteAddr().String(), address)
	} else if b.errors.IsHostUnreachableError(err) {
		b.logger.Connect.HostUnreachable(client.RemoteAddr().String(), address)
	}

	b.sender.SendFailAndClose(config, client)

	b.logger.Errors.UnknownError(client.RemoteAddr().String(), address, err)
	b.logger.Connect.Failed(client.RemoteAddr().String(), address)
}

func (b ErrorHandler) HandleConnectIOErrorWithHost(config v43.Config, err error, address string, client net.Conn, host net.Conn) {
	b.sender.SendFailAndClose(config, client)

	_ = host.Close()

	b.logger.Errors.UnknownError(client.RemoteAddr().String(), address, err)
	b.logger.Connect.Failed(client.RemoteAddr().String(), address)
}

func (b ErrorHandler) HandleBindIOError(config v43.Config, err error, address string, client net.Conn) {
	b.sender.SendFailAndClose(config, client)

	b.logger.Errors.UnknownError(client.RemoteAddr().String(), address, err)
	b.logger.Bind.Failed(client.RemoteAddr().String(), address)
}

func (b ErrorHandler) HandleBindIOErrorWithHost(config v43.Config, err error, address string, client net.Conn, host net.Conn) {
	b.sender.SendFailAndClose(config, client)

	_ = host.Close()

	b.logger.Errors.UnknownError(client.RemoteAddr().String(), address, err)
	b.logger.Bind.Failed(client.RemoteAddr().String(), address)
}

func (b ErrorHandler) HandleAddressParsingError(config v43.Config, err error, address string, client net.Conn, host net.Conn) {
	b.sender.SendFailAndClose(config, client)

	_ = host.Close()

	b.logger.Errors.AddressParsingError(client.RemoteAddr().String(), host.RemoteAddr().String(), address, err)
}

func (b ErrorHandler) HandleAddressDeterminationError(config v43.Config, err error, address string, client net.Conn, host net.Conn) {
	b.sender.SendFailAndClose(config, client)

	_ = host.Close()

	b.logger.Errors.AddressDeterminationError(client.RemoteAddr().String(), host.RemoteAddr().String(), address, err)
}

func (b ErrorHandler) HandleInvalidAddressTypeError(config v43.Config, address string, client net.Conn, host net.Conn) {
	b.sender.SendFailAndClose(config, client)

	_ = host.Close()

	b.logger.Errors.InvalidAddressTypeError(client.RemoteAddr().String(), host.RemoteAddr().String(), address)
}

func (b ErrorHandler) HandleBindManagerBindError(config v43.Config, err error, address string, client net.Conn) {
	b.sender.SendFailAndClose(config, client)

	b.logger.Errors.BindError(client.RemoteAddr().String(), address, err)
}

func (b ErrorHandler) HandleBindManagerReceiveHostError(config v43.Config, err error, address string, client net.Conn) {
	b.sender.SendFailAndClose(config, client)

	if err == managers.TimeoutError {
		b.logger.Bind.Timeout(client.RemoteAddr().String(), address)
	} else {
		b.logger.Errors.ReceiveHostError(client.RemoteAddr().String(), address, err)
	}
}

func (b ErrorHandler) HandleBindManagerSendClientError(err error, address string, client net.Conn, host net.Conn) {
	_ = client.Close()
	_ = host.Close()

	b.logger.Errors.SendClientError(client.RemoteAddr().String(), host.RemoteAddr().String(), address, err)
}

func (b ErrorHandler) HandleChunkParseError(config v43.Config, err error, client net.Conn) {
	b.sender.SendFailAndClose(config, client)

	b.logger.Errors.ParseError(client.RemoteAddr().String(), err)
}

func (b ErrorHandler) HandleTransferError(err error, client net.Conn, host net.Conn) {
	_ = client.Close()
	_ = host.Close()

	b.logger.Errors.ConfigError(client.RemoteAddr().String(), err)
}
