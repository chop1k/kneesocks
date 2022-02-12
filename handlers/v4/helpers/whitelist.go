package helpers

import (
	"socks/config/v4"
	"socks/managers"
)

type Whitelist interface {
	IsWhitelisted(config v4.Config, address string) bool
}

type BaseWhitelist struct {
	whitelist managers.WhitelistManager
}

func NewBaseWhitelist(
	whitelist managers.WhitelistManager,
) (BaseWhitelist, error) {
	return BaseWhitelist{
		whitelist: whitelist,
	}, nil
}

func (b BaseWhitelist) IsWhitelisted(config v4.Config, address string) bool {
	return b.whitelist.IsWhitelisted(config.Restrictions.WhiteList, address)
}
