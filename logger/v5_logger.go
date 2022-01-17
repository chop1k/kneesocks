package logger

import (
	"github.com/rs/zerolog"
)

type SocksV5Logger interface {
	ConnectRequest(client string, address string)
	ConnectFailed(client string, address string)
	ConnectSuccessful(client string, address string)
	ConnectNotAllowed(client string, address string)
	ConnectTimeout(client string, address string)
	ConnectRefused(client string, address string)
	ConnectHostUnreachable(client string, address string)
	ConnectNetworkUnreachable(client string, address string)
	BindRequest(client string, address string)
	BindFailed(client string, address string)
	BindSuccessful(client string, address string)
	BindNotAllowed(client string, address string)
	BindTimeout(client string, address string)
	Bound(client string, host string)
	UdpAssociationRequest(client string)
	UdpAssociationSuccessful(client string, address string)
	UdpAssociationFailed(client string, address string)
	UdpAssociationNotAllowed(client string)
	AuthenticationSuccessful(client string)
	AuthenticationFailed(client string)
	TransferFinished(client string, host string)
}

type BaseSocksV5Logger struct {
	logger zerolog.Logger
}

func NewBaseSocksV5Logger(logger zerolog.Logger) (BaseSocksV5Logger, error) {
	return BaseSocksV5Logger{
		logger: logger,
	}, nil
}

func (b BaseSocksV5Logger) ConnectRequest(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Received connect request.")
}

func (b BaseSocksV5Logger) ConnectFailed(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Connect failed.")
}

func (b BaseSocksV5Logger) ConnectSuccessful(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Connect successful.")
}

func (b BaseSocksV5Logger) ConnectNotAllowed(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Connect not allowed due to ruleset.")
}

func (b BaseSocksV5Logger) ConnectTimeout(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Connect failed due to timeout. ")
}

func (b BaseSocksV5Logger) ConnectRefused(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Connect refused by host. ")
}

func (b BaseSocksV5Logger) ConnectHostUnreachable(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Host unreachable. ")
}

func (b BaseSocksV5Logger) ConnectNetworkUnreachable(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Network unreachable. ")
}

func (b BaseSocksV5Logger) BindRequest(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Received bind request. ")
}

func (b BaseSocksV5Logger) BindFailed(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Bind failed. ")
}

func (b BaseSocksV5Logger) BindSuccessful(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Bind successful. ")
}

func (b BaseSocksV5Logger) BindNotAllowed(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Bind not allowed due to ruleset. ")
}

func (b BaseSocksV5Logger) BindTimeout(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Bind failed due to timeout. ")
}

func (b BaseSocksV5Logger) Bound(client string, host string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host_bound", host).
		Msg("Bound successfully. ")
}

func (b BaseSocksV5Logger) UdpAssociationRequest(client string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Msg("Received udp association request. ")
}

func (b BaseSocksV5Logger) UdpAssociationSuccessful(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("bind_address", address).
		Msg("Udp associate successful. ")
}

func (b BaseSocksV5Logger) UdpAssociationFailed(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("bind_address", address).
		Msg("Udp associate failed. ")
}

func (b BaseSocksV5Logger) UdpAssociationNotAllowed(client string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Msg("Udp associate not allowed due to ruleset. ")
}

func (b BaseSocksV5Logger) AuthenticationSuccessful(client string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Msg("Authentication successful. ")
}

func (b BaseSocksV5Logger) AuthenticationFailed(client string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Msg("Authentication failed. ")
}

func (b BaseSocksV5Logger) TransferFinished(client string, host string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", host).
		Msg("Transfer finished. ")
}
