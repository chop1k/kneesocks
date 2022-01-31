package udp

import "github.com/rs/zerolog"

type PacketLogger interface {
	PacketAccepted(address string)
	PacketDenied(address string)
	PacketNotAllowed(address string)
	PacketNotAllowedByWhitelist(address string)
	PacketNotAllowedByBlacklist(address string)
	PacketHostUnreachable(address string)
	PacketNetworkUnreachable(address string)
}

type BasePacketLogger struct {
	logger zerolog.Logger
}

func NewBasePacketLogger(logger zerolog.Logger) (BasePacketLogger, error) {
	return BasePacketLogger{
		logger: logger,
	}, nil
}

func (b BasePacketLogger) PacketAccepted(client string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Msg("Packet accepted.")
}

func (b BasePacketLogger) PacketDenied(client string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Msg("Packet denied.")
}

func (b BasePacketLogger) PacketNotAllowed(client string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Msg("Packet not allowed due to ruleset.")
}

func (b BasePacketLogger) PacketNotAllowedByWhitelist(client string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Msg("Packet not allowed by whitelist.")
}

func (b BasePacketLogger) PacketNotAllowedByBlacklist(client string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Msg("Packet not allowed by blacklist.")
}

func (b BasePacketLogger) PacketHostUnreachable(client string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Msg("Host unreachable.")
}

func (b BasePacketLogger) PacketNetworkUnreachable(client string) {
	e := b.logger.Warn()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Msg("Network unreachable.")
}
