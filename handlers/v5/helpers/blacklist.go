package helpers

import (
	"socks/config/v5"
	"socks/managers"
)

type Blacklist interface {
	IsBlacklisted(name string, address string) bool
}

type BaseBlacklist struct {
	config    v5.UsersConfig
	blacklist managers.BlacklistManager
}

func NewBaseBlacklist(config v5.UsersConfig, blacklist managers.BlacklistManager) (BaseBlacklist, error) {
	return BaseBlacklist{config: config, blacklist: blacklist}, nil
}

func (b BaseBlacklist) IsBlacklisted(name string, address string) bool {
	restrictions, err := b.config.GetRestrictions(name)

	if err != nil {
		panic(err)
	}

	//if !ok {
	//	return false
	//}

	return b.blacklist.IsBlacklisted(restrictions.BlackList, address)
}
