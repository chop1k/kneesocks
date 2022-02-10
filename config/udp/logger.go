package udp

import (
	"errors"
	"github.com/Jeffail/gabs"
	"github.com/mitchellh/mapstructure"
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
	config gabs.Container
}

func NewBaseLoggerConfig(config gabs.Container) (BaseLoggerConfig, error) {
	return BaseLoggerConfig{
		config: config,
	}, nil
}

func (b BaseLoggerConfig) GetLevel() (int, error) {
	if !b.config.ExistsP("Log.Udp") {
		return 0, LoggerDisabledError
	}

	level, ok := b.config.Path("Log.Udp.Level").Data().(float64)

	if !ok {
		return 0, errors.New("Log.Udp.Level: Not specified or have invalid type. ")
	}

	return int(level), nil
}

func (b BaseLoggerConfig) GetConsoleOutput() (tree.ConsoleOutputConfig, error) {
	if !b.config.ExistsP("Log.Udp") {
		return tree.ConsoleOutputConfig{}, LoggerDisabledError
	}

	if !b.config.ExistsP("Log.Udp.Console") {
		return tree.ConsoleOutputConfig{}, ConsoleOutputDisabledError
	}

	output, ok := b.config.Path("Log.Udp.Console").Data().(map[string]interface{})

	if !ok {
		return tree.ConsoleOutputConfig{}, errors.New("Log.Udp.Console: Not specified or have invalid type. ")
	}

	_output := tree.ConsoleOutputConfig{}

	return _output, mapstructure.Decode(output, &_output)
}

func (b BaseLoggerConfig) GetFileOutput() (tree.FileOutputConfig, error) {
	if !b.config.ExistsP("Log.Udp") {
		return tree.FileOutputConfig{}, LoggerDisabledError
	}

	if !b.config.ExistsP("Log.Udp.File") {
		return tree.FileOutputConfig{}, FileOutputDisabledError
	}

	output, ok := b.config.Path("Log.Udp.File").Data().(map[string]interface{})

	if !ok {
		return tree.FileOutputConfig{}, errors.New("Log.Udp.File: Not specified or have invalid type. ")
	}

	_output := tree.FileOutputConfig{}

	return _output, mapstructure.Decode(output, &_output)
}
