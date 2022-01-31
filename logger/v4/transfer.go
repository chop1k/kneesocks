package v4

import "github.com/rs/zerolog"

type TransferLogger interface {
	TransferFinished(client string, host string)
}

type BaseTransferLogger struct {
	logger zerolog.Logger
}

func NewBaseTransferLogger(logger zerolog.Logger) (BaseTransferLogger, error) {
	return BaseTransferLogger{
		logger: logger,
	}, nil
}

func (b BaseTransferLogger) TransferFinished(client string, host string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", host).
		Msg("Transfer finished. ")
}
