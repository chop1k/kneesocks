package helpers

import (
	"net"
	"socks/config/v4a"
	v4a2 "socks/logger/v4a"
	v4a3 "socks/protocol/v4a"
)

type Validator interface {
	ValidateRestrictions(command byte, address string, client net.Conn) bool
}

type BaseValidator struct {
	config    v4a.Config
	whitelist Whitelist
	blacklist Blacklist
	sender    v4a3.Sender
	logger    v4a2.Logger
	limiter   Limiter
}

func NewBaseValidator(
	config v4a.Config,
	whitelist Whitelist,
	blacklist Blacklist,
	sender v4a3.Sender,
	logger v4a2.Logger,
	limiter Limiter,
) (BaseValidator, error) {
	return BaseValidator{
		config:    config,
		whitelist: whitelist,
		blacklist: blacklist,
		sender:    sender,
		logger:    logger,
		limiter:   limiter,
	}, nil
}

func (b BaseValidator) ValidateRestrictions(command byte, address string, client net.Conn) bool {
	connectAllowed, connectErr := b.config.IsConnectAllowed()

	if connectErr != nil {
		panic(connectErr)
	}

	if !connectAllowed && command == 1 {
		b.sender.SendFailAndClose(client)

		b.logger.Restrictions.NotAllowed(client.RemoteAddr().String(), address)

		return false
	}

	bindAllowed, bindErr := b.config.IsBindAllowed()

	if bindErr != nil {
		panic(bindErr)
	}

	if !bindAllowed && command == 2 {
		b.sender.SendFailAndClose(client)

		b.logger.Restrictions.NotAllowed(client.RemoteAddr().String(), address)

		return false
	}

	whitelisted := b.whitelist.IsWhitelisted(address)

	if whitelisted {
		b.sender.SendFailAndClose(client)

		b.logger.Restrictions.NotAllowedByWhitelist(client.RemoteAddr().String(), address)

		return false
	}

	blacklisted := b.blacklist.IsBlacklisted(address)

	if blacklisted {
		b.sender.SendFailAndClose(client)

		b.logger.Restrictions.NotAllowedByBlacklist(client.RemoteAddr().String(), address)

		return false
	}

	limited := b.limiter.IsLimited()

	if limited {
		b.sender.SendFailAndClose(client)

		b.logger.Restrictions.NotAllowedByBlacklist(client.RemoteAddr().String(), address) // TODO: log

		return false
	}

	return true
}
