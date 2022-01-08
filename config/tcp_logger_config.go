package config

import "socks/config/tree"

type TcpLoggerConfig interface {
	GetConnectionAcceptedFormat() string
	GetConnectionDeniedFormat() string
	GetConnectionProtocolDeterminedFormat() string
	GetConnectionBoundFormat() string
	GetListenFormat() string
	IsConsoleOutputEnabled() bool
	IsFileOutputEnabled() bool
	GetFilePathFormat() string
}

type BaseTcpLoggerConfig struct {
	config tree.Config
}

func NewBaseTcpLoggerConfig(config tree.Config) BaseTcpLoggerConfig {
	return BaseTcpLoggerConfig{config: config}
}

func (b BaseTcpLoggerConfig) GetConnectionAcceptedFormat() string {
	return b.config.Log.Loggers.Tcp.Formats.ConnectionAccepted
}

func (b BaseTcpLoggerConfig) GetConnectionDeniedFormat() string {
	return b.config.Log.Loggers.Tcp.Formats.ConnectionDenied
}

func (b BaseTcpLoggerConfig) GetConnectionProtocolDeterminedFormat() string {
	return b.config.Log.Loggers.Tcp.Formats.ConnectionProtocolDetermined
}

func (b BaseTcpLoggerConfig) GetConnectionBoundFormat() string {
	return b.config.Log.Loggers.Tcp.Formats.ConnectionBound
}

func (b BaseTcpLoggerConfig) GetListenFormat() string {
	return b.config.Log.Loggers.Tcp.Formats.Listen
}

func (b BaseTcpLoggerConfig) IsConsoleOutputEnabled() bool {
	return b.config.Log.Loggers.Tcp.Outputs.Console != nil
}

func (b BaseTcpLoggerConfig) IsFileOutputEnabled() bool {
	return b.config.Log.Loggers.Tcp.Outputs.File != nil
}

func (b BaseTcpLoggerConfig) GetFilePathFormat() string {
	return b.config.Log.Loggers.Tcp.Outputs.File.Path
}
