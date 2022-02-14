package helpers

import (
	"net"
	v5 "socks/config/v5"
	v52 "socks/logger/v5"
	"socks/managers"
	v53 "socks/protocol/v5"
)

type Validator struct {
	whitelist managers.WhitelistManager
	blacklist managers.BlacklistManager
	sender    v53.Sender
	logger    v52.Logger
	limiter   Limiter
}

func NewValidator(
	whitelist managers.WhitelistManager,
	blacklist managers.BlacklistManager,
	sender v53.Sender,
	logger v52.Logger,
	limiter Limiter,
) (Validator, error) {
	return Validator{
		whitelist: whitelist,
		blacklist: blacklist,
		sender:    sender,
		logger:    logger,
		limiter:   limiter,
	}, nil
}

func (b Validator) ValidateRestrictions(config v5.Config, command byte, name string, addressType byte, address string, client net.Conn) bool {
	if !config.AllowIPv4 && addressType == 1 {
		b.sender.SendAddressNotSupportedAndClose(config, client)

		b.logger.Restrictions.IPv4AddressNotAllowed(client.RemoteAddr().String(), address)

		return false
	}

	if !config.AllowDomain && addressType == 3 {
		b.sender.SendAddressNotSupportedAndClose(config, client)

		b.logger.Restrictions.DomainAddressNotAllowed(client.RemoteAddr().String(), address)

		return false
	}

	if !config.AllowIPv6 && addressType == 4 {
		b.sender.SendAddressNotSupportedAndClose(config, client)

		b.logger.Restrictions.IPv6AddressNotAllowed(client.RemoteAddr().String(), address)

		return false
	}

	if !config.AllowConnect && command == 1 {
		b.sender.SendConnectionNotAllowedAndClose(config, client)

		b.logger.Restrictions.NotAllowed(client.RemoteAddr().String(), address)

		return false
	}

	if !config.AllowBind && command == 2 {
		b.sender.SendConnectionNotAllowedAndClose(config, client)

		b.logger.Restrictions.NotAllowed(client.RemoteAddr().String(), address)

		return false
	}

	if !config.AllowUdpAssociation && command == 3 {
		b.sender.SendConnectionNotAllowedAndClose(config, client)

		b.logger.Restrictions.NotAllowed(client.RemoteAddr().String(), address)

		return false
	}

	user, ok := config.Users[name]

	if !ok {
		return true
	}

	whitelisted := b.whitelist.IsWhitelisted(user.Restrictions.WhiteList, address)

	if whitelisted {
		b.sender.SendConnectionNotAllowedAndClose(config, client)

		b.logger.Restrictions.NotAllowedByWhitelist(client.RemoteAddr().String(), address)

		return false
	}

	blacklisted := b.blacklist.IsBlacklisted(user.Restrictions.BlackList, address)

	if blacklisted {
		b.sender.SendConnectionNotAllowedAndClose(config, client)

		b.logger.Restrictions.NotAllowedByBlacklist(client.RemoteAddr().String(), address)

		return false
	}

	limited := b.limiter.IsLimited(config, name)

	if limited {
		b.sender.SendFailAndClose(config, client)

		b.logger.Restrictions.NotAllowedByConnectionLimits(client.RemoteAddr().String(), address)

		return false
	}

	return true
}
