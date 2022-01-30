package tcp

import (
	"errors"
	"socks/config/tree"
)

var (
	LoggerDisabledError        = errors.New("Tcp logger is disabled. ")
	ConsoleOutputDisabledError = errors.New("Tcp console output is disabled. ")
	FileOutputDisabledError    = errors.New("Tcp console output is disabled. ")
)

type LoggerConfig interface {
	GetLevel() (int, error)
	GetConsoleOutput() (tree.ConsoleOutputConfig, error)
	GetFileOutput() (tree.FileOutputConfig, error)
}

type BaseLoggerConfig struct {
	config tree.Config
}

func NewBaseLoggerConfig(config tree.Config) (BaseLoggerConfig, error) {
	return BaseLoggerConfig{
		config: config,
	}, nil
}

func (b BaseLoggerConfig) GetLevel() (int, error) {
	if b.config.Log.Tcp == nil {
		return 0, LoggerDisabledError
	}

	return b.config.Log.Tcp.Level, nil
}

func (b BaseLoggerConfig) GetConsoleOutput() (tree.ConsoleOutputConfig, error) {
	if b.config.Log.SocksV4 == nil {
		return tree.ConsoleOutputConfig{}, LoggerDisabledError
	}

	if b.config.Log.Tcp.Console == nil {
		return tree.ConsoleOutputConfig{}, ConsoleOutputDisabledError
	}

	return *b.config.Log.Tcp.Console, nil
}

func (b BaseLoggerConfig) GetFileOutput() (tree.FileOutputConfig, error) {
	if b.config.Log.SocksV4 == nil {
		return tree.FileOutputConfig{}, LoggerDisabledError
	}

	if b.config.Log.Tcp.File == nil {
		return tree.FileOutputConfig{}, FileOutputDisabledError
	}

	return *b.config.Log.Tcp.File, nil
}
