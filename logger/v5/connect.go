package v5

import "github.com/rs/zerolog"

type ConnectLogger interface {
	ConnectRequest(client string, address string)
	ConnectFailed(client string, address string)
	ConnectSuccessful(client string, address string)
	ConnectTimeout(client string, address string)
	ConnectRefused(client string, address string)
	ConnectHostUnreachable(client string, address string)
	ConnectNetworkUnreachable(client string, address string)
}

type BaseConnectLogger struct {
	logger zerolog.Logger
}

func NewBaseConnectLogger(logger zerolog.Logger) (BaseConnectLogger, error) {
	return BaseConnectLogger{
		logger: logger,
	}, nil
}

func (b BaseConnectLogger) ConnectRequest(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Received connect request.")
}

func (b BaseConnectLogger) ConnectFailed(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Connect failed.")
}

func (b BaseConnectLogger) ConnectSuccessful(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Connect successful.")
}

func (b BaseConnectLogger) ConnectTimeout(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Connect failed due to timeout. ")
}

func (b BaseConnectLogger) ConnectRefused(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Connect refused by host. ")
}

func (b BaseConnectLogger) ConnectHostUnreachable(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Host unreachable. ")
}

func (b BaseConnectLogger) ConnectNetworkUnreachable(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Network unreachable. ")
}
