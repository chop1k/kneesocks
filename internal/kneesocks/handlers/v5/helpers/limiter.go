package helpers

import (
	"fmt"
	"socks/internal/kneesocks/config/tree"
	"socks/internal/kneesocks/config/v5"
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

func (b Limiter) IsLimited(config v5.Config, name string) bool {
	var limit tree.RateRestrictions

	user, ok := config.Users[name]

	if !ok {
		limit = tree.RateRestrictions{
			MaxSimultaneousConnections:  -1,
			HostReadBuffersPerSecond:    -1,
			HostWriteBuffersPerSecond:   -1,
			ClientReadBuffersPerSecond:  -1,
			ClientWriteBuffersPerSecond: -1,
		}
	} else {
		limit = user.Restrictions.Rate
	}

	if limit.MaxSimultaneousConnections <= 0 {
		return false
	}

	id := fmt.Sprintf("v5.%s", name)

	b.manager.Increment(id)

	limited := b.manager.IsExceed(id, limit.MaxSimultaneousConnections)

	if limited {
		b.manager.Decrement(id)

		return true
	} else {
		return false
	}
}
