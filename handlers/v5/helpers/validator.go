package helpers

import (
	"net"
	v5 "socks/config/v5"
	v52 "socks/logger/v5"
	v53 "socks/protocol/v5"
)

type Validator interface {
	ValidateRestrictions(command byte, name string, addressType byte, address string, client net.Conn) bool
}

type BaseValidator struct {
	config    v5.Config
	whitelist Whitelist
	blacklist Blacklist
	sender    v53.Sender
	logger    v52.Logger
}

func NewBaseValidator(
	config v5.Config,
	whitelist Whitelist,
	blacklist Blacklist,
	sender v53.Sender,
	logger v52.Logger,
) (BaseValidator, error) {
	return BaseValidator{
		config:    config,
		whitelist: whitelist,
		blacklist: blacklist,
		sender:    sender,
		logger:    logger,
	}, nil
}

func (b BaseValidator) ValidateRestrictions(command byte, name string, addressType byte, address string, client net.Conn) bool {
	if !b.config.IsIPv4Allowed() && addressType == 1 {
		b.sender.SendAddressNotSupportedAndClose(client)

		b.logger.Restrictions.IPv4AddressNotAllowed(client.RemoteAddr().String(), address)

		return false
	}

	if !b.config.IsDomainAllowed() && addressType == 3 {
		b.sender.SendAddressNotSupportedAndClose(client)

		b.logger.Restrictions.DomainAddressNotAllowed(client.RemoteAddr().String(), address)

		return false
	}

	if !b.config.IsIPv6Allowed() && addressType == 4 {
		b.sender.SendAddressNotSupportedAndClose(client)

		b.logger.Restrictions.IPv6AddressNotAllowed(client.RemoteAddr().String(), address)

		return false
	}

	if !b.config.IsConnectAllowed() && command == 1 {
		b.sender.SendConnectionNotAllowedAndClose(client)

		b.logger.Restrictions.NotAllowed(client.RemoteAddr().String(), address)

		return false
	}

	if !b.config.IsBindAllowed() && command == 2 {
		b.sender.SendConnectionNotAllowedAndClose(client)

		b.logger.Restrictions.NotAllowed(client.RemoteAddr().String(), address)

		return false
	}

	if !b.config.IsUdpAssociationAllowed() && command == 3 {
		b.sender.SendConnectionNotAllowedAndClose(client)

		b.logger.Restrictions.NotAllowed(client.RemoteAddr().String(), address)

		return false
	}

	whitelisted := b.whitelist.IsWhitelisted(name, address)

	if whitelisted {
		b.sender.SendConnectionNotAllowedAndClose(client)

		b.logger.Restrictions.NotAllowedByWhitelist(client.RemoteAddr().String(), address)

		return false
	}

	blacklisted := b.blacklist.IsBlacklisted(name, address)

	if blacklisted {
		b.sender.SendConnectionNotAllowedAndClose(client)

		b.logger.Restrictions.NotAllowedByBlacklist(client.RemoteAddr().String(), address)

		return false
	}

	return true
}
