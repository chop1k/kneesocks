package helpers

import (
	"socks/config/v4a"
	"socks/managers"
)

type Limiter interface {
	IsLimited(config v4a.Config) bool
}

type BaseLimiter struct {
	manager *managers.ConnectionsManager
}

func NewBaseLimiter(
	manager *managers.ConnectionsManager,
) (BaseLimiter, error) {
	return BaseLimiter{
		manager: manager,
	}, nil
}

func (b BaseLimiter) IsLimited(config v4a.Config) bool {
	rate := config.Restrictions.Rate

	if rate.MaxSimultaneousConnections <= 0 {
		return false
	}

	b.manager.Increment("v4a.anonymous")

	limited := b.manager.IsExceed("v4a.anonymous", rate.MaxSimultaneousConnections)

	if limited {
		b.manager.Decrement("v4a.anonymous")

		return true
	} else {
		return false
	}
}
