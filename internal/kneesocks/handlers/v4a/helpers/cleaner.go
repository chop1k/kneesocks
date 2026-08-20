package helpers

import (
	"socks/internal/kneesocks/managers"
)

type Cleaner struct {
	manager *managers.ConnectionsManager
}

func NewCleaner(manager *managers.ConnectionsManager) Cleaner {
	return Cleaner{manager: manager}
}

func (b Cleaner) Clean() {
	b.manager.Decrement("v4a.anonymous")
}
