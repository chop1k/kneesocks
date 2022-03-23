package v5

import "github.com/rs/zerolog"

type AssociationLogger struct {
	logger zerolog.Logger
}

func NewAssociationLogger(logger zerolog.Logger) (AssociationLogger, error) {
	return AssociationLogger{
		logger: logger,
	}, nil
}

func (b AssociationLogger) Request(client string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Msg("Received udp association request.")
}

func (b AssociationLogger) Successful(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("bind_address", address).
		Msg("Udp associate successful.")
}

func (b AssociationLogger) Failed(client string, address string) {
	e := b.logger.Warn()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("bind_address", address).
		Msg("Udp associate failed.")
}
