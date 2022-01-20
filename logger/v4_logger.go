package logger

import (
	"github.com/rs/zerolog"
)

type SocksV4Logger interface {
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

type BaseSocksV4Logger struct {
	logger zerolog.Logger
}

func NewBaseSocksV4Logger(logger zerolog.Logger) (BaseSocksV4Logger, error) {
	return BaseSocksV4Logger{
		logger: logger,
	}, nil
}

func (b BaseSocksV4Logger) ConnectRequest(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Received connect request.")
}

func (b BaseSocksV4Logger) ConnectFailed(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Connect failed.")
}

func (b BaseSocksV4Logger) ConnectSuccessful(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Connect successful.")
}

func (b BaseSocksV4Logger) ConnectNotAllowed(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Connect not allowed due to ruleset.")
}

func (b BaseSocksV4Logger) ConnectNotAllowedByWhitelist(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Connect not allowed due to whitelist.")
}

func (b BaseSocksV4Logger) ConnectNotAllowedByBlacklist(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Connect not allowed due to blacklist.")
}

func (b BaseSocksV4Logger) ConnectHostUnreachable(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Host unreachable.")
}

func (b BaseSocksV4Logger) ConnectNetworkUnreachable(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Network unreachable.")
}

func (b BaseSocksV4Logger) ConnectRefused(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Connect refused by host.")
}

func (b BaseSocksV4Logger) ConnectTimeout(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Connect failed due to timeout.")
}

func (b BaseSocksV4Logger) BindRequest(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Received bind request.")
}

func (b BaseSocksV4Logger) BindFailed(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Bind failed.")
}

func (b BaseSocksV4Logger) BindSuccessful(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Bind successful.")
}

func (b BaseSocksV4Logger) BindNotAllowed(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Bind failed due to ruleset.")
}

func (b BaseSocksV4Logger) BindNotAllowedByWhitelist(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Bind not allowed due to whitelist.")
}

func (b BaseSocksV4Logger) BindNotAllowedByBlacklist(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Bind not allowed due to blacklist.")
}

func (b BaseSocksV4Logger) BindTimeout(client string, address string) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Bind failed due to timeout.")
}

func (b BaseSocksV4Logger) Bound(client string, host string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", host).
		Msg("Bound successfully.")
}

func (b BaseSocksV4Logger) TransferFinished(client string, host string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", host).
		Msg("Transfer finished.")
}

func (b BaseSocksV4Logger) ParseError(client string, err error) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Err(err).
		Msg("Cannot parse v4 request due to error.")
}

func (b BaseSocksV4Logger) UnknownError(client string, address string, err error) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Err(err).
		Msg("Cannot handle v4 request due to error.")
}

func (b BaseSocksV4Logger) AddressParsingError(client string, host string, address string, err error) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", host).
		Str("address", address).
		Err(err).
		Msg("Cannot handle v4 request due to address parsing error.")
}

func (b BaseSocksV4Logger) AddressDeterminationError(client string, host string, address string, err error) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", host).
		Str("address", address).
		Err(err).
		Msg("Cannot handle v4 request due to address determination error.")
}

func (b BaseSocksV4Logger) InvalidAddressTypeError(client string, host string, address string) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", host).
		Str("address", address).
		Msg("Cannot handle v4 request due to invalid address type.")
}

func (b BaseSocksV4Logger) BindError(client string, address string, err error) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Err(err).
		Msg("Cannot handle v4 request due to bind error in bind manager.")
}

func (b BaseSocksV4Logger) ReceiveHostError(client string, address string, err error) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Err(err).
		Msg("Cannot handle v4 request due to bind error in bind manager.")
}

func (b BaseSocksV4Logger) SendClientError(client string, host string, address string, err error) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", host).
		Str("address", address).
		Err(err).
		Msg("Cannot handle v4 request due to bind error in bind manager.")
}
