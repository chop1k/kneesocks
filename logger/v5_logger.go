package logger

import (
	"fmt"
	"github.com/rs/zerolog"
	v5 "socks/protocol/v5"
)

type SocksV5Logger interface {
	ConnectRequest(client string, chunk v5.RequestChunk)
	ConnectFailed(client string, chunk v5.RequestChunk)
	ConnectSuccessful(client string, chunk v5.RequestChunk)
	ConnectNowAllowed(client string, chunk v5.RequestChunk)
	ConnectTimeout(client string, chunk v5.RequestChunk)
	ConnectUnreachable(client string, chunk v5.RequestChunk)
	BindRequest(client string, chunk v5.RequestChunk)
	BindFailed(client string, chunk v5.RequestChunk)
	BindSuccessful(client string, chunk v5.RequestChunk)
	BindNotAllowed(client string, chunk v5.RequestChunk)
	BindTimeout(client string, chunk v5.RequestChunk)
	Bound(client string, host string, chunk v5.RequestChunk)
	UdpAssociationRequest(client string, chunk v5.RequestChunk)
	UdpAssociationSuccessful(client string, chunk v5.RequestChunk)
	UdpAssociationFailed(client string, chunk v5.RequestChunk)
	UdpAssociationNotAllowed(client string, chunk v5.RequestChunk)
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

func (b BaseSocksV5Logger) ConnectRequest(client string, chunk v5.RequestChunk) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", fmt.Sprintf("%s:%d", chunk.Address, chunk.Port)).
		Msg("Received connect request.")
}

func (b BaseSocksV5Logger) ConnectFailed(client string, chunk v5.RequestChunk) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", fmt.Sprintf("%s:%d", chunk.Address, chunk.Port)).
		Msg("Connect failed.")
}

func (b BaseSocksV5Logger) ConnectSuccessful(client string, chunk v5.RequestChunk) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", fmt.Sprintf("%s:%d", chunk.Address, chunk.Port)).
		Msg("Connect successful.")
}

func (b BaseSocksV5Logger) ConnectNowAllowed(client string, chunk v5.RequestChunk) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", fmt.Sprintf("%s:%d", chunk.Address, chunk.Port)).
		Msg("Connect not allowed due to ruleset.")
}

func (b BaseSocksV5Logger) ConnectUnreachable(client string, chunk v5.RequestChunk) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", fmt.Sprintf("%s:%d", chunk.Address, chunk.Port)).
		Msg("Host unreachable.")
}

func (b BaseSocksV5Logger) ConnectTimeout(client string, chunk v5.RequestChunk) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", fmt.Sprintf("%s:%d", chunk.Address, chunk.Port)).
		Msg("Connect failed due to timeout. ")
}

func (b BaseSocksV5Logger) BindRequest(client string, chunk v5.RequestChunk) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", fmt.Sprintf("%s:%d", chunk.Address, chunk.Port)).
		Msg("Received bind request. ")
}

func (b BaseSocksV5Logger) BindFailed(client string, chunk v5.RequestChunk) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", fmt.Sprintf("%s:%d", chunk.Address, chunk.Port)).
		Msg("Bind failed. ")
}

func (b BaseSocksV5Logger) BindSuccessful(client string, chunk v5.RequestChunk) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", fmt.Sprintf("%s:%d", chunk.Address, chunk.Port)).
		Msg("Bind successful. ")
}

func (b BaseSocksV5Logger) BindNotAllowed(client string, chunk v5.RequestChunk) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", fmt.Sprintf("%s:%d", chunk.Address, chunk.Port)).
		Msg("Bind not allowed due to ruleset. ")
}

func (b BaseSocksV5Logger) BindTimeout(client string, chunk v5.RequestChunk) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", fmt.Sprintf("%s:%d", chunk.Address, chunk.Port)).
		Msg("Bind failed due to timeout. ")
}

func (b BaseSocksV5Logger) Bound(client string, host string, chunk v5.RequestChunk) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", fmt.Sprintf("%s:%d", chunk.Address, chunk.Port)).
		Str("host_bound", host).
		Msg("Bound successfully. ")
}

func (b BaseSocksV5Logger) UdpAssociationRequest(client string, chunk v5.RequestChunk) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", fmt.Sprintf("%s:%d", chunk.Address, chunk.Port)).
		Msg("Received udp association request. ")
}

func (b BaseSocksV5Logger) UdpAssociationSuccessful(client string, chunk v5.RequestChunk) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", fmt.Sprintf("%s:%d", chunk.Address, chunk.Port)).
		Msg("Udp associate successful. ")
}

func (b BaseSocksV5Logger) UdpAssociationFailed(client string, chunk v5.RequestChunk) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", fmt.Sprintf("%s:%d", chunk.Address, chunk.Port)).
		Msg("Udp associate failed. ")
}

func (b BaseSocksV5Logger) UdpAssociationNotAllowed(client string, chunk v5.RequestChunk) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", fmt.Sprintf("%s:%d", chunk.Address, chunk.Port)).
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
