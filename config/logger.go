package config

import (
	"errors"
	"github.com/Jeffail/gabs"
	"github.com/mitchellh/mapstructure"
	"socks/config/tree"
)

var (
	LoggerDisabledError        = errors.New("Server logger is disabled. ")
	ConsoleOutputDisabledError = errors.New("Server console output is disabled. ")
	FileOutputDisabledError    = errors.New("Server file output is disabled. ")
)

type ServerLoggerConfig interface {
	GetLevel() (int, error)
	GetConsoleOutput() (tree.ConsoleOutputConfig, error)
	GetFileOutput() (tree.FileOutputConfig, error)
}

type BaseServerLoggerConfig struct {
	config gabs.Container
}

func NewBaseServerLoggerConfig(config gabs.Container) (BaseServerLoggerConfig, error) {
	return BaseServerLoggerConfig{
		config: config,
	}, nil
}

func (b BaseServerLoggerConfig) GetLevel() (int, error) {
	if !b.config.ExistsP("Log.Server") {
		return 0, LoggerDisabledError
	}

	level, ok := b.config.Path("Log.Server.Level").Data().(float64)

	if !ok {
		return 0, errors.New("Log.Server.Level: Not specified or have invalid type. ")
	}

	return int(level), nil
}

func (b BaseServerLoggerConfig) GetConsoleOutput() (tree.ConsoleOutputConfig, error) {
	if !b.config.ExistsP("Log.Server") {
		return tree.ConsoleOutputConfig{}, LoggerDisabledError
	}

	if !b.config.ExistsP("Log.Server.Console") {
		return tree.ConsoleOutputConfig{}, ConsoleOutputDisabledError
	}

	output, ok := b.config.Path("Log.Server.Console").Data().(map[string]interface{})

	if !ok {
		return tree.ConsoleOutputConfig{}, errors.New("Log.Server.Console: Not specified or have invalid type. ")
	}

	_output := tree.ConsoleOutputConfig{}

	return _output, mapstructure.Decode(output, &_output)
}

func (b BaseServerLoggerConfig) GetFileOutput() (tree.FileOutputConfig, error) {
	if !b.config.ExistsP("Log.Server") {
		return tree.FileOutputConfig{}, LoggerDisabledError
	}

	if !b.config.ExistsP("Log.Server.File") {
		return tree.FileOutputConfig{}, FileOutputDisabledError
	}

	output, ok := b.config.Path("Log.Server.File").Data().(map[string]interface{})

	if !ok {
		return tree.FileOutputConfig{}, errors.New("Log.Server.File: Not specified or have invalid type. ")
	}

	_output := tree.FileOutputConfig{}

	return _output, mapstructure.Decode(output, &_output)
}
