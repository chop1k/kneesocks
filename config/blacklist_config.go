package config

import "socks/config/tree"

type BlacklistConfig interface {
	GetBlacklist(name string) ([]string, error)
}

type BaseBlacklistConfig struct {
	config tree.Config
}

func NewBaseBlacklistConfig(config tree.Config) BaseBlacklistConfig {
	return BaseBlacklistConfig{
		config: config,
	}
}

func (b BaseBlacklistConfig) GetBlacklist(name string) ([]string, error) {
	if b.config.SocksV5 == nil {

	}

	user, ok := b.config.SocksV5.Users[name]

	if !ok {

	}

	return user.Restrictions.BlackList, nil
}
