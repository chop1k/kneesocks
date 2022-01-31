package tcp

import "github.com/rs/zerolog"

type ConnectionLogger interface {
	Accepted(client string)
	Denied(client string)
	ProtocolDetermined(client string, protocol string)
	Bound(client string, host string)
	ExchangeTimeout(client string)
}

type BaseConnectionLogger struct {
	logger zerolog.Logger
}

func NewBaseConnectionLogger(logger zerolog.Logger) (BaseConnectionLogger, error) {
	return BaseConnectionLogger{
		logger: logger,
	}, nil
}

func (b BaseConnectionLogger) Accepted(client string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Msg("Tcp connection accepted.")
}

func (b BaseConnectionLogger) Denied(client string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Msg("Tcp connection denied.")
}

func (b BaseConnectionLogger) ProtocolDetermined(client string, protocol string) {
	e := b.logger.Debug()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("protocol", protocol).
		Msg("Connection protocol determined.")
}

func (b BaseConnectionLogger) Bound(client string, host string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", host).
		Msg("Connection bound.")
}

func (b BaseConnectionLogger) ExchangeTimeout(client string) {
	e := b.logger.Warn()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Msg("Connection exchange timeout.")
}
