package logger

import (
	"fmt"
	"github.com/rs/zerolog"
	v4 "socks/protocol/v4"
)

type SocksV4Logger interface {
	ConnectRequest(client string, chunk v4.RequestChunk)
	ConnectFailed(client string, chunk v4.RequestChunk)
	ConnectSuccessful(client string, chunk v4.RequestChunk)
	ConnectNotAllowed(client string, chunk v4.RequestChunk)
	ConnectUnreachable(client string, chunk v4.RequestChunk)
	ConnectTimeout(client string, chunk v4.RequestChunk)
	BindRequest(client string, chunk v4.RequestChunk)
	BindFailed(client string, chunk v4.RequestChunk)
	BindSuccessful(client string, chunk v4.RequestChunk)
	BindNotAllowed(client string, chunk v4.RequestChunk)
	BindTimeout(client string, chunk v4.RequestChunk)
	Bound(client string, host string, chunk v4.RequestChunk)
	TransferFinished(client string, host string)
}

type BaseSocksV4Logger struct {
	logger zerolog.Logger
}

func NewBaseSocksV4Logger(logger zerolog.Logger) (BaseSocksV4Logger, error) {
	return BaseSocksV4Logger{
		logger: logger,
	}, nil
}

func (b BaseSocksV4Logger) ConnectRequest(client string, chunk v4.RequestChunk) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", fmt.Sprintf("%s:%d", chunk.DestinationIp, chunk.DestinationPort)).
		Str("user_id", chunk.UserId).
		Msg("Received connect request. ")
}

func (b BaseSocksV4Logger) ConnectFailed(client string, chunk v4.RequestChunk) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", fmt.Sprintf("%s:%d", chunk.DestinationIp, chunk.DestinationPort)).
		Str("user_id", chunk.UserId).
		Msg("Connect failed. ")
}

func (b BaseSocksV4Logger) ConnectSuccessful(client string, chunk v4.RequestChunk) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", fmt.Sprintf("%s:%d", chunk.DestinationIp, chunk.DestinationPort)).
		Str("user_id", chunk.UserId).
		Msg("Connect successful. ")
}

func (b BaseSocksV4Logger) ConnectNotAllowed(client string, chunk v4.RequestChunk) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", fmt.Sprintf("%s:%d", chunk.DestinationIp, chunk.DestinationPort)).
		Str("user_id", chunk.UserId).
		Msg("Connect not allowed due to ruleset. ")
}

func (b BaseSocksV4Logger) ConnectUnreachable(client string, chunk v4.RequestChunk) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", fmt.Sprintf("%s:%d", chunk.DestinationIp, chunk.DestinationPort)).
		Str("user_id", chunk.UserId).
		Msg("Host unreachable. ")
}

func (b BaseSocksV4Logger) ConnectTimeout(client string, chunk v4.RequestChunk) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", fmt.Sprintf("%s:%d", chunk.DestinationIp, chunk.DestinationPort)).
		Str("user_id", chunk.UserId).
		Msg("Connect failed due to timeout. ")
}

func (b BaseSocksV4Logger) BindRequest(client string, chunk v4.RequestChunk) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", fmt.Sprintf("%s:%d", chunk.DestinationIp, chunk.DestinationPort)).
		Str("user_id", chunk.UserId).
		Msg("Received bind request. ")
}

func (b BaseSocksV4Logger) BindFailed(client string, chunk v4.RequestChunk) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", fmt.Sprintf("%s:%d", chunk.DestinationIp, chunk.DestinationPort)).
		Str("user_id", chunk.UserId).
		Msg("Bind failed. ")
}

func (b BaseSocksV4Logger) BindSuccessful(client string, chunk v4.RequestChunk) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", fmt.Sprintf("%s:%d", chunk.DestinationIp, chunk.DestinationPort)).
		Str("user_id", chunk.UserId).
		Msg("Bind successful. ")
}

func (b BaseSocksV4Logger) BindNotAllowed(client string, chunk v4.RequestChunk) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", fmt.Sprintf("%s:%d", chunk.DestinationIp, chunk.DestinationPort)).
		Str("user_id", chunk.UserId).
		Msg("Bind failed due to ruleset. ")
}

func (b BaseSocksV4Logger) BindTimeout(client string, chunk v4.RequestChunk) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", fmt.Sprintf("%s:%d", chunk.DestinationIp, chunk.DestinationPort)).
		Str("user_id", chunk.UserId).
		Msg("Bind failed due to timeout. ")
}

func (b BaseSocksV4Logger) Bound(client string, host string, chunk v4.RequestChunk) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", host).
		Str("host", fmt.Sprintf("%s:%d", chunk.DestinationIp, chunk.DestinationPort)).
		Str("user_id", chunk.UserId).
		Msg("Bound successfully. ")
}

func (b BaseSocksV4Logger) TransferFinished(client string, host string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", host).
		Msg("Transfer finished. ")
}
