package helpers

import (
	"net"
	v4 "socks/config/v4"
	v42 "socks/logger/v4"
	"socks/managers"
	v43 "socks/protocol/v4"
)

type Validator interface {
	ValidateRestrictions(config v4.Config, command byte, address string, client net.Conn) bool
}

type BaseValidator struct {
	whitelist managers.WhitelistManager
	blacklist managers.BlacklistManager
	sender    v43.Sender
	logger    v42.Logger
	limiter   Limiter
}

func NewBaseValidator(
	whitelist managers.WhitelistManager,
	blacklist managers.BlacklistManager,
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

	if b.whitelist.IsWhitelisted(config.Restrictions.WhiteList, address) {
		b.sender.SendFailAndClose(config, client)

		b.logger.Restrictions.NotAllowedByWhitelist(client.RemoteAddr().String(), address)

		return false
	}

	if b.blacklist.IsBlacklisted(config.Restrictions.BlackList, address) {
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
