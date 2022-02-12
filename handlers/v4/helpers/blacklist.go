package helpers

import (
	"socks/config/v4"
	"socks/managers"
)

type Blacklist interface {
	IsBlacklisted(config v4.Config, address string) bool
}

type BaseBlacklist struct {
	blacklist managers.BlacklistManager
}

func NewBaseBlacklist(
	blacklist managers.BlacklistManager,
) (BaseBlacklist, error) {
	return BaseBlacklist{
		blacklist: blacklist,
	}, nil
}

func (b BaseBlacklist) IsBlacklisted(config v4.Config, address string) bool {
	return b.blacklist.IsBlacklisted(config.Restrictions.BlackList, address)
}
