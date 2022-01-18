package logger

import "github.com/rs/zerolog"

type ServerLogger interface {
	WhitelistMatchError(address string, pattern string, err error)
	BlacklistMatchError(address string, pattern string, err error)
}

type BaseServerLogger struct {
	logger zerolog.Logger
}

func NewBaseServerLogger(logger zerolog.Logger) (BaseServerLogger, error) {
	return BaseServerLogger{
		logger: logger,
	}, nil
}

func (b BaseServerLogger) WhitelistMatchError(address string, pattern string, err error) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("address", address).
		Str("pattern", pattern).
		Err(err).
		Msg("Got error while matching in whitelist.")
}

func (b BaseServerLogger) BlacklistMatchError(address string, pattern string, err error) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("address", address).
		Str("pattern", pattern).
		Err(err).
		Msg("Got error while matching in blacklist.")
}
