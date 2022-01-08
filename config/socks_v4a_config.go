package config

import "socks/config/tree"

type SocksV4aConfig interface {
	IsConnectAllowed() bool
	IsBindAllowed() bool
	GetConnectDeadline() uint
	GetBindDeadline() uint
}

type BaseSocksV4aConfig struct {
	config tree.Config
}

func NewBaseSocksV4aConfig(config tree.Config) BaseSocksV4aConfig {
	return BaseSocksV4aConfig{config: config}
}

func (b BaseSocksV4aConfig) IsConnectAllowed() bool {
	return b.config.SocksV4a.AllowConnect
}

func (b BaseSocksV4aConfig) IsBindAllowed() bool {
	return b.config.SocksV4a.AllowBind
}

func (b BaseSocksV4aConfig) GetConnectDeadline() uint {
	return b.config.SocksV4a.ConnectDeadline
}

func (b BaseSocksV4aConfig) GetBindDeadline() uint {
	return b.config.SocksV4a.BindDeadline
}
