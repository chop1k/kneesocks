package v4

import (
	"errors"
	"socks/config/tree"
)

var (
	LoggerDisabledError        = errors.New("SocksV4 logger is disabled. ")
	ConsoleOutputDisabledError = errors.New("SocksV4 console output is disabled. ")
	FileOutputDisabledError    = errors.New("SocksV4 file output is disabled. ")
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
	if b.config.SocksV4 == nil {
		return 0, LoggerDisabledError
	}

	return b.config.SocksV4.Level, nil
}

func (b LoggerConfig) GetConsoleOutput() (tree.ConsoleOutputConfig, error) {
	if b.config.SocksV4 == nil {
		return tree.ConsoleOutputConfig{}, LoggerDisabledError
	}

	if b.config.SocksV4.Console == nil {
		return tree.ConsoleOutputConfig{}, ConsoleOutputDisabledError
	}

	return *b.config.SocksV4.Console, nil
}

func (b LoggerConfig) GetFileOutput() (tree.FileOutputConfig, error) {
	if b.config.SocksV4 == nil {
		return tree.FileOutputConfig{}, LoggerDisabledError
	}

	if b.config.SocksV4.File == nil {
		return tree.FileOutputConfig{}, FileOutputDisabledError
	}

	return *b.config.SocksV4.File, nil
}
