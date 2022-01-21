package logger

import (
	"github.com/rs/zerolog"
)

type SocksV5Logger interface {
	ConnectRequest(client string, address string)
	ConnectFailed(client string, address string)
	ConnectSuccessful(client string, address string)
	ConnectNotAllowed(client string, address string)
	ConnectNotAllowedByWhitelist(client string, address string)
	ConnectNotAllowedByBlacklist(client string, address string)
	ConnectTimeout(client string, address string)
	ConnectRefused(client string, address string)
	ConnectHostUnreachable(client string, address string)
	ConnectNetworkUnreachable(client string, address string)
	BindRequest(client string, address string)
	BindFailed(client string, address string)
	BindSuccessful(client string, address string)
	BindNotAllowed(client string, address string)
	BindNotAllowedByWhitelist(client string, address string)
	BindNotAllowedByBlacklist(client string, address string)
	BindTimeout(client string, address string)
	Bound(client string, host string)
	UdpAssociationRequest(client string)
	UdpAssociationSuccessful(client string, address string)
	UdpAssociationFailed(client string, address string)
	UdpAssociationNotAllowed(client string)
	AuthenticationSuccessful(client string, name string)
	AuthenticationFailed(client string)
	TransferFinished(client string, host string)
	ParseError(client string, err error)
	UnknownError(client string, address string, err error)
	AddressParsingError(client string, host string, address string, err error)
	UdpAddressParsingError(client string, err error)
	AddressDeterminationError(client string, host string, address string, err error)
	InvalidAddressTypeError(client string, addressType byte, address string)
	UnknownCommandError(client string, command byte, address string)
	SelectMethodsError(client string, err error)
	ReceiveRequestError(client string, err error)
	PasswordResponseError(client string, user string, err error)
	ParseMethodsError(client string, err error)
	ReceiveHostError(client string, address string, err error)
	SendClientError(client string, host string, address string, err error)
	BindError(client string, address string, err error)
	IPv4AddressNotAllowed(client string, address string)
	DomainAddressNotAllowed(client string, address string)
	IPv6AddressNotAllowed(client string, address string)
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

func (b BaseSocksV5Logger) ConnectNotAllowedByWhitelist(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Connect not allowed due to whitelist.")
}

func (b BaseSocksV5Logger) ConnectNotAllowedByBlacklist(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Connect not allowed due to blacklist.")
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

func (b BaseSocksV5Logger) BindNotAllowedByWhitelist(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Bind not allowed due to whitelist.")
}

func (b BaseSocksV5Logger) BindNotAllowedByBlacklist(client string, address string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("host", address).
		Msg("Bind not allowed due to blacklist.")
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

func (b BaseSocksV5Logger) AuthenticationSuccessful(client string, name string) {
	e := b.logger.Info()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("name", name).
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

func (b BaseSocksV5Logger) ParseError(client string, err error) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Err(err).
		Msg("Cannot parse v5 request due to error. ")
}

func (b BaseSocksV5Logger) UnknownError(client string, address string, err error) {
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

func (b BaseSocksV5Logger) AddressParsingError(client string, host string, address string, err error) {
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

func (b BaseSocksV5Logger) UdpAddressParsingError(client string, err error) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Err(err).
		Msg("Cannot handle v5 request due to address parsing error. ")
}

func (b BaseSocksV5Logger) AddressDeterminationError(client string, host string, address string, err error) {
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

func (b BaseSocksV5Logger) InvalidAddressTypeError(client string, addressType byte, address string) {
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

func (b BaseSocksV5Logger) UnknownCommandError(client string, command byte, address string) {
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

func (b BaseSocksV5Logger) SelectMethodsError(client string, err error) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Err(err).
		Msg("Cannot handle v5 request due to methods selection error. ")
}

func (b BaseSocksV5Logger) ReceiveRequestError(client string, err error) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Err(err).
		Msg("Cannot handle v5 request due to receive request error. ")
}

func (b BaseSocksV5Logger) PasswordResponseError(client string, user string, err error) {
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

func (b BaseSocksV5Logger) ParseMethodsError(client string, err error) {
	e := b.logger.Error()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Err(err).
		Msg("Cannot handle v5 request due to parse methods error. ")
}

func (b BaseSocksV5Logger) ReceiveHostError(client string, address string, err error) {
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

func (b BaseSocksV5Logger) SendClientError(client string, host string, address string, err error) {
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

func (b BaseSocksV5Logger) BindError(client string, address string, err error) {
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

func (b BaseSocksV5Logger) IPv4AddressNotAllowed(client string, address string) {
	e := b.logger.Warn()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("address", address).
		Msg("IPv4 address not allowed by ruleset. ")
}

func (b BaseSocksV5Logger) DomainAddressNotAllowed(client string, address string) {
	e := b.logger.Warn()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("address", address).
		Msg("Domain address not allowed by ruleset. ")
}

func (b BaseSocksV5Logger) IPv6AddressNotAllowed(client string, address string) {
	e := b.logger.Warn()

	if !e.Enabled() {
		return
	}

	e.
		Str("client", client).
		Str("address", address).
		Msg("IPv6 address not allowed by ruleset. ")
}
