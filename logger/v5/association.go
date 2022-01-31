package v5

import "github.com/rs/zerolog"

type AssociationLogger interface {
	UdpAssociationRequest(client string)
	UdpAssociationSuccessful(client string, address string)
	UdpAssociationFailed(client string, address string)
	UdpAssociationNotAllowed(client string)
}

type BaseAssociationLogger struct {
	logger zerolog.Logger
}

func NewBaseAssociationLogger(logger zerolog.Logger) (BaseAssociationLogger, error) {
	return BaseAssociationLogger{
		logger: logger,
	}, nil
}

func (b BaseAssociationLogger) UdpAssociationRequest(client string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Msg("Received udp association request. ")
}

func (b BaseAssociationLogger) UdpAssociationSuccessful(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("bind_address", address).
		Msg("Udp associate successful. ")
}

func (b BaseAssociationLogger) UdpAssociationFailed(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("bind_address", address).
		Msg("Udp associate failed. ")
}

func (b BaseAssociationLogger) UdpAssociationNotAllowed(client string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Msg("Udp associate not allowed due to ruleset. ")
}
