package udp

import "github.com/rs/zerolog"

type ErrorsLogger interface {
	ListenError(address string, err error)
	AcceptError(err error)
	WriteError(client string, err error)
	ParseAddressError(client string, address string, err error)
	ManagerError(client string, err error)
	ResolveError(client string, address string, err error)
	DialError(client string, address string, err error)
	WriteMsgError(client string, address string, err error)
	DeadlineError(client string, err error)
	DeterminationError(client string, address string, err error)
	BuildError(client string, err error)
}

type BaseErrorsLogger struct {
	logger zerolog.Logger
}

func NewBaseErrorsLogger(logger zerolog.Logger) (BaseErrorsLogger, error) {
	return BaseErrorsLogger{
		logger: logger,
	}, nil
}

func (b BaseErrorsLogger) ListenError(address string, err error) {
	e := b.logger.Fatal()

	if !e.Enabled() {
		return
	}

	e.
		Str("address", address).
		Err(err).
		Msg("Got listen error.")
}

func (b BaseErrorsLogger) AcceptError(err error) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Err(err).
		Msg("Got listen error.")
}

func (b BaseErrorsLogger) WriteError(client string, err error) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Err(err).
		Msg("Got write error.")
}

func (b BaseErrorsLogger) ParseAddressError(client string, address string, err error) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("address", address).
		Err(err).
		Msg("Got address parsing error.")
}

func (b BaseErrorsLogger) ManagerError(client string, err error) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Err(err).
		Msg("Got error from manager.")
}

func (b BaseErrorsLogger) ResolveError(client string, address string, err error) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("address", address).
		Err(err).
		Msg("Got resolve error.")
}

func (b BaseErrorsLogger) DialError(client string, address string, err error) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("address", address).
		Err(err).
		Msg("Got dial error.")
}

func (b BaseErrorsLogger) WriteMsgError(client string, address string, err error) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("address", address).
		Err(err).
		Msg("Got writeMsg error.")
}

func (b BaseErrorsLogger) DeadlineError(client string, err error) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Err(err).
		Msg("Got deadline error.")
}

func (b BaseErrorsLogger) DeterminationError(client string, address string, err error) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("address", address).
		Err(err).
		Msg("Got address determination error.")
}

func (b BaseErrorsLogger) BuildError(client string, err error) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Err(err).
		Msg("Got build error.")
}
