package helpers

import (
	"socks/config/v5"
	"socks/managers"
)

type Whitelist interface {
	IsWhitelisted(name string, address string) bool
}

type BaseWhitelist struct {
	config    v5.Config
	whitelist managers.WhitelistManager
}

func NewBaseWhitelist(config v5.Config, whitelist managers.WhitelistManager) (BaseWhitelist, error) {
	return BaseWhitelist{config: config, whitelist: whitelist}, nil
}

func (b BaseWhitelist) IsWhitelisted(name string, address string) bool {
	users := b.config.GetUsers()

	user, ok := users[name]

	if !ok {
		return false
	}

	return b.whitelist.IsWhitelisted(user.Restrictions.WhiteList, address)
}
