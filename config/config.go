package config

import (
	"github.com/Jeffail/gabs"
)

type Config interface {
	IsSocksV4Enabled() bool
	IsSocksV4aEnabled() bool
	IsSocksV5Enabled() bool
}

type BaseConfig struct {
	config gabs.Container
}

func NewBaseConfig(config gabs.Container) BaseConfig {
	return BaseConfig{config: config}
}

func (b BaseConfig) IsSocksV4Enabled() bool {
	return b.config.ExistsP("SocksV4")
}

func (b BaseConfig) IsSocksV4aEnabled() bool {
	return b.config.ExistsP("SocksV4a")
}

func (b BaseConfig) IsSocksV5Enabled() bool {
	return b.config.ExistsP("SocksV5")
}
