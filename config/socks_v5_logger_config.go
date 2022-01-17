package config

import (
	"errors"
	"socks/config/tree"
)

var (
	SocksV5LoggerDisabledError        = errors.New("SocksV5 logger is disabled. ")
	SocksV5ConsoleOutputDisabledError = errors.New("SocksV5 console output is disabled. ")
	SocksV5FileOutputDisabledError    = errors.New("SocksV5 console output is disabled. ")
)

type SocksV5LoggerConfig interface {
	GetLevel() (int, error)
	GetConsoleOutput() (tree.ConsoleOutputConfig, error)
	GetFileOutput() (tree.FileOutputConfig, error)
}

type BaseSocksV5LoggerConfig struct {
	config tree.Config
}

func NewBaseSocksV5LoggerConfig(config tree.Config) (BaseSocksV5LoggerConfig, error) {
	return BaseSocksV5LoggerConfig{
		config: config,
	}, nil
}

func (b BaseSocksV5LoggerConfig) GetLevel() (int, error) {
	if b.config.Log.SocksV5 == nil {
		return 0, SocksV5LoggerDisabledError
	}

	return b.config.Log.SocksV5.Level, nil
}

func (b BaseSocksV5LoggerConfig) GetConsoleOutput() (tree.ConsoleOutputConfig, error) {
	if b.config.Log.SocksV4 == nil {
		return tree.ConsoleOutputConfig{}, SocksV5LoggerDisabledError
	}

	if b.config.Log.SocksV5.Console == nil {
		return tree.ConsoleOutputConfig{}, SocksV5ConsoleOutputDisabledError
	}

	return *b.config.Log.SocksV5.Console, nil
}

func (b BaseSocksV5LoggerConfig) GetFileOutput() (tree.FileOutputConfig, error) {
	if b.config.Log.SocksV4 == nil {
		return tree.FileOutputConfig{}, SocksV5LoggerDisabledError
	}

	if b.config.Log.SocksV5.File == nil {
		return tree.FileOutputConfig{}, SocksV5FileOutputDisabledError
	}

	return *b.config.Log.SocksV5.File, nil
}
