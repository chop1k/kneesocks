package v5

import (
	"errors"
	"socks/internal/kneesocks/config/tree"
)

var (
	LoggerDisabledError        = errors.New("SocksV5 logger is disabled. ")
	ConsoleOutputDisabledError = errors.New("SocksV5 console output is disabled. ")
	FileOutputDisabledError    = errors.New("SocksV5 file output is disabled. ")
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
	if b.config.SocksV5 == nil {
		return 0, LoggerDisabledError
	}

	return b.config.SocksV5.Level, nil
}

func (b LoggerConfig) GetConsoleOutput() (tree.ConsoleOutputConfig, error) {
	if b.config.SocksV5 == nil {
		return tree.ConsoleOutputConfig{}, LoggerDisabledError
	}

	if b.config.SocksV5.Console == nil {
		return tree.ConsoleOutputConfig{}, ConsoleOutputDisabledError
	}

	return *b.config.SocksV5.Console, nil
}

func (b LoggerConfig) GetFileOutput() (tree.FileOutputConfig, error) {
	if b.config.SocksV5 == nil {
		return tree.FileOutputConfig{}, LoggerDisabledError
	}

	if b.config.SocksV5.File == nil {
		return tree.FileOutputConfig{}, FileOutputDisabledError
	}

	return *b.config.SocksV5.File, nil
}
