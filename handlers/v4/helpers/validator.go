package helpers

import (
	"net"
	v4 "socks/config/v4"
	v42 "socks/logger/v4"
	v43 "socks/protocol/v4"
)

type Validator interface {
	ValidateRestrictions(config v4.Config, command byte, address string, client net.Conn) bool
}

type BaseValidator struct {
	whitelist Whitelist
	blacklist Blacklist
	sender    v43.Sender
	logger    v42.Logger
	limiter   Limiter
}

func NewBaseValidator(
	whitelist Whitelist,
	blacklist Blacklist,
	sender v43.Sender,
	logger v42.Logger,
	limiter Limiter,
) (BaseValidator, error) {
	return BaseValidator{
		whitelist: whitelist,
		blacklist: blacklist,
		sender:    sender,
		logger:    logger,
		limiter:   limiter,
	}, nil
}

func (b BaseValidator) ValidateRestrictions(config v4.Config, command byte, address string, client net.Conn) bool {
	if !config.AllowConnect && command == 1 {
		b.sender.SendFailAndClose(config, client)

		b.logger.Restrictions.NotAllowed(client.RemoteAddr().String(), address)

		return false
	}

	if !config.AllowBind && command == 2 {
		b.sender.SendFailAndClose(config, client)

		b.logger.Restrictions.NotAllowed(client.RemoteAddr().String(), address)

		return false
	}

	whitelisted := b.whitelist.IsWhitelisted(config, address)

	if whitelisted {
		b.sender.SendFailAndClose(config, client)

		b.logger.Restrictions.NotAllowedByWhitelist(client.RemoteAddr().String(), address)

		return false
	}

	blacklisted := b.blacklist.IsBlacklisted(config, address)

	if blacklisted {
		b.sender.SendFailAndClose(config, client)

		b.logger.Restrictions.NotAllowedByBlacklist(client.RemoteAddr().String(), address)

		return false
	}

	limited := b.limiter.IsLimited(config)

	if limited {
		b.sender.SendFailAndClose(config, client)

		b.logger.Restrictions.NotAllowedByConnectionLimits(client.RemoteAddr().String(), address)

		return false
	}

	return true
}
