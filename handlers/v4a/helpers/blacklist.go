package helpers

import (
	"socks/config/v4a"
	"socks/managers"
)

type Blacklist interface {
	IsBlacklisted(address string) bool
}

type BaseBlacklist struct {
	config    v4a.RestrictionsConfig
	blacklist managers.BlacklistManager
}

func NewBaseBlacklist(config v4a.RestrictionsConfig, blacklist managers.BlacklistManager) (BaseBlacklist, error) {
	return BaseBlacklist{config: config, blacklist: blacklist}, nil
}

func (b BaseBlacklist) IsBlacklisted(address string) bool {
	blacklist, err := b.config.GetBlacklist()

	if err != nil {
		panic(err)
	}

	return b.blacklist.IsBlacklisted(blacklist, address)
}
