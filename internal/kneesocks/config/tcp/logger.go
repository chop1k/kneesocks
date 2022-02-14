package tcp

import (
	"errors"
	"socks/internal/kneesocks/config/tree"
)

var (
	LoggerDisabledError        = errors.New("Tcp logger is disabled. ")
	ConsoleOutputDisabledError = errors.New("Tcp console output is disabled. ")
	FileOutputDisabledError    = errors.New("Tcp file output is disabled. ")
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
	if b.config.Tcp == nil {
		return 0, LoggerDisabledError
	}

	return b.config.Tcp.Level, nil
}

func (b LoggerConfig) GetConsoleOutput() (tree.ConsoleOutputConfig, error) {
	if b.config.Tcp == nil {
		return tree.ConsoleOutputConfig{}, LoggerDisabledError
	}

	if b.config.Tcp.Console == nil {
		return tree.ConsoleOutputConfig{}, ConsoleOutputDisabledError
	}

	return *b.config.Tcp.Console, nil
}

func (b LoggerConfig) GetFileOutput() (tree.FileOutputConfig, error) {
	if b.config.Tcp == nil {
		return tree.FileOutputConfig{}, LoggerDisabledError
	}

	if b.config.Tcp.File == nil {
		return tree.FileOutputConfig{}, FileOutputDisabledError
	}

	return *b.config.Tcp.File, nil
}
