package v5

import "github.com/rs/zerolog"

type AuthLogger struct {
	logger zerolog.Logger
}

func NewAuthLogger(logger zerolog.Logger) (AuthLogger, error) {
	return AuthLogger{
		logger: logger,
	}, nil
}

func (b AuthLogger) Successful(client string, name string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("name", name).
		Msg("Authentication successful.")
}

func (b AuthLogger) Failed(client string) {
	e := b.logger.Warn()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Msg("Authentication failed.")
}
