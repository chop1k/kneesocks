package udp

import "github.com/rs/zerolog"

type PacketLogger interface {
	Accepted(address string)
	Denied(address string)
	NotAllowed(address string)
	NotAllowedByWhitelist(address string)
	NotAllowedByBlacklist(address string)
	HostUnreachable(address string)
	NetworkUnreachable(address string)
}

type BasePacketLogger struct {
	logger zerolog.Logger
}

func NewBasePacketLogger(logger zerolog.Logger) (BasePacketLogger, error) {
	return BasePacketLogger{
		logger: logger,
	}, nil
}

func (b BasePacketLogger) Accepted(client string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Msg("Packet accepted.")
}

func (b BasePacketLogger) Denied(client string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Msg("Packet denied.")
}

func (b BasePacketLogger) NotAllowed(client string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Msg("Packet not allowed due to ruleset.")
}

func (b BasePacketLogger) NotAllowedByWhitelist(client string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Msg("Packet not allowed by whitelist.")
}

func (b BasePacketLogger) NotAllowedByBlacklist(client string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Msg("Packet not allowed by blacklist.")
}

func (b BasePacketLogger) HostUnreachable(client string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Msg("Host unreachable.")
}

func (b BasePacketLogger) NetworkUnreachable(client string) {
	e := b.logger.Warn()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Msg("Network unreachable.")
}
