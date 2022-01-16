package server

import (
	"socks/config/tree"
)

type WhitelistManager interface {
	IsWhitelisted(addrType byte, addr string) bool
}

type BaseWhitelist struct {
	config tree.Config
}

func NewBaseWhitelist(config tree.Config) BaseWhitelist {
	return BaseWhitelist{
		config: config,
	}
}

func (b BaseWhitelist) IsWhitelisted(addrType byte, addr string) bool {
	panic("b")
}
