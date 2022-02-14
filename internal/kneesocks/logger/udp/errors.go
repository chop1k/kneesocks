package udp

import "github.com/rs/zerolog"

type ErrorsLogger struct {
	logger zerolog.Logger
}

func NewErrorsLogger(logger zerolog.Logger) (ErrorsLogger, error) {
	return ErrorsLogger{
		logger: logger,
	}, nil
}

func (b ErrorsLogger) ListenError(address string, err error) {
	e := b.logger.Fatal()

	if !e.Enabled() {
		return
	}

	e.
		Str("address", address).
		Err(err).
		Msg("Got listen error.")
}

func (b ErrorsLogger) AcceptError(err error) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Err(err).
		Msg("Got listen error.")
}

func (b ErrorsLogger) WriteError(client string, err error) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Err(err).
		Msg("Got write error.")
}

func (b ErrorsLogger) ParseAddressError(client string, address string, err error) {
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

func (b ErrorsLogger) ManagerError(client string, err error) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Err(err).
		Msg("Got error from manager.")
}

func (b ErrorsLogger) ResolveError(client string, address string, err error) {
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

func (b ErrorsLogger) DialError(client string, address string, err error) {
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

func (b ErrorsLogger) WriteMsgError(client string, address string, err error) {
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

func (b ErrorsLogger) DeadlineError(client string, err error) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Err(err).
		Msg("Got deadline error.")
}

func (b ErrorsLogger) DeterminationError(client string, address string, err error) {
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

func (b ErrorsLogger) BuildError(client string, err error) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Err(err).
		Msg("Got build error.")
}
