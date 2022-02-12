package helpers

import (
	v4 "socks/config/v4"
	"socks/managers"
)

type Limiter interface {
	IsLimited(config v4.Config) bool
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

func (b BaseLimiter) IsLimited(config v4.Config) bool {
	rate := config.Restrictions.Rate

	if rate.MaxSimultaneousConnections <= 0 {
		return false
	}

	b.manager.Increment("v4.anonymous")

	limited := b.manager.IsExceed("v4.anonymous", rate.MaxSimultaneousConnections)

	if limited {
		b.manager.Decrement("v4.anonymous")

		return true
	} else {
		return false
	}
}
