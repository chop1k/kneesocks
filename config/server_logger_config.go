package config

import (
	"errors"
	"socks/config/tree"
)

var (
	ServerLoggerDisabledError        = errors.New("Server logger is disabled. ")
	ServerConsoleOutputDisabledError = errors.New("Server console output is disabled. ")
	ServerFileOutputDisabledError    = errors.New("Server console output is disabled. ")
)

type ServerLoggerConfig interface {
	GetLevel() (int, error)
	GetConsoleOutput() (tree.ConsoleOutputConfig, error)
	GetFileOutput() (tree.FileOutputConfig, error)
}

type BaseServerLoggerConfig struct {
	config tree.Config
}

func NewBaseServerLoggerConfig(config tree.Config) (BaseServerLoggerConfig, error) {
	return BaseServerLoggerConfig{
		config: config,
	}, nil
}

func (b BaseServerLoggerConfig) GetLevel() (int, error) {
	if b.config.Log.Server == nil {
		return 0, ServerLoggerDisabledError
	}

	return b.config.Log.Server.Level, nil
}

func (b BaseServerLoggerConfig) GetConsoleOutput() (tree.ConsoleOutputConfig, error) {
	if b.config.Log.Server == nil {
		return tree.ConsoleOutputConfig{}, ServerLoggerDisabledError
	}

	if b.config.Log.Server.Console == nil {
		return tree.ConsoleOutputConfig{}, ServerConsoleOutputDisabledError
	}

	return *b.config.Log.Server.Console, nil
}

func (b BaseServerLoggerConfig) GetFileOutput() (tree.FileOutputConfig, error) {
	if b.config.Log.Server == nil {
		return tree.FileOutputConfig{}, ServerLoggerDisabledError
	}

	if b.config.Log.Server.File == nil {
		return tree.FileOutputConfig{}, ServerFileOutputDisabledError
	}

	return *b.config.Log.Server.File, nil
}
