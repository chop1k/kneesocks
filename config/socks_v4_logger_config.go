package config

import "socks/config/tree"

type SocksV4LoggerConfig interface {
	GetConnectRequestFormat() string
	GetConnectFailedFormat() string
	GetConnectSuccessfulFormat() string
	GetBindRequestFormat() string
	GetBindFailedFormat() string
	GetBindSuccessfulFormat() string
	GetBoundFormat() string
	IsConsoleOutputEnabled() bool
	IsFileOutputEnabled() bool
	GetFilePathFormat() string
}

type BaseSocksV4LoggerConfig struct {
	config tree.Config
}

func NewBaseSocksV4LoggerConfig(config tree.Config) BaseSocksV4LoggerConfig {
	return BaseSocksV4LoggerConfig{config: config}
}

func (b BaseSocksV4LoggerConfig) GetConnectRequestFormat() string {
	return b.config.Log.Loggers.SocksV4.Formats.ConnectRequest
}

func (b BaseSocksV4LoggerConfig) GetConnectFailedFormat() string {
	return b.config.Log.Loggers.SocksV4.Formats.ConnectFailed
}

func (b BaseSocksV4LoggerConfig) GetConnectSuccessfulFormat() string {
	return b.config.Log.Loggers.SocksV4.Formats.ConnectSuccessful
}

func (b BaseSocksV4LoggerConfig) GetBindRequestFormat() string {
	return b.config.Log.Loggers.SocksV4.Formats.BindRequest
}

func (b BaseSocksV4LoggerConfig) GetBindFailedFormat() string {
	return b.config.Log.Loggers.SocksV4.Formats.BindFailed
}

func (b BaseSocksV4LoggerConfig) GetBindSuccessfulFormat() string {
	return b.config.Log.Loggers.SocksV4.Formats.BindSuccessful
}

func (b BaseSocksV4LoggerConfig) GetBoundFormat() string {
	return b.config.Log.Loggers.SocksV4.Formats.Bound
}

func (b BaseSocksV4LoggerConfig) IsConsoleOutputEnabled() bool {
	return b.config.Log.Loggers.SocksV4.Outputs.Console != nil
}

func (b BaseSocksV4LoggerConfig) IsFileOutputEnabled() bool {
	return b.config.Log.Loggers.SocksV4.Outputs.File != nil
}

func (b BaseSocksV4LoggerConfig) GetFilePathFormat() string {
	return b.config.Log.Loggers.SocksV4.Outputs.File.Path
}
