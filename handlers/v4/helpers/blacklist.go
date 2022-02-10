package helpers

import (
	"socks/config/v4"
	"socks/managers"
)

type Blacklist interface {
	IsBlacklisted(address string) bool
}

type BaseBlacklist struct {
	config    v4.RestrictionsConfig
	blacklist managers.BlacklistManager
}

func NewBaseBlacklist(
	config v4.RestrictionsConfig,
	blacklist managers.BlacklistManager,
) (BaseBlacklist, error) {
	return BaseBlacklist{
		config:    config,
		blacklist: blacklist,
	}, nil
}

func (b BaseBlacklist) IsBlacklisted(address string) bool {
	blacklist, err := b.config.GetBlacklist()

	if err != nil {
		return false
	}

	return b.blacklist.IsBlacklisted(blacklist, address)
}
