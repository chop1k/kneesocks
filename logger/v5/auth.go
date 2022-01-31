package v5

import "github.com/rs/zerolog"

type AuthLogger interface {
	AuthenticationSuccessful(client string, name string)
	AuthenticationFailed(client string)
}

type BaseAuthLogger struct {
	logger zerolog.Logger
}

func NewBaseAuthLogger(logger zerolog.Logger) (BaseAuthLogger, error) {
	return BaseAuthLogger{
		logger: logger,
	}, nil
}

func (b BaseAuthLogger) AuthenticationSuccessful(client string, name string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("name", name).
		Msg("Authentication successful. ")
}

func (b BaseAuthLogger) AuthenticationFailed(client string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Msg("Authentication failed. ")
}
