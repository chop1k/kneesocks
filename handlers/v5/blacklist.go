package v5

import (
	"socks/config"
	"socks/managers"
)

type Blacklist interface {
	IsBlacklisted(name string, address string) bool
}

type BaseBlacklist struct {
	config    config.SocksV5Config
	blacklist managers.BlacklistManager
}

func NewBaseBlacklist(config config.SocksV5Config, blacklist managers.BlacklistManager) (BaseBlacklist, error) {
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
