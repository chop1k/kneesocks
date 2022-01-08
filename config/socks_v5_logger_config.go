package config

import "socks/config/tree"

type SocksV5LoggerConfig interface {
	SocksV4aLoggerConfig
	GetAuthenticationSuccessfulFormat() string
	GetAuthenticationFailedFormat() string
	GetUdpAssociationRequestFormat() string
	GetUdpAssociationSuccessfulFormat() string
	GetUdpAssociationFailedFormat() string
}

type BaseSocksV5LoggerConfig struct {
	config tree.Config
}

func NewBaseSocksV5LoggerConfig(config tree.Config) BaseSocksV5LoggerConfig {
	return BaseSocksV5LoggerConfig{config: config}
}

func (b BaseSocksV5LoggerConfig) GetConnectRequestFormat() string {
	return b.config.Log.Loggers.SocksV5.Formats.ConnectRequest
}

func (b BaseSocksV5LoggerConfig) GetConnectFailedFormat() string {
	return b.config.Log.Loggers.SocksV5.Formats.ConnectFailed
}

func (b BaseSocksV5LoggerConfig) GetConnectSuccessfulFormat() string {
	return b.config.Log.Loggers.SocksV5.Formats.ConnectSuccessful
}

func (b BaseSocksV5LoggerConfig) GetBindRequestFormat() string {
	return b.config.Log.Loggers.SocksV5.Formats.BindRequest
}

func (b BaseSocksV5LoggerConfig) GetBindFailedFormat() string {
	return b.config.Log.Loggers.SocksV5.Formats.BindFailed
}

func (b BaseSocksV5LoggerConfig) GetBindSuccessfulFormat() string {
	return b.config.Log.Loggers.SocksV5.Formats.BindSuccessful
}

func (b BaseSocksV5LoggerConfig) GetBoundFormat() string {
	return b.config.Log.Loggers.SocksV5.Formats.Bound
}

func (b BaseSocksV5LoggerConfig) IsConsoleOutputEnabled() bool {
	return b.config.Log.Loggers.SocksV5.Outputs.Console != nil
}

func (b BaseSocksV5LoggerConfig) IsFileOutputEnabled() bool {
	return b.config.Log.Loggers.SocksV5.Outputs.File != nil
}

func (b BaseSocksV5LoggerConfig) GetFilePathFormat() string {
	return b.config.Log.Loggers.SocksV5.Outputs.File.Path
}

func (b BaseSocksV5LoggerConfig) GetAuthenticationSuccessfulFormat() string {
	return b.config.Log.Loggers.SocksV5.Formats.AuthenticationSuccessful
}

func (b BaseSocksV5LoggerConfig) GetAuthenticationFailedFormat() string {
	return b.config.Log.Loggers.SocksV5.Formats.AuthenticationFailed
}

func (b BaseSocksV5LoggerConfig) GetUdpAssociationRequestFormat() string {
	return b.config.Log.Loggers.SocksV5.Formats.UdpAssociationRequest
}

func (b BaseSocksV5LoggerConfig) GetUdpAssociationSuccessfulFormat() string {
	return b.config.Log.Loggers.SocksV5.Formats.UdpAssociationSuccessful
}

func (b BaseSocksV5LoggerConfig) GetUdpAssociationFailedFormat() string {
	return b.config.Log.Loggers.SocksV5.Formats.UdpAssociationFailed
}
