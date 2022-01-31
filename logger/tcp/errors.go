package tcp

import "github.com/rs/zerolog"

type ErrorsLogger interface {
	ListenError(address string, err error)
	AcceptError(err error)
	AddressParseError(address string, err error)
	LookupError(address string, err error)
	ReceiveClientError(address string, err error)
	SendHostError(address string, err error)
	WriteRequestError(client string, host string, err error)
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
		Msg("Got accept error.")
}

func (b BaseErrorsLogger) AddressParseError(address string, err error) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("address", address).
		Err(err).
		Msg("Cannot parse address.")
}

func (b BaseErrorsLogger) LookupError(address string, err error) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("address", address).
		Err(err).
		Msg("Got lookup error.")
}

func (b BaseErrorsLogger) ReceiveClientError(address string, err error) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("address", address).
		Err(err).
		Msg("Cannot receive client because of error.")
}

func (b BaseErrorsLogger) SendHostError(address string, err error) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("address", address).
		Err(err).
		Msg("Cannot send host because of error.")
}

func (b BaseErrorsLogger) WriteRequestError(client string, host string, err error) {
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
