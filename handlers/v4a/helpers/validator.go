package helpers

import (
	"net"
	"socks/config/v4a"
	v4a2 "socks/logger/v4a"
	"socks/managers"
	v4a3 "socks/protocol/v4a"
)

type Validator interface {
	ValidateRestrictions(config v4a.Config, command byte, address string, client net.Conn) bool
}

type BaseValidator struct {
	whitelist managers.WhitelistManager
	blacklist managers.BlacklistManager
	sender    v4a3.Sender
	logger    v4a2.Logger
	limiter   Limiter
}

func NewBaseValidator(
	whitelist managers.WhitelistManager,
	blacklist managers.BlacklistManager,
	sender v4a3.Sender,
	logger v4a2.Logger,
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

func (b BaseValidator) ValidateRestrictions(config v4a.Config, command byte, address string, client net.Conn) bool {
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

	if b.limiter.IsLimited(config) {
		b.sender.SendFailAndClose(config, client)

		b.logger.Restrictions.NotAllowedByConnectionLimits(client.RemoteAddr().String(), address)

		return false
	}

	return true
}
