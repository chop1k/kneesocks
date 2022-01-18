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
	Listen(address string)
	ListenError(address string, err error)
	AcceptError(err error)
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
		Msg("Tcp connection accepted.")
}

func (b BaseTcpLogger) ConnectionDenied(client string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Msg("Tcp connection denied.")
}

func (b BaseTcpLogger) ConnectionProtocolDetermined(client string, protocol string) {
	e := b.logger.Debug()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("protocol", protocol).
		Msg("Connection protocol determined.")
}

func (b BaseTcpLogger) ConnectionBound(client string, host string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", host).
		Msg("Connection bound.")
}

func (b BaseTcpLogger) ConnectionExchangeTimeout(client string) {
	e := b.logger.Warn()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Msg("Connection exchange timeout.")
}

func (b BaseTcpLogger) Listen(address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("address", address).
		Msg("Listening for tcp connection.")
}

func (b BaseTcpLogger) ListenError(address string, err error) {
	e := b.logger.Fatal()

	if !e.Enabled() {
		return
	}

	e.
		Str("address", address).
		Err(err).
		Msg("Got listen error.")
}

func (b BaseTcpLogger) AcceptError(err error) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Err(err).
		Msg("Got accept error.")
}
