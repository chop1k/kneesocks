package helpers

import (
	"socks/config/v4a"
	"socks/managers"
)

type Limiter interface {
	IsLimited() (bool, error)
}

type BaseLimiter struct {
	config  v4a.RestrictionsConfig
	manager *managers.ConnectionsManager
}

func NewBaseLimiter(
	config v4a.RestrictionsConfig,
	manager *managers.ConnectionsManager,
) (BaseLimiter, error) {
	return BaseLimiter{
		config:  config,
		manager: manager,
	}, nil
}

func (b BaseLimiter) IsLimited() (bool, error) {
	limit, err := b.config.GetRate()

	if err != nil {
		return false, err
	}

	if limit.MaxSimultaneousConnections <= 0 {
		return false, nil
	}

	b.manager.Increment("v4a.anonymous")

	limited := b.manager.IsExceed("v4a.anonymous", limit.MaxSimultaneousConnections)

	if limited {
		b.manager.Decrement("v4a.anonymous")

		return true, nil
	} else {
		return false, nil
	}
}
