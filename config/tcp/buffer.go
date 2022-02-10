package tcp

import (
	"errors"
	"github.com/Jeffail/gabs"
)

type BufferConfig interface {
	GetClientSize() (uint, error)
	GetHostSize() (uint, error)
}

type BaseBufferConfig struct {
	config gabs.Container
}

func NewBaseBufferConfig(config gabs.Container) (BaseBufferConfig, error) {
	return BaseBufferConfig{config: config}, nil
}

func (b BaseBufferConfig) GetClientSize() (uint, error) {
	buffer, ok := b.config.Path("Tcp.Buffer.ClientSize").Data().(float64)

	if !ok {
		return 0, errors.New("Tcp.Buffer.ClientSize: Not specified or have invalid type. ")
	}

	return uint(buffer), nil
}

func (b BaseBufferConfig) GetHostSize() (uint, error) {
	buffer, ok := b.config.Path("Tcp.Buffer.HostSize").Data().(float64)

	if !ok {
		return 0, errors.New("Tcp.Buffer.HostSize: Not specified or have invalid type. ")
	}

	return uint(buffer), nil
}
