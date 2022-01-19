package config

import "socks/config/tree"

type UdpConfig interface {
	GetBindIp() string
	GetBindPort() uint16
	GetBindZone() string
	GetBufferSize() uint
}

type BaseUdpConfig struct {
	config tree.Config
}

func NewBaseUdpConfig(config tree.Config) BaseUdpConfig {
	return BaseUdpConfig{
		config: config,
	}
}

func (b BaseUdpConfig) GetBindIp() string {
	return b.config.Udp.BindIp
}

func (b BaseUdpConfig) GetBindPort() uint16 {
	return b.config.Udp.BindPort
}

func (b BaseUdpConfig) GetBindZone() string {
	return b.config.Udp.BindZone
}

func (b BaseUdpConfig) GetBufferSize() uint {
	return b.config.Udp.BufferSize
}
