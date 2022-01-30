package v4a

import (
	"socks/config"
	"socks/managers"
)

type Blacklist interface {
	IsBlacklisted(address string) bool
}

type BaseBlacklist struct {
	config    config.SocksV4aConfig
	blacklist managers.BlacklistManager
}

func NewBaseBlacklist(config config.SocksV4aConfig, blacklist managers.BlacklistManager) (BaseBlacklist, error) {
	return BaseBlacklist{config: config, blacklist: blacklist}, nil
}

func (b BaseBlacklist) IsBlacklisted(address string) bool {
	restrictions := b.config.GetRestrictions()

	return b.blacklist.IsBlacklisted(restrictions.BlackList, address)
}
