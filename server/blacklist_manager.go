package server

import (
	"socks/config/tree"
)

type BlacklistManager interface {
	IsBlacklisted(addrType byte, addr string) bool
}

type BaseBlacklist struct {
	config tree.Config
}

func NewBaseBlacklist(config tree.Config) BaseBlacklist {
	return BaseBlacklist{
		config: config,
	}
}

func (b BaseBlacklist) IsBlacklisted(addrType byte, addr string) bool {
	panic("a")
}
