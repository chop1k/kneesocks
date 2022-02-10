package udp

import (
	"errors"
	"github.com/Jeffail/gabs"
)

type BufferConfig interface {
	GetPacketSize() (uint, error)
}

type BaseBufferConfig struct {
	config gabs.Container
}

func NewBaseBufferConfig(config gabs.Container) (BaseBufferConfig, error) {
	return BaseBufferConfig{config: config}, nil
}

func (b BaseBufferConfig) GetPacketSize() (uint, error) {
	buffer, ok := b.config.Path("Udp.Buffer.PacketSize").Data().(float64)

	if !ok {
		return 0, errors.New("Udp.Buffer.PacketSize: Not specified or have invalid type. ")
	}

	return uint(buffer), nil
}
