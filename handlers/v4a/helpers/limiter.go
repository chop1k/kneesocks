package helpers

import (
	"socks/config/v4a"
	"socks/managers"
)

type Limiter interface {
	IsLimited() bool
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

func (b BaseLimiter) IsLimited() bool {
	limit, err := b.config.GetRate()

	if err != nil {
		panic(err)
	}

	if limit.MaxSimultaneousConnections <= 0 {
		return false
	}

	b.manager.Increment("v4a.anonymous")

	limited := b.manager.IsExceed("v4a.anonymous", limit.MaxSimultaneousConnections)

	if limited {
		b.manager.Decrement("v4a.anonymous")

		return true
	} else {
		return false
	}
}
