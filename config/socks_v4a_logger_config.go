package config

import "socks/config/tree"

type SocksV4aLoggerConfig interface {
	SocksV4LoggerConfig
}

type BaseSocksV4aLoggerConfig struct {
	config tree.Config
}

func NewBaseSocksV4aLoggerConfig(config tree.Config) BaseSocksV4aLoggerConfig {
	return BaseSocksV4aLoggerConfig{config: config}
}

func (b BaseSocksV4aLoggerConfig) GetConnectRequestFormat() string {
	return b.config.Log.Loggers.SocksV4a.Formats.ConnectRequest
}

func (b BaseSocksV4aLoggerConfig) GetConnectFailedFormat() string {
	return b.config.Log.Loggers.SocksV4a.Formats.ConnectFailed
}

func (b BaseSocksV4aLoggerConfig) GetConnectSuccessfulFormat() string {
	return b.config.Log.Loggers.SocksV4a.Formats.ConnectSuccessful
}

func (b BaseSocksV4aLoggerConfig) GetBindRequestFormat() string {
	return b.config.Log.Loggers.SocksV4a.Formats.BindRequest
}

func (b BaseSocksV4aLoggerConfig) GetBindFailedFormat() string {
	return b.config.Log.Loggers.SocksV4a.Formats.BindFailed
}

func (b BaseSocksV4aLoggerConfig) GetBindSuccessfulFormat() string {
	return b.config.Log.Loggers.SocksV4a.Formats.BindSuccessful
}

func (b BaseSocksV4aLoggerConfig) GetBoundFormat() string {
	return b.config.Log.Loggers.SocksV4a.Formats.Bound
}

func (b BaseSocksV4aLoggerConfig) IsConsoleOutputEnabled() bool {
	return b.config.Log.Loggers.SocksV4a.Outputs.Console != nil
}

func (b BaseSocksV4aLoggerConfig) IsFileOutputEnabled() bool {
	return b.config.Log.Loggers.SocksV4a.Outputs.File != nil
}

func (b BaseSocksV4aLoggerConfig) GetFilePathFormat() string {
	return b.config.Log.Loggers.SocksV4a.Outputs.File.Path
}
