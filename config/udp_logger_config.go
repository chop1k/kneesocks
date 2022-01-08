package config

import "socks/config/tree"

type UdpLoggerConfig interface {
	GetPacketReceivedFormat() string
	GetPacketDeniedFormat() string
	GetPacketSentFormat() string
	IsConsoleOutputEnabled() bool
	IsFileOutputEnabled() bool
	GetFilePathFormat() string
}

type BaseUdpLoggerConfig struct {
	config tree.Config
}

func NewBaseUdpLoggerConfig(config tree.Config) BaseUdpLoggerConfig {
	return BaseUdpLoggerConfig{config: config}
}

func (b BaseUdpLoggerConfig) GetPacketReceivedFormat() string {
	return b.config.Log.Loggers.Udp.Formats.PacketReceived
}

func (b BaseUdpLoggerConfig) GetPacketDeniedFormat() string {
	return b.config.Log.Loggers.Udp.Formats.PacketDenied
}

func (b BaseUdpLoggerConfig) GetPacketSentFormat() string {
	return b.config.Log.Loggers.Udp.Formats.PacketSent
}

func (b BaseUdpLoggerConfig) IsConsoleOutputEnabled() bool {
	return b.config.Log.Loggers.Udp.Outputs.Console != nil
}

func (b BaseUdpLoggerConfig) IsFileOutputEnabled() bool {
	return b.config.Log.Loggers.Udp.Outputs.File != nil
}

func (b BaseUdpLoggerConfig) GetFilePathFormat() string {
	return b.config.Log.Loggers.Udp.Outputs.File.Path
}
