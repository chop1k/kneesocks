package config

import "socks/config/tree"

type WhitelistConfig interface {
	GetWhitelist() []string
}

type BaseWhitelistConfig struct {
	config tree.Config
}

func NewBaseWhitelistConfig(config tree.Config) BaseWhitelistConfig {
	return BaseWhitelistConfig{
		config: config,
	}
}

func (b BaseWhitelistConfig) GetWhitelist() []string {
	return b.config.WhiteList
}
