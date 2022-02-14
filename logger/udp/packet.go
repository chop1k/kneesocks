package udp

import "github.com/rs/zerolog"

type PacketLogger struct {
	logger zerolog.Logger
}

func NewPacketLogger(logger zerolog.Logger) (PacketLogger, error) {
	return PacketLogger{
		logger: logger,
	}, nil
}

func (b PacketLogger) Accepted(client string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Msg("Packet accepted.")
}

func (b PacketLogger) Denied(client string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Msg("Packet denied.")
}

func (b PacketLogger) NotAllowed(client string) {
	e := b.logger.Warn()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Msg("Packet not allowed due to ruleset.")
}

func (b PacketLogger) NotAllowedByWhitelist(client string) {
	e := b.logger.Warn()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Msg("Packet not allowed by whitelist.")
}

func (b PacketLogger) NotAllowedByBlacklist(client string) {
	e := b.logger.Warn()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Msg("Packet not allowed by blacklist.")
}

func (b PacketLogger) HostUnreachable(client string) {
	e := b.logger.Warn()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Msg("Host unreachable.")
}

func (b PacketLogger) NetworkUnreachable(client string) {
	e := b.logger.Warn()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Msg("Network unreachable.")
}

func (b PacketLogger) InvalidFragment(address string) {
	e := b.logger.Warn()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", address).
		Msg("Invalid fragment, dropping the package.")
}
