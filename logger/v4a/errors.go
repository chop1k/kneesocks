package v4a

import "github.com/rs/zerolog"

type ErrorsLogger struct {
	logger zerolog.Logger
}

func NewErrorsLogger(logger zerolog.Logger) (ErrorsLogger, error) {
	return ErrorsLogger{
		logger: logger,
	}, nil
}

func (b ErrorsLogger) ParseError(client string, err error) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Err(err).
		Msg("Cannot parse v4a request due to error. ")
}

func (b ErrorsLogger) UnknownError(client string, address string, err error) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Err(err).
		Msg("Cannot handle v4a request due to error. ")
}

func (b ErrorsLogger) AddressParsingError(client string, host string, address string, err error) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", host).
		Str("address", address).
		Err(err).
		Msg("Cannot handle v4a request due to address parsing error.")
}

func (b ErrorsLogger) AddressDeterminationError(client string, host string, address string, err error) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", host).
		Str("address", address).
		Err(err).
		Msg("Cannot handle v4a request due to address determination error.")
}

func (b ErrorsLogger) InvalidAddressTypeError(client string, host string, address string) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", host).
		Str("address", address).
		Msg("Cannot handle v4a request due to invalid address type.")
}

func (b ErrorsLogger) BindError(client string, address string, err error) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Err(err).
		Msg("Cannot handle v4a request due to bind error in bind manager.")
}

func (b ErrorsLogger) ReceiveHostError(client string, address string, err error) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Err(err).
		Msg("Cannot handle v4a request due to bind error in bind manager.")
}

func (b ErrorsLogger) SendClientError(client string, host string, address string, err error) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", host).
		Str("address", address).
		Err(err).
		Msg("Cannot handle v4a request due to bind error in bind manager.")
}

func (b ErrorsLogger) ConfigError(client string, err error) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Err(err).
		Msg("Got config error.")
}
