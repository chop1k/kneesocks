package helpers

import (
	"socks/config/v4"
	"socks/managers"
)

type Whitelist interface {
	IsWhitelisted(address string) bool
}

type BaseWhitelist struct {
	config    v4.RestrictionsConfig
	whitelist managers.WhitelistManager
}

func NewBaseWhitelist(
	config v4.RestrictionsConfig,
	whitelist managers.WhitelistManager,
) (BaseWhitelist, error) {
	return BaseWhitelist{
		config:    config,
		whitelist: whitelist,
	}, nil
}

func (b BaseWhitelist) IsWhitelisted(address string) bool {
	whitelist, err := b.config.GetWhitelist()

	if err != nil {
		panic(err)
	}

	return b.whitelist.IsWhitelisted(whitelist, address)
}
