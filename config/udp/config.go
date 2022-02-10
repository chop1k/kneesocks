package udp

import (
	"errors"
	"github.com/Jeffail/gabs"
)

type BindConfig interface {
	GetAddress() (string, error)
	GetPort() (uint16, error)
}

type BaseBindConfig struct {
	config gabs.Container
}

func NewBaseBindConfig(config gabs.Container) (BaseBindConfig, error) {
	return BaseBindConfig{
		config: config,
	}, nil
}

func (b BaseBindConfig) GetAddress() (string, error) {
	address, ok := b.config.Path("Udp.Bind.Address").Data().(string)

	if !ok {
		return "", errors.New("Udp.Bind.Address: Not specified or have invalid type. ")
	}

	return address, nil
}

func (b BaseBindConfig) GetPort() (uint16, error) {
	port, ok := b.config.Path("Udp.Bind.Port").Data().(float64)

	if !ok {
		return 0, errors.New("Udp.Bind.Port: Not specified or have invalid type. ")
	}

	return uint16(port), nil
}