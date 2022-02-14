package tcp

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
		Msg("Got accept error.")
}

func (b ErrorsLogger) AddressParseError(address string, err error) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("address", address).
		Err(err).
		Msg("Cannot parse address.")
}

func (b ErrorsLogger) LookupError(address string, err error) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("address", address).
		Err(err).
		Msg("Got lookup error.")
}

func (b ErrorsLogger) ReceiveClientError(address string, err error) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("address", address).
		Err(err).
		Msg("Cannot receive client because of error.")
}

func (b ErrorsLogger) SendHostError(address string, err error) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("address", address).
		Err(err).
		Msg("Cannot send host because of error.")
}

func (b ErrorsLogger) WriteRequestError(client string, host string, err error) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", host).
		Err(err).
		Msg("Cannot send first packet because of error.")
}

func (b ErrorsLogger) ReceiveWelcomeError(client string, err error) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Err(err).
		Msg("Cannot receive welcome chunk because of error.")
}
