package helpers

import (
	"socks/managers"
)

type Cleaner struct {
	manager *managers.ConnectionsManager
}

func NewCleaner(manager *managers.ConnectionsManager) (Cleaner, error) {
	return Cleaner{manager: manager}, nil
}

func (b Cleaner) Clean() {
	b.manager.Decrement("v4a.anonymous")
}
