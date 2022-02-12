package helpers

import (
	"fmt"
	"socks/managers"
)

type Cleaner interface {
	Clean(name string)
}

type BaseCleaner struct {
	manager *managers.ConnectionsManager
}

func NewBaseCleaner(manager *managers.ConnectionsManager) (BaseCleaner, error) {
	return BaseCleaner{manager: manager}, nil
}

func (b BaseCleaner) Clean(name string) {
	id := fmt.Sprintf("v5.%s", name)

	b.manager.Decrement(id)
}
