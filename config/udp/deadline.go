package udp

import (
	"errors"
	"github.com/Jeffail/gabs"
	"time"
)

type DeadlineConfig interface {
	GetReadDeadline() (time.Duration, error)
}

type BaseDeadlineConfig struct {
	config gabs.Container
}

func NewBaseDeadlineConfig(config gabs.Container) (BaseDeadlineConfig, error) {
	return BaseDeadlineConfig{config: config}, nil
}

func (b BaseDeadlineConfig) GetReadDeadline() (time.Duration, error) {
	deadline, ok := b.config.Path("Udp.Deadline.Read").Data().(float64)

	if !ok {
		return 0, errors.New("Udp.Deadline.Read: Not specified or have invalid type. ")
	}

	return time.Second * time.Duration(deadline), nil
}
