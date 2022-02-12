package helpers

import (
	"socks/config/v4a"
	"socks/managers"
)

type Whitelist interface {
	IsWhitelisted(address string) (bool, error)
}

type BaseWhitelist struct {
	config    v4a.RestrictionsConfig
	whitelist managers.WhitelistManager
}

func NewBaseWhitelist(config v4a.RestrictionsConfig, whitelist managers.WhitelistManager) (BaseWhitelist, error) {
	return BaseWhitelist{config: config, whitelist: whitelist}, nil
}

func (b BaseWhitelist) IsWhitelisted(address string) (bool, error) {
	whitelist, err := b.config.GetWhitelist()

	if err != nil {
		return false, err
	}

	return b.whitelist.IsWhitelisted(whitelist, address), nil
}
