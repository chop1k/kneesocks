package tcp

import "github.com/rs/zerolog"

type ConnectionLogger interface {
	ConnectionAccepted(client string)
	ConnectionDenied(client string)
	ConnectionProtocolDetermined(client string, protocol string)
	ConnectionBound(client string, host string)
	ConnectionExchangeTimeout(client string)
}

type BaseConnectionLogger struct {
	logger zerolog.Logger
}

func NewBaseConnectionLogger(logger zerolog.Logger) (BaseConnectionLogger, error) {
	return BaseConnectionLogger{
		logger: logger,
	}, nil
}

func (b BaseConnectionLogger) ConnectionAccepted(client string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Msg("Tcp connection accepted.")
}

func (b BaseConnectionLogger) ConnectionDenied(client string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Msg("Tcp connection denied.")
}

func (b BaseConnectionLogger) ConnectionProtocolDetermined(client string, protocol string) {
	e := b.logger.Debug()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("protocol", protocol).
		Msg("Connection protocol determined.")
}

func (b BaseConnectionLogger) ConnectionBound(client string, host string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", host).
		Msg("Connection bound.")
}

func (b BaseConnectionLogger) ConnectionExchangeTimeout(client string) {
	e := b.logger.Warn()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Msg("Connection exchange timeout.")
}
