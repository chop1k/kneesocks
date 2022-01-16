package config

import "socks/config/tree"

type TcpConfig interface {
	GetBindIP() string
	GetBindPort() int
	GetBindZone() string
	GetClientBufferSize() uint
	GetHostBufferSize() uint
	GetExchangeDeadline() uint
}

type BaseTcpConfig struct {
	config tree.Config
}

func NewBaseTcpConfig(config tree.Config) BaseTcpConfig {
	return BaseTcpConfig{config: config}
}

func (b BaseTcpConfig) GetBindIP() string {
	return b.config.Tcp.BindIp
}

func (b BaseTcpConfig) GetBindPort() int {
	return int(b.config.Tcp.BindPort)
}

func (b BaseTcpConfig) GetBindZone() string {
	return b.config.Tcp.BindZone
}

func (b BaseTcpConfig) GetClientBufferSize() uint {
	return b.config.Tcp.ClientBufferSize
}

func (b BaseTcpConfig) GetHostBufferSize() uint {
	return b.config.Tcp.HostBufferSize
}

func (b BaseTcpConfig) GetExchangeDeadline() uint {
	return b.config.Tcp.ExchangeDeadline
}
