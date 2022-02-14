package udp

import (
	"errors"
	"socks/config/tree"
)

var (
	LoggerDisabledError        = errors.New("Udp logger is disabled. ")
	ConsoleOutputDisabledError = errors.New("Udp console output is disabled. ")
	FileOutputDisabledError    = errors.New("Udp file output is disabled. ")
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
	if b.config.Udp == nil {
		return 0, LoggerDisabledError
	}

	return b.config.Udp.Level, nil
}

func (b LoggerConfig) GetConsoleOutput() (tree.ConsoleOutputConfig, error) {
	if b.config.Udp == nil {
		return tree.ConsoleOutputConfig{}, LoggerDisabledError
	}

	if b.config.Udp.Console == nil {
		return tree.ConsoleOutputConfig{}, ConsoleOutputDisabledError
	}

	return *b.config.Udp.Console, nil
}

func (b LoggerConfig) GetFileOutput() (tree.FileOutputConfig, error) {
	if b.config.Udp == nil {
		return tree.FileOutputConfig{}, LoggerDisabledError
	}

	if b.config.Udp.File == nil {
		return tree.FileOutputConfig{}, FileOutputDisabledError
	}

	return *b.config.Udp.File, nil
}
