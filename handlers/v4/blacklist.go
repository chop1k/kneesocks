package v4

import (
	"socks/config/v4"
	"socks/managers"
)

type Blacklist interface {
	IsBlacklisted(address string) bool
}

type BaseBlacklist struct {
	config    v4.Config
	blacklist managers.BlacklistManager
}

func NewBaseBlacklist(config v4.Config, blacklist managers.BlacklistManager) (BaseBlacklist, error) {
	return BaseBlacklist{config: config, blacklist: blacklist}, nil
}

func (b BaseBlacklist) IsBlacklisted(address string) bool {
	restrictions := b.config.GetRestrictions()

	return b.blacklist.IsBlacklisted(restrictions.BlackList, address)
}