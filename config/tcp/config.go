package tcp

import "socks/config/tree"

type Config interface {
	GetBindIP() string
	GetBindPort() uint16
	GetBindZone() string
	GetClientBufferSize() uint
	GetHostBufferSize() uint
	GetExchangeDeadline() uint
}

type BaseConfig struct {
	config tree.Config
}

func NewBaseConfig(config tree.Config) (BaseConfig, error) {
	return BaseConfig{config: config}, nil
}

func (b BaseConfig) GetBindIP() string {
	return b.config.Tcp.BindIp
}

func (b BaseConfig) GetBindPort() uint16 {
	return b.config.Tcp.BindPort
}

func (b BaseConfig) GetBindZone() string {
	return b.config.Tcp.BindZone
}

func (b BaseConfig) GetClientBufferSize() uint {
	return b.config.Tcp.ClientBufferSize
}

func (b BaseConfig) GetHostBufferSize() uint {
	return b.config.Tcp.HostBufferSize
}

func (b BaseConfig) GetExchangeDeadline() uint {
	return b.config.Tcp.ExchangeDeadline
}
