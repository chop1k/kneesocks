package config

import "socks/config/tree"

type BlacklistConfig interface {
	GetBlacklist() []string
}

type BaseBlacklistConfig struct {
	config tree.Config
}

func NewBaseBlacklistConfig(config tree.Config) BaseBlacklistConfig {
	return BaseBlacklistConfig{
		config: config,
	}
}

func (b BaseBlacklistConfig) GetBlacklist() []string {
	return b.config.BlackList
}
