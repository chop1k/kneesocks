package helpers

import (
	"socks/managers"
)

type Cleaner interface {
	Clean()
}

type BaseCleaner struct {
	manager *managers.ConnectionsManager
}

func NewBaseCleaner(manager *managers.ConnectionsManager) (BaseCleaner, error) {
	return BaseCleaner{manager: manager}, nil
}

func (b BaseCleaner) Clean() {
	b.manager.Decrement("v4a.anonymous")
}