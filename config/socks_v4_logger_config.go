package config

import (
	"errors"
	"socks/config/tree"
)

var (
	SocksV4LoggerDisabledError        = errors.New("SocksV4 logger is disabled. ")
	SocksV4ConsoleOutputDisabledError = errors.New("SocksV4 console output is disabled. ")
	SocksV4FileOutputDisabledError    = errors.New("SocksV4 console output is disabled. ")
)

type SocksV4LoggerConfig interface {
	GetLevel() (int, error)
	GetConsoleOutput() (tree.ConsoleOutputConfig, error)
	GetFileOutput() (tree.FileOutputConfig, error)
}

type BaseSocksV4LoggerConfig struct {
	config tree.Config
}

func NewBaseSocksV4LoggerConfig(config tree.Config) (BaseSocksV4LoggerConfig, error) {
	return BaseSocksV4LoggerConfig{
		config: config,
	}, nil
}

func (b BaseSocksV4LoggerConfig) GetLevel() (int, error) {
	if b.config.Log.SocksV4 == nil {
		return 0, SocksV4LoggerDisabledError
	}

	return b.config.Log.SocksV4.Level, nil
}

func (b BaseSocksV4LoggerConfig) GetConsoleOutput() (tree.ConsoleOutputConfig, error) {
	if b.config.Log.SocksV4 == nil {
		return tree.ConsoleOutputConfig{}, SocksV4LoggerDisabledError
	}

	if b.config.Log.SocksV4.Console == nil {
		return tree.ConsoleOutputConfig{}, SocksV4ConsoleOutputDisabledError
	}

	return *b.config.Log.SocksV4.Console, nil
}

func (b BaseSocksV4LoggerConfig) GetFileOutput() (tree.FileOutputConfig, error) {
	if b.config.Log.SocksV4 == nil {
		return tree.FileOutputConfig{}, SocksV4LoggerDisabledError
	}

	if b.config.Log.SocksV4.File == nil {
		return tree.FileOutputConfig{}, SocksV4FileOutputDisabledError
	}

	return *b.config.Log.SocksV4.File, nil
}
