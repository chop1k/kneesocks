package v5

import "github.com/rs/zerolog"

type AssociationLogger interface {
	Request(client string)
	Successful(client string, address string)
	Failed(client string, address string)
}

type BaseAssociationLogger struct {
	logger zerolog.Logger
}

func NewBaseAssociationLogger(logger zerolog.Logger) (BaseAssociationLogger, error) {
	return BaseAssociationLogger{
		logger: logger,
	}, nil
}

func (b BaseAssociationLogger) Request(client string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Msg("Received udp association request. ")
}

func (b BaseAssociationLogger) Successful(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("bind_address", address).
		Msg("Udp associate successful. ")
}

func (b BaseAssociationLogger) Failed(client string, address string) {
	e := b.logger.Warn()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("bind_address", address).
		Msg("Udp associate failed. ")
}
