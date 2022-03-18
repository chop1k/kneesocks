package v4a

import (
	"errors"
	"socks/internal/kneesocks/config/tree"
)

var (
	LoggerDisabledError        = errors.New("SocksV4a logger is disabled. ")
	ConsoleOutputDisabledError = errors.New("SocksV4a console output is disabled. ")
	FileOutputDisabledError    = errors.New("SocksV4a file output is disabled. ")
)

type LoggerConfig struct {
	config *tree.SocksV4aLoggerConfig
}

func NewLoggerConfig(config *tree.SocksV4aLoggerConfig) (LoggerConfig, error) {
	return LoggerConfig{
		config: config,
	}, nil
}

func (b LoggerConfig) GetLevel() (int, error) {
	if b.config == nil {
		return 0, LoggerDisabledError
	}

	return b.config.Level, nil
}

func (b LoggerConfig) GetConsoleOutput() (tree.ConsoleOutputConfig, error) {
	if b.config == nil {
		return tree.ConsoleOutputConfig{}, LoggerDisabledError
	}

	if b.config.Console == nil {
		return tree.ConsoleOutputConfig{}, ConsoleOutputDisabledError
	}

	return *b.config.Console, nil
}

func (b LoggerConfig) GetFileOutput() (tree.FileOutputConfig, error) {
	if b.config == nil {
		return tree.FileOutputConfig{}, LoggerDisabledError
	}

	if b.config.File == nil {
		return tree.FileOutputConfig{}, FileOutputDisabledError
	}

	return *b.config.File, nil
}
