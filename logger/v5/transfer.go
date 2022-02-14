package v5

import "github.com/rs/zerolog"

type TransferLogger struct {
	logger zerolog.Logger
}

func NewTransferLogger(logger zerolog.Logger) (TransferLogger, error) {
	return TransferLogger{
		logger: logger,
	}, nil
}

func (b TransferLogger) TransferFinished(client string, host string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", host).
		Msg("Transfer finished. ")
}
