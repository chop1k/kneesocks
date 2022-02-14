package v5

import "github.com/rs/zerolog"

type RestrictionsLogger struct {
	logger zerolog.Logger
}

func NewRestrictionsLogger(logger zerolog.Logger) (RestrictionsLogger, error) {
	return RestrictionsLogger{
		logger: logger,
	}, nil
}

func (b RestrictionsLogger) NotAllowed(client string, address string) {
	e := b.logger.Warn()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Command not allowed due to ruleset.")
}

func (b RestrictionsLogger) NotAllowedByWhitelist(client string, address string) {
	e := b.logger.Warn()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Not allowed due to whitelist.")
}

func (b RestrictionsLogger) NotAllowedByBlacklist(client string, address string) {
	e := b.logger.Warn()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Not allowed due to blacklist.")
}

func (b RestrictionsLogger) NotAllowedByConnectionLimits(client string, address string) {
	e := b.logger.Warn()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Not allowed due to maximum simultaneous connections.")
}

func (b RestrictionsLogger) IPv4AddressNotAllowed(client string, address string) {
	e := b.logger.Warn()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("address", address).
		Msg("IPv4 address type not allowed by ruleset. ")
}

func (b RestrictionsLogger) DomainAddressNotAllowed(client string, address string) {
	e := b.logger.Warn()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("address", address).
		Msg("Domain address type not allowed by ruleset. ")
}

func (b RestrictionsLogger) IPv6AddressNotAllowed(client string, address string) {
	e := b.logger.Warn()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("address", address).
		Msg("IPv6 address type not allowed by ruleset. ")
}
