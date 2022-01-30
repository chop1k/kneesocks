package udp

import "socks/config/tree"

type Config interface {
	GetBindIp() string
	GetBindPort() uint16
	GetBindZone() string
	GetBufferSize() uint
}

type BaseConfig struct {
	config tree.Config
}

func NewBaseConfig(config tree.Config) (BaseConfig, error) {
	return BaseConfig{
		config: config,
	}, nil
}

func (b BaseConfig) GetBindIp() string {
	return b.config.Udp.BindIp
}

func (b BaseConfig) GetBindPort() uint16 {
	return b.config.Udp.BindPort
}

func (b BaseConfig) GetBindZone() string {
	return b.config.Udp.BindZone
}

func (b BaseConfig) GetBufferSize() uint {
	return b.config.Udp.BufferSize
}
