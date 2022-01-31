package udp

import "github.com/rs/zerolog"

type ErrorsLogger interface {
	ListenError(address string, err error)
	AcceptError(err error)
}

type BaseErrorsLogger struct {
	logger zerolog.Logger
}

func NewBaseErrorsLogger(logger zerolog.Logger) (BaseErrorsLogger, error) {
	return BaseErrorsLogger{
		logger: logger,
	}, nil
}

func (b BaseErrorsLogger) ListenError(address string, err error) {
	e := b.logger.Fatal()

	if !e.Enabled() {
		return
	}

	e.
		Str("address", address).
		Err(err).
		Msg("Got listen error.")
}

func (b BaseErrorsLogger) AcceptError(err error) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Err(err).
		Msg("Got listen error.")
}
