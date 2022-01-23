package main

import "github.com/rs/zerolog"

type Logger struct {
	logger zerolog.Logger
}

func NewLogger(logger zerolog.Logger) (Logger, error) {
	return Logger{
		logger: logger,
	}, nil
}

func (l Logger) ListenTcp(address string) {
	e := l.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("address", address).
		Msg("Tcp server listening. ")
}

func (l Logger) ListenUdp(address string) {
	e := l.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("address", address).
		Msg("Udp server listening.")
}

func (l Logger) Connection(address string, bindAddress string) {
	e := l.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("address", address).
		Str("bind_address", bindAddress).
		Msg("New connection.")
}

func (l Logger) AcceptError(address string, err error) {
	e := l.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("address", address).
		Err(err).
		Msg("Cannot accept connection because of error.")
}

func (l Logger) AcceptPacketError(address string, err error) {
	e := l.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("address", address).
		Err(err).
		Msg("Cannot accept packet because of error.")
}

func (l Logger) PictureRequest(address string, picture byte) {
	e := l.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("address", address).
		Uint8("picture", picture).
		Msg("Received picture request.")
}

func (l Logger) DialError(address string, err error) {
	e := l.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("address", address).
		Err(err).
		Msg("Cannot dial because of error.")
}

func (l Logger) FileError(err error) {
	e := l.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Err(err).
		Msg("Cannot open file because of error.")
}

func (l Logger) IOError(address string, err error) {
	e := l.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("address", address).
		Err(err).
		Msg("Data transfer failed because of error.")
}

func (l Logger) InvalidPicture(address string, picture byte) {
	e := l.logger.Warn()

	if !e.Enabled() {
		return
	}

	e.
		Str("address", address).
		Uint8("picture", picture).
		Msg("Data transfer failed because of error.")
}

func (l Logger) PacketAccepted(address string, bindAddress string) {
	e := l.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("address", address).
		Str("bind_address", bindAddress).
		Msg("New packet accepted.")
}
