package config

import "socks/config/tree"

type Config interface {
	IsSocksV4Enabled() bool
	IsSocksV4aEnabled() bool
	IsSocksV5Enabled() bool
	IsUnixEnabled() bool
}

type BaseConfig struct {
	config tree.Config
}

func NewBaseConfig(config tree.Config) BaseConfig {
	return BaseConfig{config: config}
}

func (b BaseConfig) IsSocksV4Enabled() bool {
	return b.config.SocksV4 != nil
}

func (b BaseConfig) IsSocksV4aEnabled() bool {
	return b.config.SocksV4a != nil
}

func (b BaseConfig) IsSocksV5Enabled() bool {
	return b.config.SocksV5 != nil
}

func (b BaseConfig) IsUnixEnabled() bool {
	return b.config.Unix != nil
}
