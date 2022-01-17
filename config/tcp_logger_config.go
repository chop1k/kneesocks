package config

import (
	"errors"
	"socks/config/tree"
)

var (
	TcpLoggerDisabledError        = errors.New("Tcp logger is disabled. ")
	TcpConsoleOutputDisabledError = errors.New("Tcp console output is disabled. ")
	TcpFileOutputDisabledError    = errors.New("Tcp console output is disabled. ")
)

type TcpLoggerConfig interface {
	GetLevel() (int, error)
	GetConsoleOutput() (tree.ConsoleOutputConfig, error)
	GetFileOutput() (tree.FileOutputConfig, error)
}

type BaseTcpLoggerConfig struct {
	config tree.Config
}

func NewBaseTcpLoggerConfig(config tree.Config) (BaseTcpLoggerConfig, error) {
	return BaseTcpLoggerConfig{
		config: config,
	}, nil
}

func (b BaseTcpLoggerConfig) GetLevel() (int, error) {
	if b.config.Log.Tcp == nil {
		return 0, TcpLoggerDisabledError
	}

	return b.config.Log.Tcp.Level, nil
}

func (b BaseTcpLoggerConfig) GetConsoleOutput() (tree.ConsoleOutputConfig, error) {
	if b.config.Log.SocksV4 == nil {
		return tree.ConsoleOutputConfig{}, TcpLoggerDisabledError
	}

	if b.config.Log.Tcp.Console == nil {
		return tree.ConsoleOutputConfig{}, TcpConsoleOutputDisabledError
	}

	return *b.config.Log.Tcp.Console, nil
}

func (b BaseTcpLoggerConfig) GetFileOutput() (tree.FileOutputConfig, error) {
	if b.config.Log.SocksV4 == nil {
		return tree.FileOutputConfig{}, TcpLoggerDisabledError
	}

	if b.config.Log.Tcp.File == nil {
		return tree.FileOutputConfig{}, TcpFileOutputDisabledError
	}

	return *b.config.Log.Tcp.File, nil
}
