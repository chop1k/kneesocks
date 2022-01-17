package logger

import (
	"fmt"
	"github.com/rs/zerolog"
	"socks/protocol/v4a"
)

type SocksV4aLogger interface {
	ConnectRequest(client string, chunk v4a.RequestChunk)
	ConnectFailed(client string, chunk v4a.RequestChunk)
	ConnectSuccessful(client string, chunk v4a.RequestChunk)
	ConnectNotAllowed(client string, chunk v4a.RequestChunk)
	ConnectUnreachable(client string, chunk v4a.RequestChunk)
	ConnectTimeout(client string, chunk v4a.RequestChunk)
	BindRequest(client string, chunk v4a.RequestChunk)
	BindFailed(client string, chunk v4a.RequestChunk)
	BindSuccessful(client string, chunk v4a.RequestChunk)
	BindNotAllowed(client string, chunk v4a.RequestChunk)
	BindTimeout(client string, chunk v4a.RequestChunk)
	Bound(client string, host string, chunk v4a.RequestChunk)
	TransferFinished(client string, host string)
}

type BaseSocksV4aLogger struct {
	logger zerolog.Logger
}

func NewBaseSocksV4aLogger(logger zerolog.Logger) (BaseSocksV4aLogger, error) {
	return BaseSocksV4aLogger{
		logger: logger,
	}, nil
}

func (b BaseSocksV4aLogger) ConnectRequest(client string, chunk v4a.RequestChunk) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", fmt.Sprintf("%s:%d", chunk.Domain, chunk.DestinationPort)).
		Msg("Received connect request. ")
}

func (b BaseSocksV4aLogger) ConnectFailed(client string, chunk v4a.RequestChunk) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", fmt.Sprintf("%s:%d", chunk.Domain, chunk.DestinationPort)).
		Msg("Connect failed. ")
}

func (b BaseSocksV4aLogger) ConnectSuccessful(client string, chunk v4a.RequestChunk) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", fmt.Sprintf("%s:%d", chunk.Domain, chunk.DestinationPort)).
		Msg("Connect successful. ")
}

func (b BaseSocksV4aLogger) ConnectNotAllowed(client string, chunk v4a.RequestChunk) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", fmt.Sprintf("%s:%d", chunk.Domain, chunk.DestinationPort)).
		Msg("Connect not allowed due to ruleset. ")
}

func (b BaseSocksV4aLogger) ConnectUnreachable(client string, chunk v4a.RequestChunk) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", fmt.Sprintf("%s:%d", chunk.Domain, chunk.DestinationPort)).
		Msg("Host unreachable. ")
}

func (b BaseSocksV4aLogger) ConnectTimeout(client string, chunk v4a.RequestChunk) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", fmt.Sprintf("%s:%d", chunk.Domain, chunk.DestinationPort)).
		Msg("Connect failed due to timeout. ")
}

func (b BaseSocksV4aLogger) BindRequest(client string, chunk v4a.RequestChunk) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", fmt.Sprintf("%s:%d", chunk.Domain, chunk.DestinationPort)).
		Msg("Received bind request. ")
}

func (b BaseSocksV4aLogger) BindFailed(client string, chunk v4a.RequestChunk) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", fmt.Sprintf("%s:%d", chunk.Domain, chunk.DestinationPort)).
		Msg("Bind failed. ")
}

func (b BaseSocksV4aLogger) BindSuccessful(client string, chunk v4a.RequestChunk) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", fmt.Sprintf("%s:%d", chunk.Domain, chunk.DestinationPort)).
		Msg("Bind successful. ")
}

func (b BaseSocksV4aLogger) BindNotAllowed(client string, chunk v4a.RequestChunk) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", fmt.Sprintf("%s:%d", chunk.Domain, chunk.DestinationPort)).
		Msg("Bind failed due to ruleset. ")
}

func (b BaseSocksV4aLogger) BindTimeout(client string, chunk v4a.RequestChunk) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", fmt.Sprintf("%s:%d", chunk.Domain, chunk.DestinationPort)).
		Msg("Bind failed due to timeout. ")
}

func (b BaseSocksV4aLogger) Bound(client string, host string, chunk v4a.RequestChunk) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", host).
		Str("host", fmt.Sprintf("%s:%d", chunk.Domain, chunk.DestinationPort)).
		Msg("Bound successfully. ")
}

func (b BaseSocksV4aLogger) TransferFinished(client string, host string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", host).
		Msg("Transfer finished. ")
}
