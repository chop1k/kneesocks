package helpers

import (
	"net"
	"socks/internal/kneesocks/config/v4a"
	v4a2 "socks/internal/kneesocks/logger/v4a"
	"socks/internal/kneesocks/managers"
	v4a3 "socks/pkg/protocol/v4a"
)

type Validator struct {
	whitelist managers.WhitelistManager
	blacklist managers.BlacklistManager
	sender    v4a3.Sender
	logger    v4a2.Logger
	limiter   Limiter
}

func NewValidator(
	whitelist managers.WhitelistManager,
	blacklist managers.BlacklistManager,
	sender v4a3.Sender,
	logger v4a2.Logger,
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

func (b Validator) ValidateRestrictions(config v4a.Config, command byte, address string, client net.Conn) bool {
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
