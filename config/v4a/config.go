package v4a

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
	allowed, ok := b.config.Path("SocksV4a.AllowConnect").Data().(bool)

	if !ok {
		return false, errors.New("SocksV4a.AllowConnect: Not specified or have invalid type. ")
	}

	return allowed, nil
}

func (b BaseConfig) IsBindAllowed() (bool, error) {
	allowed, ok := b.config.Path("SocksV4a.AllowBind").Data().(bool)

	if !ok {
		return false, errors.New("SocksV4a.AllowBind: Not specified or have invalid type. ")
	}

	return allowed, nil
}
