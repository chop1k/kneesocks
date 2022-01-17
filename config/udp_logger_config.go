package config

import (
	"errors"
	"socks/config/tree"
)

var (
	UdpLoggerDisabledError        = errors.New("Udp logger is disabled. ")
	UdpConsoleOutputDisabledError = errors.New("Udp console output is disabled. ")
	UdpFileOutputDisabledError    = errors.New("Udp console output is disabled. ")
)

type UdpLoggerConfig interface {
	GetLevel() (int, error)
	GetConsoleOutput() (tree.ConsoleOutputConfig, error)
	GetFileOutput() (tree.FileOutputConfig, error)
}

type BaseUdpLoggerConfig struct {
	config tree.Config
}

func NewBaseUdpLoggerConfig(config tree.Config) (BaseUdpLoggerConfig, error) {
	return BaseUdpLoggerConfig{
		config: config,
	}, nil
}

func (b BaseUdpLoggerConfig) GetLevel() (int, error) {
	if b.config.Log.Udp == nil {
		return 0, UdpLoggerDisabledError
	}

	return b.config.Log.Udp.Level, nil
}

func (b BaseUdpLoggerConfig) GetConsoleOutput() (tree.ConsoleOutputConfig, error) {
	if b.config.Log.SocksV4 == nil {
		return tree.ConsoleOutputConfig{}, UdpLoggerDisabledError
	}

	if b.config.Log.Udp.Console == nil {
		return tree.ConsoleOutputConfig{}, UdpConsoleOutputDisabledError
	}

	return *b.config.Log.Udp.Console, nil
}

func (b BaseUdpLoggerConfig) GetFileOutput() (tree.FileOutputConfig, error) {
	if b.config.Log.SocksV4 == nil {
		return tree.FileOutputConfig{}, UdpLoggerDisabledError
	}

	if b.config.Log.Udp.File == nil {
		return tree.FileOutputConfig{}, UdpFileOutputDisabledError
	}

	return *b.config.Log.Udp.File, nil
}
