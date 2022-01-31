package v5

import "github.com/rs/zerolog"

type ErrorsLogger interface {
	ParseError(client string, err error)
	UnknownError(client string, address string, err error)
	AddressParsingError(client string, host string, address string, err error)
	UdpAddressParsingError(client string, err error)
	AddressDeterminationError(client string, host string, address string, err error)
	InvalidAddressTypeError(client string, addressType byte, address string)
	UnknownCommandError(client string, command byte, address string)
	SelectMethodsError(client string, err error)
	ReceiveRequestError(client string, err error)
	PasswordResponseError(client string, user string, err error)
	ParseMethodsError(client string, err error)
	ReceiveHostError(client string, address string, err error)
	SendClientError(client string, host string, address string, err error)
	BindError(client string, address string, err error)
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
		Msg("Cannot parse v5 request due to error. ")
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
		Msg("Cannot handle v5 request due to error. ")
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
		Msg("Cannot handle v5 request due to address parsing error. ")
}

func (b BaseErrorsLogger) UdpAddressParsingError(client string, err error) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Err(err).
		Msg("Cannot handle v5 request due to address parsing error. ")
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
		Msg("Cannot handle v5 request due to address determination error. ")
}

func (b BaseErrorsLogger) InvalidAddressTypeError(client string, addressType byte, address string) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Uint8("addressType", addressType).
		Str("address", address).
		Msg("Cannot handle v5 request due to invalid address type. ")
}

func (b BaseErrorsLogger) UnknownCommandError(client string, command byte, address string) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Uint8("command", command).
		Str("address", address).
		Msg("Cannot handle v5 request due to invalid command. ")
}

func (b BaseErrorsLogger) SelectMethodsError(client string, err error) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Err(err).
		Msg("Cannot handle v5 request due to methods selection error. ")
}

func (b BaseErrorsLogger) ReceiveRequestError(client string, err error) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Err(err).
		Msg("Cannot handle v5 request due to receive request error. ")
}

func (b BaseErrorsLogger) PasswordResponseError(client string, user string, err error) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("user", user).
		Err(err).
		Msg("Cannot handle v5 request due to password response error. ")
}

func (b BaseErrorsLogger) ParseMethodsError(client string, err error) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Err(err).
		Msg("Cannot handle v5 request due to parse methods error. ")
}

func (b BaseErrorsLogger) ReceiveHostError(client string, address string, err error) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("address", address).
		Err(err).
		Msg("Cannot handle v5 request due to receive host error in bind manager. ")
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
		Msg("Cannot handle v5 request due to send client error in bind manager. ")
}

func (b BaseErrorsLogger) BindError(client string, address string, err error) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("address", address).
		Err(err).
		Msg("Cannot handle v5 request due to bind error. ")
}
