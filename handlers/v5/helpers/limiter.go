package helpers

import (
	"fmt"
	"socks/config/tree"
	v5 "socks/config/v5"
	"socks/managers"
)

type Limiter interface {
	IsLimited(name string) bool
}

type BaseLimiter struct {
	config  v5.UsersConfig
	manager *managers.ConnectionsManager
}

func NewBaseLimiter(
	config v5.UsersConfig,
	manager *managers.ConnectionsManager,
) (BaseLimiter, error) {
	return BaseLimiter{
		config:  config,
		manager: manager,
	}, nil
}

func (b BaseLimiter) IsLimited(name string) bool {
	limit, err := b.config.GetRate(name)

	if err != nil && err == v5.UserNotExistsError {
		limit = tree.RateRestrictions{
			MaxSimultaneousConnections:  -1,
			HostReadBuffersPerSecond:    -1,
			HostWriteBuffersPerSecond:   -1,
			ClientReadBuffersPerSecond:  -1,
			ClientWriteBuffersPerSecond: -1,
		}
	} else if err != nil {
		panic(err)
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
