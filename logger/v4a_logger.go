package logger

import (
	"github.com/rs/zerolog"
)

type SocksV4aLogger interface {
	ConnectRequest(client string, address string)
	ConnectFailed(client string, address string)
	ConnectSuccessful(client string, address string)
	ConnectNotAllowed(client string, address string)
	ConnectNotAllowedByWhitelist(client string, address string)
	ConnectNotAllowedByBlacklist(client string, address string)
	ConnectHostUnreachable(client string, address string)
	ConnectNetworkUnreachable(client string, address string)
	ConnectRefused(client string, address string)
	ConnectTimeout(client string, address string)
	BindRequest(client string, address string)
	BindFailed(client string, address string)
	BindSuccessful(client string, address string)
	BindNotAllowed(client string, address string)
	BindNotAllowedByWhitelist(client string, address string)
	BindNotAllowedByBlacklist(client string, address string)
	BindTimeout(client string, address string)
	Bound(client string, host string)
	TransferFinished(client string, host string)
	ParseError(client string, err error)
	UnknownError(client string, address string, err error)
	AddressParsingError(client string, host string, address string, err error)
	AddressDeterminationError(client string, host string, address string, err error)
	InvalidAddressTypeError(client string, host string, address string)
	BindError(client string, address string, err error)
	ReceiveHostError(client string, address string, err error)
	SendClientError(client string, host string, address string, err error)
}

type BaseSocksV4aLogger struct {
	logger zerolog.Logger
}

func NewBaseSocksV4aLogger(logger zerolog.Logger) (BaseSocksV4aLogger, error) {
	return BaseSocksV4aLogger{
		logger: logger,
	}, nil
}

func (b BaseSocksV4aLogger) ConnectRequest(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Received connect request. ")
}

func (b BaseSocksV4aLogger) ConnectFailed(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Connect failed. ")
}

func (b BaseSocksV4aLogger) ConnectSuccessful(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Connect successful. ")
}

func (b BaseSocksV4aLogger) ConnectNotAllowed(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Connect not allowed due to ruleset. ")
}

func (b BaseSocksV4aLogger) ConnectNotAllowedByWhitelist(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Connect not allowed due to whitelist.")
}

func (b BaseSocksV4aLogger) ConnectNotAllowedByBlacklist(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Connect not allowed due to blacklist.")
}
func (b BaseSocksV4aLogger) ConnectHostUnreachable(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Host unreachable. ")
}

func (b BaseSocksV4aLogger) ConnectTimeout(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Connect failed due to timeout. ")
}

func (b BaseSocksV4aLogger) ConnectNetworkUnreachable(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Network unreachable. ")
}

func (b BaseSocksV4aLogger) ConnectRefused(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Connect refused by host. ")
}

func (b BaseSocksV4aLogger) BindRequest(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Received bind request. ")
}

func (b BaseSocksV4aLogger) BindFailed(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Bind failed. ")
}

func (b BaseSocksV4aLogger) BindSuccessful(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Bind successful. ")
}

func (b BaseSocksV4aLogger) BindNotAllowed(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Bind failed due to ruleset. ")
}

func (b BaseSocksV4aLogger) BindNotAllowedByWhitelist(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Bind not allowed due to whitelist.")
}

func (b BaseSocksV4aLogger) BindNotAllowedByBlacklist(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Bind not allowed due to blacklist.")
}

func (b BaseSocksV4aLogger) BindTimeout(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Bind failed due to timeout. ")
}

func (b BaseSocksV4aLogger) Bound(client string, host string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", host).
		Msg("Bound successfully. ")
}

func (b BaseSocksV4aLogger) TransferFinished(client string, host string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", host).
		Msg("Transfer finished. ")
}

func (b BaseSocksV4aLogger) ParseError(client string, err error) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Err(err).
		Msg("Cannot parse v4a request due to error. ")
}

func (b BaseSocksV4aLogger) UnknownError(client string, address string, err error) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Err(err).
		Msg("Cannot handle v4a request due to error. ")
}

func (b BaseSocksV4aLogger) AddressParsingError(client string, host string, address string, err error) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", host).
		Str("address", address).
		Err(err).
		Msg("Cannot handle v4a request due to address parsing error.")
}

func (b BaseSocksV4aLogger) AddressDeterminationError(client string, host string, address string, err error) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", host).
		Str("address", address).
		Err(err).
		Msg("Cannot handle v4a request due to address determination error.")
}

func (b BaseSocksV4aLogger) InvalidAddressTypeError(client string, host string, address string) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", host).
		Str("address", address).
		Msg("Cannot handle v4a request due to invalid address type.")
}

func (b BaseSocksV4aLogger) BindError(client string, address string, err error) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Err(err).
		Msg("Cannot handle v4a request due to bind error in bind manager.")
}

func (b BaseSocksV4aLogger) ReceiveHostError(client string, address string, err error) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Err(err).
		Msg("Cannot handle v4a request due to bind error in bind manager.")
}

func (b BaseSocksV4aLogger) SendClientError(client string, host string, address string, err error) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", host).
		Str("address", address).
		Err(err).
		Msg("Cannot handle v4a request due to bind error in bind manager.")
}
