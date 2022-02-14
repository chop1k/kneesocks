package v5

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
		Msg("Cannot parse v5 request due to error. ")
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
		Msg("Cannot handle v5 request due to error. ")
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
		Msg("Cannot handle v5 request due to address parsing error. ")
}

func (b ErrorsLogger) UdpAddressParsingError(client string, err error) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Err(err).
		Msg("Cannot handle v5 request due to address parsing error. ")
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
		Msg("Cannot handle v5 request due to address determination error. ")
}

func (b ErrorsLogger) InvalidAddressTypeError(client string, addressType byte, address string) {
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

func (b ErrorsLogger) UnknownCommandError(client string, command byte, address string) {
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

func (b ErrorsLogger) SelectMethodsError(client string, err error) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Err(err).
		Msg("Cannot handle v5 request due to methods selection error. ")
}

func (b ErrorsLogger) ReceiveRequestError(client string, err error) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Err(err).
		Msg("Cannot handle v5 request due to receive request error. ")
}

func (b ErrorsLogger) PasswordResponseError(client string, user string, err error) {
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

func (b ErrorsLogger) ParseMethodsError(client string, err error) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Err(err).
		Msg("Cannot handle v5 request due to parse methods error. ")
}

func (b ErrorsLogger) ReceiveHostError(client string, address string, err error) {
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
		Msg("Cannot handle v5 request due to send client error in bind manager. ")
}

func (b ErrorsLogger) BindError(client string, address string, err error) {
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
