package v4

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
