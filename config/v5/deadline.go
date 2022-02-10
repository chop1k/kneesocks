package v5

import (
	"errors"
	"github.com/Jeffail/gabs"
	"time"
)

type DeadlineConfig interface {
	GetRequestDeadline() (time.Duration, error)
	GetConnectDeadline() (time.Duration, error)
	GetBindDeadline() (time.Duration, error)
	GetPasswordDeadline() (time.Duration, error)
	GetResponseDeadline() (time.Duration, error)
	GetSelectionDeadline() (time.Duration, error)
	GetPasswordResponseDeadline() (time.Duration, error)
}

type BaseDeadlineConfig struct {
	config gabs.Container
}

func NewBaseDeadlineConfig(config gabs.Container) (BaseDeadlineConfig, error) {
	return BaseDeadlineConfig{config: config}, nil
}

func (b BaseDeadlineConfig) GetRequestDeadline() (time.Duration, error) {
	deadline, ok := b.config.Path("SocksV5.Deadline.Request").Data().(float64)

	if !ok {
		return 0, errors.New("SocksV5.Deadline.Request: Not specified or have invalid type. ")
	}

	return time.Second * time.Duration(deadline), nil
}

func (b BaseDeadlineConfig) GetConnectDeadline() (time.Duration, error) {
	deadline, ok := b.config.Path("SocksV5.Deadline.Connect").Data().(float64)

	if !ok {
		return 0, errors.New("SocksV5.Deadline.Connect: Not specified or have invalid type. ")
	}

	return time.Second * time.Duration(deadline), nil
}

func (b BaseDeadlineConfig) GetBindDeadline() (time.Duration, error) {
	deadline, ok := b.config.Path("SocksV5.Deadline.Bind").Data().(float64)

	if !ok {
		return 0, errors.New("SocksV5.Deadline.Bind: Not specified or have invalid type. ")
	}

	return time.Second * time.Duration(deadline), nil
}

func (b BaseDeadlineConfig) GetPasswordDeadline() (time.Duration, error) {
	deadline, ok := b.config.Path("SocksV5.Deadline.Password").Data().(float64)

	if !ok {
		return 0, errors.New("SocksV5.Deadline.Password: Not specified or have invalid type. ")
	}

	return time.Second * time.Duration(deadline), nil
}

func (b BaseDeadlineConfig) GetResponseDeadline() (time.Duration, error) {
	deadline, ok := b.config.Path("SocksV5.Deadline.Response").Data().(float64)

	if !ok {
		return 0, errors.New("SocksV5.Deadline.Response: Not specified or have invalid type. ")
	}

	return time.Second * time.Duration(deadline), nil
}

func (b BaseDeadlineConfig) GetSelectionDeadline() (time.Duration, error) {
	deadline, ok := b.config.Path("SocksV5.Deadline.Selection").Data().(float64)

	if !ok {
		return 0, errors.New("SocksV5.Deadline.Selection: Not specified or have invalid type. ")
	}

	return time.Second * time.Duration(deadline), nil
}

func (b BaseDeadlineConfig) GetPasswordResponseDeadline() (time.Duration, error) {
	deadline, ok := b.config.Path("SocksV5.Deadline.PasswordResponse").Data().(float64)

	if !ok {
		return 0, errors.New("SocksV5.Deadline.PasswordResponse: Not specified or have invalid type. ")
	}

	return time.Second * time.Duration(deadline), nil
}
