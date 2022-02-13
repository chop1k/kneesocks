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
	if b.config.Udp == nil {
		return 0, LoggerDisabledError
	}

	return b.config.Udp.Level, nil
}

func (b BaseLoggerConfig) GetConsoleOutput() (tree.ConsoleOutputConfig, error) {
	if b.config.Udp == nil {
		return tree.ConsoleOutputConfig{}, LoggerDisabledError
	}

	if b.config.Udp.Console == nil {
		return tree.ConsoleOutputConfig{}, ConsoleOutputDisabledError
	}

	return *b.config.Udp.Console, nil
}

func (b BaseLoggerConfig) GetFileOutput() (tree.FileOutputConfig, error) {
	if b.config.Udp == nil {
		return tree.FileOutputConfig{}, LoggerDisabledError
	}

	if b.config.Udp.File == nil {
		return tree.FileOutputConfig{}, FileOutputDisabledError
	}

	return *b.config.Udp.File, nil
}
