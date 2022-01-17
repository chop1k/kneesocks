package config

import (
	"errors"
	"socks/config/tree"
)

var (
	SocksV4aLoggerDisabledError        = errors.New("SocksV4a logger is disabled. ")
	SocksV4aConsoleOutputDisabledError = errors.New("SocksV4a console output is disabled. ")
	SocksV4aFileOutputDisabledError    = errors.New("SocksV4a console output is disabled. ")
)

type SocksV4aLoggerConfig interface {
	GetLevel() (int, error)
	GetConsoleOutput() (tree.ConsoleOutputConfig, error)
	GetFileOutput() (tree.FileOutputConfig, error)
}

type BaseSocksV4aLoggerConfig struct {
	config tree.Config
}

func NewBaseSocksV4aLoggerConfig(config tree.Config) (BaseSocksV4aLoggerConfig, error) {
	return BaseSocksV4aLoggerConfig{
		config: config,
	}, nil
}

func (b BaseSocksV4aLoggerConfig) GetLevel() (int, error) {
	if b.config.Log.SocksV4a == nil {
		return 0, SocksV4aLoggerDisabledError
	}

	return b.config.Log.SocksV4a.Level, nil
}

func (b BaseSocksV4aLoggerConfig) GetConsoleOutput() (tree.ConsoleOutputConfig, error) {
	if b.config.Log.SocksV4 == nil {
		return tree.ConsoleOutputConfig{}, SocksV4aLoggerDisabledError
	}

	if b.config.Log.SocksV4a.Console == nil {
		return tree.ConsoleOutputConfig{}, SocksV4aConsoleOutputDisabledError
	}

	return *b.config.Log.SocksV4a.Console, nil
}

func (b BaseSocksV4aLoggerConfig) GetFileOutput() (tree.FileOutputConfig, error) {
	if b.config.Log.SocksV4 == nil {
		return tree.FileOutputConfig{}, SocksV4aLoggerDisabledError
	}

	if b.config.Log.SocksV4a.File == nil {
		return tree.FileOutputConfig{}, SocksV4aFileOutputDisabledError
	}

	return *b.config.Log.SocksV4a.File, nil
}
