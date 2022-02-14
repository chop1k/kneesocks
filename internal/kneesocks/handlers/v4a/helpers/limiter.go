package helpers

import (
	"socks/internal/kneesocks/config/v4a"
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

func (b Limiter) IsLimited(config v4a.Config) bool {
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
