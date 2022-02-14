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

type LoggerConfig struct {
	config tree.LogConfig
}

func NewLoggerConfig(config tree.LogConfig) (LoggerConfig, error) {
	return LoggerConfig{
		config: config,
	}, nil
}

func (b LoggerConfig) GetLevel() (int, error) {
	if b.config.SocksV4a == nil {
		return 0, LoggerDisabledError
	}

	return b.config.SocksV4a.Level, nil
}

func (b LoggerConfig) GetConsoleOutput() (tree.ConsoleOutputConfig, error) {
	if b.config.SocksV4a == nil {
		return tree.ConsoleOutputConfig{}, LoggerDisabledError
	}

	if b.config.SocksV4a.Console == nil {
		return tree.ConsoleOutputConfig{}, ConsoleOutputDisabledError
	}

	return *b.config.SocksV4a.Console, nil
}

func (b LoggerConfig) GetFileOutput() (tree.FileOutputConfig, error) {
	if b.config.SocksV4a == nil {
		return tree.FileOutputConfig{}, LoggerDisabledError
	}

	if b.config.SocksV4a.File == nil {
		return tree.FileOutputConfig{}, FileOutputDisabledError
	}

	return *b.config.SocksV4a.File, nil
}
