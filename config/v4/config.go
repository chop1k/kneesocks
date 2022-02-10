package v4

import (
	"errors"
	"github.com/Jeffail/gabs"
)

type Config interface {
	IsConnectAllowed() (bool, error)
	IsBindAllowed() (bool, error)
}

type BaseConfig struct {
	config gabs.Container
}

func NewBaseConfig(config gabs.Container) (BaseConfig, error) {
	return BaseConfig{config: config}, nil
}

func (b BaseConfig) IsConnectAllowed() (bool, error) {
	allowed, ok := b.config.Path("SocksV4.AllowConnect").Data().(bool)

	if !ok {
		return false, errors.New("SocksV4.AllowConnect: Not specified or have invalid type. ")
	}

	return allowed, nil
}

func (b BaseConfig) IsBindAllowed() (bool, error) {
	allowed, ok := b.config.Path("SocksV4.AllowBind").Data().(bool)

	if !ok {
		return false, errors.New("SocksV4.AllowBind: Not specified or have invalid type. ")
	}

	return allowed, nil
}
