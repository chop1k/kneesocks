package v4a

import "socks/config/tree"

type Config interface {
	IsConnectAllowed() bool
	IsBindAllowed() bool
	GetConnectDeadline() uint
	GetBindDeadline() uint
	GetRestrictions() tree.Restrictions
}

type BaseConfig struct {
	config tree.Config
}

func NewBaseConfig(config tree.Config) (BaseConfig, error) {
	return BaseConfig{config: config}, nil
}

func (b BaseConfig) IsConnectAllowed() bool {
	return b.config.SocksV4a.AllowConnect
}

func (b BaseConfig) IsBindAllowed() bool {
	return b.config.SocksV4a.AllowBind
}

func (b BaseConfig) GetConnectDeadline() uint {
	return b.config.SocksV4a.ConnectDeadline
}

func (b BaseConfig) GetBindDeadline() uint {
	return b.config.SocksV4a.BindDeadline
}

func (b BaseConfig) GetRestrictions() tree.Restrictions {
	return b.config.SocksV4a.Restrictions
}
