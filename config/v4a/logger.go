package v4a

import (
	"errors"
	"github.com/Jeffail/gabs"
	"github.com/mitchellh/mapstructure"
	"socks/config/tree"
)

var (
	LoggerDisabledError        = errors.New("SocksV4a logger is disabled. ")
	ConsoleOutputDisabledError = errors.New("SocksV4a console output is disabled. ")
	FileOutputDisabledError    = errors.New("SocksV4a file output is disabled. ")
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
	if !b.config.ExistsP("Log.SocksV4a") {
		return 0, LoggerDisabledError
	}

	level, ok := b.config.Path("Log.SocksV4a.Level").Data().(float64)

	if !ok {
		return 0, errors.New("Log.SocksV4a.Level: Not specified or have invalid type. ")
	}

	return int(level), nil
}

func (b BaseLoggerConfig) GetConsoleOutput() (tree.ConsoleOutputConfig, error) {
	if !b.config.ExistsP("Log.SocksV4a") {
		return tree.ConsoleOutputConfig{}, LoggerDisabledError
	}

	if !b.config.ExistsP("Log.SocksV4a.Console") {
		return tree.ConsoleOutputConfig{}, ConsoleOutputDisabledError
	}

	output, ok := b.config.Path("Log.SocksV4a.Console").Data().(map[string]interface{})

	if !ok {
		return tree.ConsoleOutputConfig{}, errors.New("Log.SocksV4a.Console: Not specified or have invalid type. ")
	}

	_output := tree.ConsoleOutputConfig{}

	return _output, mapstructure.Decode(output, &_output)
}

func (b BaseLoggerConfig) GetFileOutput() (tree.FileOutputConfig, error) {
	if !b.config.ExistsP("Log.SocksV4a") {
		return tree.FileOutputConfig{}, LoggerDisabledError
	}

	if !b.config.ExistsP("Log.SocksV4a.File") {
		return tree.FileOutputConfig{}, FileOutputDisabledError
	}

	output, ok := b.config.Path("Log.SocksV4a.File").Data().(map[string]interface{})

	if !ok {
		return tree.FileOutputConfig{}, errors.New("Log.SocksV4a.File: Not specified or have invalid type. ")
	}

	_output := tree.FileOutputConfig{}

	return _output, mapstructure.Decode(output, &_output)
}
