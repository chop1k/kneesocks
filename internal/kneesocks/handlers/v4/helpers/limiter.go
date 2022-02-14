package helpers

import (
	"socks/internal/kneesocks/config/v4"
	"socks/internal/kneesocks/managers"
)

type Limiter struct {
	manager *managers.ConnectionsManager
}

func NewLimiter(
	manager *managers.ConnectionsManager,
) (Limiter, error) {
	return Limiter{
		manager: manager,
	}, nil
}

func (b Limiter) IsLimited(config v4.Config) bool {
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
