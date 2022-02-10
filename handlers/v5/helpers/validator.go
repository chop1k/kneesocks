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
	ipV4Allowed, ipV4Err := b.config.IsIPv4Allowed()

	if ipV4Err != nil {
		panic(ipV4Err)
	}

	if !ipV4Allowed && addressType == 1 {
		b.sender.SendAddressNotSupportedAndClose(client)

		b.logger.Restrictions.IPv4AddressNotAllowed(client.RemoteAddr().String(), address)

		return false
	}

	domainAllowed, domainErr := b.config.IsIPv4Allowed()

	if domainErr != nil {
		panic(domainErr)
	}

	if !domainAllowed && addressType == 3 {
		b.sender.SendAddressNotSupportedAndClose(client)

		b.logger.Restrictions.DomainAddressNotAllowed(client.RemoteAddr().String(), address)

		return false
	}

	ipV6Allowed, ipV6Err := b.config.IsIPv4Allowed()

	if ipV6Err != nil {
		panic(ipV6Err)
	}

	if !ipV6Allowed && addressType == 4 {
		b.sender.SendAddressNotSupportedAndClose(client)

		b.logger.Restrictions.IPv6AddressNotAllowed(client.RemoteAddr().String(), address)

		return false
	}

	connectAllowed, connectErr := b.config.IsConnectAllowed()

	if connectErr != nil {
		panic(connectErr)
	}

	if !connectAllowed && command == 1 {
		b.sender.SendConnectionNotAllowedAndClose(client)

		b.logger.Restrictions.NotAllowed(client.RemoteAddr().String(), address)

		return false
	}

	bindAllowed, bindErr := b.config.IsBindAllowed()

	if bindErr != nil {
		panic(bindErr)
	}

	if !bindAllowed && command == 2 {
		b.sender.SendConnectionNotAllowedAndClose(client)

		b.logger.Restrictions.NotAllowed(client.RemoteAddr().String(), address)

		return false
	}

	associationAllowed, associationErr := b.config.IsUdpAssociationAllowed()

	if associationErr != nil {
		panic(associationErr)
	}

	if !associationAllowed && command == 3 {
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
