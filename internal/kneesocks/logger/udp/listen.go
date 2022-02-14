package udp

import "github.com/rs/zerolog"

type ListenLogger struct {
	logger zerolog.Logger
}

func NewListenLogger(logger zerolog.Logger) (ListenLogger, error) {
	return ListenLogger{
		logger: logger,
	}, nil
}

func (b ListenLogger) Listen(address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("address", address).
		Msg("Listening for udp packets.")
}
