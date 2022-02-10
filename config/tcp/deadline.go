package tcp

import (
	"errors"
	"github.com/Jeffail/gabs"
	"time"
)

type DeadlineConfig interface {
	GetExchangeDeadline() (time.Duration, error)
	GetWelcomeDeadline() (time.Duration, error)
}

type BaseDeadlineConfig struct {
	config gabs.Container
}

func NewBaseDeadlineConfig(config gabs.Container) (BaseDeadlineConfig, error) {
	return BaseDeadlineConfig{config: config}, nil
}

func (b BaseDeadlineConfig) GetExchangeDeadline() (time.Duration, error) {
	deadline, ok := b.config.Path("Tcp.Deadline.Exchange").Data().(float64)

	if !ok {
		return 0, errors.New("Tcp.Deadline.Exchange: Not specified or have invalid type. ")
	}

	return time.Second * time.Duration(deadline), nil
}

func (b BaseDeadlineConfig) GetWelcomeDeadline() (time.Duration, error) {
	deadline, ok := b.config.Path("Tcp.Deadline.Welcome").Data().(float64)

	if !ok {
		return 0, errors.New("Tcp.Deadline.Welcome: Not specified or have invalid type. ")
	}

	return time.Second * time.Duration(deadline), nil
}
