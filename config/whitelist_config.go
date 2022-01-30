package config

import "socks/config/tree"

type WhitelistConfig interface {
	GetWhitelist(name string) ([]string, error)
}

type BaseWhitelistConfig struct {
	config tree.Config
}

func NewBaseWhitelistConfig(config tree.Config) BaseWhitelistConfig {
	return BaseWhitelistConfig{
		config: config,
	}
}

func (b BaseWhitelistConfig) GetWhitelist(name string) ([]string, error) {
	if b.config.SocksV5 == nil {

	}

	user, ok := b.config.SocksV5.Users[name]

	if !ok {

	}

	return user.Restrictions.WhiteList, nil
}
