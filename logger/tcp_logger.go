package logger

import (
	"github.com/rs/zerolog"
)

type TcpLogger interface {
	ConnectionAccepted(client string)
	ConnectionDenied(client string)
	ConnectionProtocolDetermined(client string, protocol string)
	ConnectionBound(client string, host string)
	ConnectionExchangeTimeout(client string)
	Listen(addr string)
}

type BaseTcpLogger struct {
	logger zerolog.Logger
}

func NewBaseTcpLogger(logger zerolog.Logger) (BaseTcpLogger, error) {
	return BaseTcpLogger{
		logger: logger,
	}, nil
}

func (b BaseTcpLogger) ConnectionAccepted(client string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Msg("New tcp connection. ")
}

func (b BaseTcpLogger) ConnectionDenied(client string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Msg("New tcp connection. ")
}

func (b BaseTcpLogger) ConnectionProtocolDetermined(client string, protocol string) {
	e := b.logger.Debug()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("protocol", protocol).
		Msg("Connection protocol determined. ")
}

func (b BaseTcpLogger) ConnectionBound(client string, host string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", host).
		Msg("Connection bound. ")
}

func (b BaseTcpLogger) ConnectionExchangeTimeout(client string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Msg("Connection exchange timeout. ")
}

func (b BaseTcpLogger) Listen(addr string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("addr", addr).
		Msg("Listening for tcp connection. ")
}
