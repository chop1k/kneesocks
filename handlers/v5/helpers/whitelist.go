package helpers

import (
	"socks/config/v5"
	"socks/managers"
)

type Whitelist interface {
	IsWhitelisted(name string, address string) bool
}

type BaseWhitelist struct {
	config    v5.UsersConfig
	whitelist managers.WhitelistManager
}

func NewBaseWhitelist(config v5.UsersConfig, whitelist managers.WhitelistManager) (BaseWhitelist, error) {
	return BaseWhitelist{config: config, whitelist: whitelist}, nil
}

func (b BaseWhitelist) IsWhitelisted(name string, address string) bool {
	restrictions, err := b.config.GetRestrictions(name)

	if err != nil {
		panic(err)
	}

	//if !ok {
	//	return false
	//}

	return b.whitelist.IsWhitelisted(restrictions.WhiteList, address)
}
