package helpers

import (
	"socks/config/v5"
	"socks/managers"
)

type Blacklist interface {
	IsBlacklisted(name string, address string) bool
}

type BaseBlacklist struct {
	config    v5.Config
	blacklist managers.BlacklistManager
}

func NewBaseBlacklist(config v5.Config, blacklist managers.BlacklistManager) (BaseBlacklist, error) {
	return BaseBlacklist{config: config, blacklist: blacklist}, nil
}

func (b BaseBlacklist) IsBlacklisted(name string, address string) bool {
	users := b.config.GetUsers()

	user, ok := users[name]

	if !ok {
		return false
	}

	return b.blacklist.IsBlacklisted(user.Restrictions.BlackList, address)
}
