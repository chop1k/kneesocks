package tcp

import "github.com/rs/zerolog"

type ListenLogger interface {
	Listen(address string)
}

type BaseListenLogger struct {
	logger zerolog.Logger
}

func NewBaseListenLogger(logger zerolog.Logger) (BaseListenLogger, error) {
	return BaseListenLogger{
		logger: logger,
	}, nil
}

func (b BaseListenLogger) Listen(address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("address", address).
		Msg("Listening for tcp connection.")
}
