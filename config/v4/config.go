package v4

import "socks/config/tree"

type Config interface {
	IsConnectAllowed() bool
	IsBindAllowed() bool
	GetRestrictions() tree.Restrictions
}

type BaseConfig struct {
	config tree.Config
}

func NewBaseConfig(config tree.Config) (BaseConfig, error) {
	return BaseConfig{config: config}, nil
}

func (b BaseConfig) IsConnectAllowed() bool {
	return b.config.SocksV4.AllowConnect
}

func (b BaseConfig) IsBindAllowed() bool {
	return b.config.SocksV4.AllowBind
}

func (b BaseConfig) GetRestrictions() tree.Restrictions {
	return b.config.SocksV4.Restrictions
}
