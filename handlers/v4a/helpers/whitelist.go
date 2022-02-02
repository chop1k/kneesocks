package helpers

import (
	"socks/config/v4a"
	"socks/managers"
)

type Whitelist interface {
	IsWhitelisted(address string) bool
}

type BaseWhitelist struct {
	config    v4a.Config
	whitelist managers.WhitelistManager
}

func NewBaseWhitelist(config v4a.Config, whitelist managers.WhitelistManager) (BaseWhitelist, error) {
	return BaseWhitelist{config: config, whitelist: whitelist}, nil
}

func (b BaseWhitelist) IsWhitelisted(address string) bool {
	restrictions := b.config.GetRestrictions()

	return b.whitelist.IsWhitelisted(restrictions.WhiteList, address)
}
