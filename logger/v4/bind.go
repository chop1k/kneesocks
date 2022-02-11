package v4

import "github.com/rs/zerolog"

type BindLogger interface {
	Request(client string, address string)
	Failed(client string, address string)
	Successful(client string, address string)
	Timeout(client string, address string)
	Bound(client string, host string)
}

type BaseBindLogger struct {
	logger zerolog.Logger
}

func NewBaseBindLogger(logger zerolog.Logger) (BaseBindLogger, error) {
	return BaseBindLogger{
		logger: logger,
	}, nil
}

func (b BaseBindLogger) Request(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Received bind request.")
}

func (b BaseBindLogger) Failed(client string, address string) {
	e := b.logger.Warn()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Bind failed.")
}

func (b BaseBindLogger) Successful(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Bind successful.")
}

func (b BaseBindLogger) Timeout(client string, address string) {
	e := b.logger.Warn()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Bind failed due to timeout.")
}

func (b BaseBindLogger) Bound(client string, host string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", host).
		Msg("Bound successfully.")
}
