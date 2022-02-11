package v4a

import "github.com/rs/zerolog"

type RestrictionsLogger interface {
	NotAllowed(client string, address string)
	NotAllowedByWhitelist(client string, address string)
	NotAllowedByBlacklist(client string, address string)
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
