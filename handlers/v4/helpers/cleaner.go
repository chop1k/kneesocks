package helpers

import (
	"socks/managers"
)

type Cleaner interface {
	Clean() error
}

type BaseCleaner struct {
	manager *managers.ConnectionsManager
}

func NewBaseCleaner(manager *managers.ConnectionsManager) (BaseCleaner, error) {
	return BaseCleaner{manager: manager}, nil
}

func (b BaseCleaner) Clean() error {
	b.manager.Decrement("v4.anonymous")

	return nil
}
