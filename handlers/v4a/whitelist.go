package v4a

import (
	"socks/config"
	"socks/managers"
)

type Whitelist interface {
	IsWhitelisted(address string) bool
}

type BaseWhitelist struct {
	config    config.SocksV4aConfig
	whitelist managers.WhitelistManager
}

func NewBaseWhitelist(config config.SocksV4aConfig, whitelist managers.WhitelistManager) (BaseWhitelist, error) {
	return BaseWhitelist{config: config, whitelist: whitelist}, nil
}

func (b BaseWhitelist) IsWhitelisted(address string) bool {
	restrictions := b.config.GetRestrictions()

	return b.whitelist.IsWhitelisted(restrictions.WhiteList, address)
}
