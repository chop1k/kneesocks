package v4

import (
	"net"
	v43 "socks/config/v4"
	"socks/logger/v4"
	"socks/managers"
	v42 "socks/protocol/v4"
	"socks/utils"
)

type ErrorHandler interface {
	HandleDialError(config v43.Config, err error, address string, client net.Conn)
	HandleConnectIOErrorWithHost(config v43.Config, err error, address string, client net.Conn, host net.Conn)
	HandleBindIOError(config v43.Config, err error, address string, client net.Conn)
	HandleBindIOErrorWithHost(config v43.Config, err error, address string, client net.Conn, host net.Conn)
	HandleAddressParsingError(config v43.Config, err error, address string, client net.Conn, host net.Conn)
	HandleAddressDeterminationError(config v43.Config, err error, address string, client net.Conn, host net.Conn)
	HandleInvalidAddressTypeError(config v43.Config, address string, client net.Conn, host net.Conn)
	HandleBindManagerBindError(config v43.Config, err error, address string, client net.Conn)
	HandleBindManagerReceiveHostError(config v43.Config, err error, address string, client net.Conn)
	HandleBindManagerSendClientError(err error, address string, client net.Conn, host net.Conn)
	HandleChunkParseError(config v43.Config, err error, client net.Conn)
	HandleTransferError(err error, client net.Conn, host net.Conn)
}

type BaseErrorHandler struct {
	logger v4.Logger
	sender v42.Sender
	errors utils.ErrorUtils
}

func NewBaseErrorHandler(
	logger v4.Logger,
	sender v42.Sender,
	errors utils.ErrorUtils,
) (BaseErrorHandler, error) {
	return BaseErrorHandler{
		logger: logger,
		sender: sender,
		errors: errors,
	}, nil
}

func (b BaseErrorHandler) HandleDialError(config v43.Config, err error, address string, client net.Conn) {
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

func (b BaseErrorHandler) HandleConnectIOErrorWithHost(config v43.Config, err error, address string, client net.Conn, host net.Conn) {
	b.sender.SendFailAndClose(config, client)

	_ = host.Close()

	b.logger.Errors.UnknownError(client.RemoteAddr().String(), address, err)
	b.logger.Connect.Failed(client.RemoteAddr().String(), address)
}

func (b BaseErrorHandler) HandleBindIOError(config v43.Config, err error, address string, client net.Conn) {
	b.sender.SendFailAndClose(config, client)

	b.logger.Errors.UnknownError(client.RemoteAddr().String(), address, err)
	b.logger.Bind.Failed(client.RemoteAddr().String(), address)
}

func (b BaseErrorHandler) HandleBindIOErrorWithHost(config v43.Config, err error, address string, client net.Conn, host net.Conn) {
	b.sender.SendFailAndClose(config, client)

	_ = host.Close()

	b.logger.Errors.UnknownError(client.RemoteAddr().String(), address, err)
	b.logger.Bind.Failed(client.RemoteAddr().String(), address)
}

func (b BaseErrorHandler) HandleAddressParsingError(config v43.Config, err error, address string, client net.Conn, host net.Conn) {
	b.sender.SendFailAndClose(config, client)

	_ = host.Close()

	b.logger.Errors.AddressParsingError(client.RemoteAddr().String(), host.RemoteAddr().String(), address, err)
}

func (b BaseErrorHandler) HandleAddressDeterminationError(config v43.Config, err error, address string, client net.Conn, host net.Conn) {
	b.sender.SendFailAndClose(config, client)

	_ = host.Close()

	b.logger.Errors.AddressDeterminationError(client.RemoteAddr().String(), host.RemoteAddr().String(), address, err)
}

func (b BaseErrorHandler) HandleInvalidAddressTypeError(config v43.Config, address string, client net.Conn, host net.Conn) {
	b.sender.SendFailAndClose(config, client)

	_ = host.Close()

	b.logger.Errors.InvalidAddressTypeError(client.RemoteAddr().String(), host.RemoteAddr().String(), address)
}

func (b BaseErrorHandler) HandleBindManagerBindError(config v43.Config, err error, address string, client net.Conn) {
	b.sender.SendFailAndClose(config, client)

	b.logger.Errors.BindError(client.RemoteAddr().String(), address, err)
}

func (b BaseErrorHandler) HandleBindManagerReceiveHostError(config v43.Config, err error, address string, client net.Conn) {
	b.sender.SendFailAndClose(config, client)

	if err == managers.TimeoutError {
		b.logger.Bind.Timeout(client.RemoteAddr().String(), address)
	} else {
		b.logger.Errors.ReceiveHostError(client.RemoteAddr().String(), address, err)
	}
}

func (b BaseErrorHandler) HandleBindManagerSendClientError(err error, address string, client net.Conn, host net.Conn) {
	_ = client.Close()
	_ = host.Close()

	b.logger.Errors.SendClientError(client.RemoteAddr().String(), host.RemoteAddr().String(), address, err)
}

func (b BaseErrorHandler) HandleChunkParseError(config v43.Config, err error, client net.Conn) {
	b.sender.SendFailAndClose(config, client)

	b.logger.Errors.ParseError(client.RemoteAddr().String(), err)
}

func (b BaseErrorHandler) HandleTransferError(err error, client net.Conn, host net.Conn) {
	_ = client.Close()
	_ = host.Close()

	b.logger.Errors.ConfigError(client.RemoteAddr().String(), err)
}
