package helpers

import (
	"net"
	v4 "socks/config/v4"
	v42 "socks/logger/v4"
)

type Validator interface {
	ValidateRestrictions(command byte, address string, client net.Conn) bool
}

type BaseValidator struct {
	config    v4.Config
	whitelist Whitelist
	blacklist Blacklist
	sender    Sender
	logger    v42.Logger
}

func NewBaseValidator(
	config v4.Config,
	whitelist Whitelist,
	blacklist Blacklist,
	sender Sender,
	logger v42.Logger,
) (BaseValidator, error) {
	return BaseValidator{
		config:    config,
		whitelist: whitelist,
		blacklist: blacklist,
		sender:    sender,
		logger:    logger,
	}, nil
}

func (b BaseValidator) ValidateRestrictions(command byte, address string, client net.Conn) bool {
	if !b.config.IsConnectAllowed() && command == 1 {
		b.sender.SendFailAndClose(client)

		b.logger.Restrictions.NotAllowed(client.RemoteAddr().String(), address)

		return false
	}

	if !b.config.IsBindAllowed() && command == 2 {
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

	return true
}
