package config

import "socks/config/tree"

type SocksV4Config interface {
	IsConnectAllowed() bool
	IsBindAllowed() bool
	GetConnectDeadline() uint
	GetBindDeadline() uint
	GetRestrictions() tree.Restrictions
}

type BaseSocksV4Config struct {
	config tree.Config
}

func NewBaseSocksV4Config(config tree.Config) BaseSocksV4Config {
	return BaseSocksV4Config{config: config}
}

func (b BaseSocksV4Config) IsConnectAllowed() bool {
	return b.config.SocksV4.AllowConnect
}

func (b BaseSocksV4Config) IsBindAllowed() bool {
	return b.config.SocksV4.AllowBind
}

func (b BaseSocksV4Config) GetConnectDeadline() uint {
	return b.config.SocksV4.ConnectDeadline
}

func (b BaseSocksV4Config) GetBindDeadline() uint {
	return b.config.SocksV4.BindDeadline
}

func (b BaseSocksV4Config) GetRestrictions() tree.Restrictions {
	return b.config.SocksV4.Restrictions
}
