package logger

import "github.com/rs/zerolog"

type UdpLogger interface {
	PacketAccepted(address string)
	PacketDenied(address string)
	PacketNotAllowed(address string)
	PacketNotAllowedByWhitelist(address string)
	PacketNotAllowedByBlacklist(address string)
	PacketHostUnreachable(address string)
	PacketNetworkUnreachable(address string)
	Listen(address string)
	ListenError(address string, err error)
	AcceptError(err error)
}

type BaseUdpLogger struct {
	logger zerolog.Logger
}

func NewBaseUdpLogger(logger zerolog.Logger) (BaseUdpLogger, error) {
	return BaseUdpLogger{
		logger: logger,
	}, nil
}

func (b BaseUdpLogger) PacketAccepted(client string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Msg("Packet accepted.")
}

func (b BaseUdpLogger) PacketDenied(client string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Msg("Packet denied.")
}

func (b BaseUdpLogger) PacketNotAllowed(client string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Msg("Packet not allowed due to ruleset.")
}

func (b BaseUdpLogger) PacketNotAllowedByWhitelist(client string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Msg("Packet not allowed by whitelist.")
}

func (b BaseUdpLogger) PacketNotAllowedByBlacklist(client string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Msg("Packet not allowed by blacklist.")
}

func (b BaseUdpLogger) PacketHostUnreachable(client string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Msg("Host unreachable.")
}

func (b BaseUdpLogger) PacketNetworkUnreachable(client string) {
	e := b.logger.Warn()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Msg("Network unreachable.")
}

func (b BaseUdpLogger) Listen(address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("address", address).
		Msg("Listening for udp packets.")
}

func (b BaseUdpLogger) ListenError(address string, err error) {
	e := b.logger.Fatal()

	if !e.Enabled() {
		return
	}

	e.
		Str("address", address).
		Err(err).
		Msg("Got listen error.")
}

func (b BaseUdpLogger) AcceptError(err error) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Err(err).
		Msg("Got listen error.")
}
