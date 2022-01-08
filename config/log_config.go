package config

import (
	"socks/config/tree"
)

type LogConfig interface {
	IsTcpLoggerEnabled() bool
	IsUdpLoggerEnabled() bool
	IsSocksV4LoggerEnabled() bool
	IsSocksV4aLoggerEnabled() bool
	IsSocksV5LoggerEnabled() bool
	IsUnixLoggerEnabled() bool
	IsErrorsLoggerEnabled() bool
	GetReplacer() string
}

type BaseLogConfig struct {
	config tree.Config
}

func NewBaseLogConfig(config tree.Config) BaseLogConfig {
	return BaseLogConfig{config: config}
}

func (b BaseLogConfig) IsTcpLoggerEnabled() bool {
	return b.config.Log.Loggers.Tcp != nil
}

func (b BaseLogConfig) IsUdpLoggerEnabled() bool {
	return b.config.Log.Loggers.Udp != nil
}

func (b BaseLogConfig) IsSocksV4LoggerEnabled() bool {
	return b.config.Log.Loggers.SocksV4 != nil
}

func (b BaseLogConfig) IsSocksV4aLoggerEnabled() bool {
	return b.config.Log.Loggers.SocksV4a != nil
}

func (b BaseLogConfig) IsSocksV5LoggerEnabled() bool {
	return b.config.Log.Loggers.SocksV5 != nil
}

func (b BaseLogConfig) IsUnixLoggerEnabled() bool {
	return b.config.Log.Loggers.Unix != nil
}

func (b BaseLogConfig) IsErrorsLoggerEnabled() bool {
	return b.config.Log.Loggers.Errors != nil
}

func (b BaseLogConfig) GetReplacer() string {
	return b.config.Log.Formats.Replacer
}
