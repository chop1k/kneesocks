package v4

import "github.com/rs/zerolog"

type ErrorsLogger interface {
	ParseError(client string, err error)
	UnknownError(client string, address string, err error)
	AddressParsingError(client string, host string, address string, err error)
	AddressDeterminationError(client string, host string, address string, err error)
	InvalidAddressTypeError(client string, host string, address string)
	BindError(client string, address string, err error)
	ReceiveHostError(client string, address string, err error)
	SendClientError(client string, host string, address string, err error)
}

type BaseErrorsLogger struct {
	logger zerolog.Logger
}

func NewBaseErrorsLogger(logger zerolog.Logger) (BaseErrorsLogger, error) {
	return BaseErrorsLogger{
		logger: logger,
	}, nil
}

func (b BaseErrorsLogger) ParseError(client string, err error) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Err(err).
		Msg("Cannot parse v4 request due to error.")
}

func (b BaseErrorsLogger) UnknownError(client string, address string, err error) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Err(err).
		Msg("Cannot handle v4 request due to error.")
}

func (b BaseErrorsLogger) AddressParsingError(client string, host string, address string, err error) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", host).
		Str("address", address).
		Err(err).
		Msg("Cannot handle v4 request due to address parsing error.")
}

func (b BaseErrorsLogger) AddressDeterminationError(client string, host string, address string, err error) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", host).
		Str("address", address).
		Err(err).
		Msg("Cannot handle v4 request due to address determination error.")
}

func (b BaseErrorsLogger) InvalidAddressTypeError(client string, host string, address string) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", host).
		Str("address", address).
		Msg("Cannot handle v4 request due to invalid address type.")
}

func (b BaseErrorsLogger) BindError(client string, address string, err error) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Err(err).
		Msg("Cannot handle v4 request due to bind error in bind manager.")
}

func (b BaseErrorsLogger) ReceiveHostError(client string, address string, err error) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Err(err).
		Msg("Cannot handle v4 request due to bind error in bind manager.")
}

func (b BaseErrorsLogger) SendClientError(client string, host string, address string, err error) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", host).
		Str("address", address).
		Err(err).
		Msg("Cannot handle v4 request due to bind error in bind manager.")
}
