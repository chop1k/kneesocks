package v4a

import "github.com/rs/zerolog"

type ConnectLogger interface {
	Request(client string, address string)
	Failed(client string, address string)
	Successful(client string, address string)
	HostUnreachable(client string, address string)
	NetworkUnreachable(client string, address string)
	Refused(client string, address string)
	Timeout(client string, address string)
}

type BaseConnectLogger struct {
	logger zerolog.Logger
}

func NewBaseConnectLogger(logger zerolog.Logger) (BaseConnectLogger, error) {
	return BaseConnectLogger{
		logger: logger,
	}, nil
}

func (b BaseConnectLogger) Request(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Received connect request. ")
}

func (b BaseConnectLogger) Failed(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Connect failed. ")
}

func (b BaseConnectLogger) Successful(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Connect successful. ")
}

func (b BaseConnectLogger) ConnectNotAllowed(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Connect not allowed due to ruleset. ")
}

func (b BaseConnectLogger) ConnectNotAllowedByWhitelist(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Connect not allowed due to whitelist.")
}

func (b BaseConnectLogger) ConnectNotAllowedByBlacklist(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Connect not allowed due to blacklist.")
}
func (b BaseConnectLogger) HostUnreachable(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Host unreachable. ")
}

func (b BaseConnectLogger) Timeout(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Connect failed due to timeout. ")
}

func (b BaseConnectLogger) NetworkUnreachable(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Network unreachable. ")
}

func (b BaseConnectLogger) Refused(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Connect refused by host. ")
}
