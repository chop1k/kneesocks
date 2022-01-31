package v5

import "github.com/rs/zerolog"

type AuthLogger interface {
	Successful(client string, name string)
	Failed(client string)
}

type BaseAuthLogger struct {
	logger zerolog.Logger
}

func NewBaseAuthLogger(logger zerolog.Logger) (BaseAuthLogger, error) {
	return BaseAuthLogger{
		logger: logger,
	}, nil
}

func (b BaseAuthLogger) Successful(client string, name string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("name", name).
		Msg("Authentication successful. ")
}

func (b BaseAuthLogger) Failed(client string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Msg("Authentication failed. ")
}
