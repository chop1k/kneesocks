package v5

import "github.com/rs/zerolog"

type RestrictionsLogger interface {
	NotAllowed(client string, address string)
	NotAllowedByWhitelist(client string, address string)
	NotAllowedByBlacklist(client string, address string)
	NotAllowedByConnectionLimits(client string, address string)
	IPv4AddressNotAllowed(client string, address string)
	DomainAddressNotAllowed(client string, address string)
	IPv6AddressNotAllowed(client string, address string)
}

type BaseRestrictionsLogger struct {
	logger zerolog.Logger
}

func NewBaseRestrictionsLogger(logger zerolog.Logger) (BaseRestrictionsLogger, error) {
	return BaseRestrictionsLogger{
		logger: logger,
	}, nil
}

func (b BaseRestrictionsLogger) NotAllowed(client string, address string) {
	e := b.logger.Warn()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Command not allowed due to ruleset.")
}

func (b BaseRestrictionsLogger) NotAllowedByWhitelist(client string, address string) {
	e := b.logger.Warn()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Not allowed due to whitelist.")
}

func (b BaseRestrictionsLogger) NotAllowedByBlacklist(client string, address string) {
	e := b.logger.Warn()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Not allowed due to blacklist.")
}

func (b BaseRestrictionsLogger) NotAllowedByConnectionLimits(client string, address string) {
	e := b.logger.Warn()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Not allowed due to maximum simultaneous connections.")
}

func (b BaseRestrictionsLogger) IPv4AddressNotAllowed(client string, address string) {
	e := b.logger.Warn()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("address", address).
		Msg("IPv4 address type not allowed by ruleset. ")
}

func (b BaseRestrictionsLogger) DomainAddressNotAllowed(client string, address string) {
	e := b.logger.Warn()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("address", address).
		Msg("Domain address type not allowed by ruleset. ")
}

func (b BaseRestrictionsLogger) IPv6AddressNotAllowed(client string, address string) {
	e := b.logger.Warn()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("address", address).
		Msg("IPv6 address type not allowed by ruleset. ")
}
