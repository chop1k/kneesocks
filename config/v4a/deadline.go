package v4a

import (
	"errors"
	"github.com/Jeffail/gabs"
	"time"
)

type DeadlineConfig interface {
	GetConnectDeadline() (time.Duration, error)
	GetBindDeadline() (time.Duration, error)
	GetResponseDeadline() (time.Duration, error)
}

type BaseDeadlineConfig struct {
	config gabs.Container
}

func NewBaseDeadlineConfig(config gabs.Container) (BaseDeadlineConfig, error) {
	return BaseDeadlineConfig{config: config}, nil
}

func (b BaseDeadlineConfig) GetConnectDeadline() (time.Duration, error) {
	deadline, ok := b.config.Path("SocksV4a.Deadline.Connect").Data().(float64)

	if !ok {
		return 0, errors.New("SocksV4a.Deadline.Connect: Not specified or have invalid type. ")
	}

	return time.Second * time.Duration(deadline), nil
}

func (b BaseDeadlineConfig) GetBindDeadline() (time.Duration, error) {
	deadline, ok := b.config.Path("SocksV4a.Deadline.Bind").Data().(float64)

	if !ok {
		return 0, errors.New("SocksV4a.Deadline.Bind: Not specified or have invalid type. ")
	}

	return time.Second * time.Duration(deadline), nil
}

func (b BaseDeadlineConfig) GetResponseDeadline() (time.Duration, error) {
	deadline, ok := b.config.Path("SocksV4a.Deadline.Response").Data().(float64)

	if !ok {
		return 0, errors.New("SocksV4a.Deadline.Response: Not specified or have invalid type. ")
	}

	return time.Second * time.Duration(deadline), nil
}
