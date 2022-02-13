package v4a

import (
	"errors"
	"socks/config/tree"
)

var (
	LoggerDisabledError        = errors.New("SocksV4a logger is disabled. ")
	ConsoleOutputDisabledError = errors.New("SocksV4a console output is disabled. ")
	FileOutputDisabledError    = errors.New("SocksV4a file output is disabled. ")
)

type LoggerConfig interface {
	GetLevel() (int, error)
	GetConsoleOutput() (tree.ConsoleOutputConfig, error)
	GetFileOutput() (tree.FileOutputConfig, error)
}

type BaseLoggerConfig struct {
	config tree.LogConfig
}

func NewBaseLoggerConfig(config tree.LogConfig) (BaseLoggerConfig, error) {
	return BaseLoggerConfig{
		config: config,
	}, nil
}

func (b BaseLoggerConfig) GetLevel() (int, error) {
	if b.config.SocksV4a == nil {
		return 0, LoggerDisabledError
	}

	return b.config.SocksV4a.Level, nil
}

func (b BaseLoggerConfig) GetConsoleOutput() (tree.ConsoleOutputConfig, error) {
	if b.config.SocksV4a == nil {
		return tree.ConsoleOutputConfig{}, LoggerDisabledError
	}

	if b.config.SocksV4a.Console == nil {
		return tree.ConsoleOutputConfig{}, ConsoleOutputDisabledError
	}

	return *b.config.SocksV4a.Console, nil
}

func (b BaseLoggerConfig) GetFileOutput() (tree.FileOutputConfig, error) {
	if b.config.SocksV4a == nil {
		return tree.FileOutputConfig{}, LoggerDisabledError
	}

	if b.config.SocksV4a.File == nil {
		return tree.FileOutputConfig{}, FileOutputDisabledError
	}

	return *b.config.SocksV4a.File, nil
}
