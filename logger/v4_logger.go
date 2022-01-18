package logger

import (
	"github.com/rs/zerolog"
)

type SocksV4Logger interface {
	ConnectRequest(client string, address string)
	ConnectFailed(client string, address string)
	ConnectSuccessful(client string, address string)
	ConnectNotAllowed(client string, address string)
	ConnectHostUnreachable(client string, address string)
	ConnectNetworkUnreachable(client string, address string)
	ConnectRefused(client string, address string)
	ConnectTimeout(client string, address string)
	BindRequest(client string, address string)
	BindFailed(client string, address string)
	BindSuccessful(client string, address string)
	BindNotAllowed(client string, address string)
	BindTimeout(client string, address string)
	Bound(client string, host string)
	TransferFinished(client string, host string)
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
		Msg("Received connect request. ")
}

func (b BaseSocksV4Logger) ConnectFailed(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Connect failed. ")
}

func (b BaseSocksV4Logger) ConnectSuccessful(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Connect successful. ")
}

func (b BaseSocksV4Logger) ConnectNotAllowed(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Connect not allowed due to ruleset. ")
}

func (b BaseSocksV4Logger) ConnectHostUnreachable(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Host unreachable. ")
}

func (b BaseSocksV4Logger) ConnectNetworkUnreachable(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Network unreachable. ")
}

func (b BaseSocksV4Logger) ConnectRefused(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Connect refused by host. ")
}

func (b BaseSocksV4Logger) ConnectTimeout(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Connect failed due to timeout. ")
}

func (b BaseSocksV4Logger) BindRequest(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Received bind request. ")
}

func (b BaseSocksV4Logger) BindFailed(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Bind failed. ")
}

func (b BaseSocksV4Logger) BindSuccessful(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Bind successful. ")
}

func (b BaseSocksV4Logger) BindNotAllowed(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Bind failed due to ruleset. ")
}

func (b BaseSocksV4Logger) BindTimeout(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Bind failed due to timeout. ")
}

func (b BaseSocksV4Logger) Bound(client string, host string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", host).
		Msg("Bound successfully. ")
}

func (b BaseSocksV4Logger) TransferFinished(client string, host string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", host).
		Msg("Transfer finished. ")
}
