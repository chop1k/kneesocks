package tcp

import "github.com/rs/zerolog"

type ConnectionLogger struct {
	logger zerolog.Logger
}

func NewConnectionLogger(logger zerolog.Logger) (ConnectionLogger, error) {
	return ConnectionLogger{
		logger: logger,
	}, nil
}

func (b ConnectionLogger) Accepted(client string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Msg("Tcp connection accepted.")
}

func (b ConnectionLogger) Denied(client string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Msg("Tcp connection denied.")
}

func (b ConnectionLogger) ProtocolDetermined(client string, protocol string) {
	e := b.logger.Debug()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("protocol", protocol).
		Msg("Connection protocol determined.")
}

func (b ConnectionLogger) Bound(client string, host string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", host).
		Msg("Connection bound.")
}

func (b ConnectionLogger) ExchangeTimeout(client string) {
	e := b.logger.Warn()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Msg("Connection exchange timeout.")
}
