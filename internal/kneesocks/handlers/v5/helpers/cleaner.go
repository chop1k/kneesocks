package helpers

import (
	"fmt"
	"socks/internal/kneesocks/managers"
)

type Cleaner struct {
	manager *managers.ConnectionsManager
}

func NewCleaner(manager *managers.ConnectionsManager) (Cleaner, error) {
	return Cleaner{manager: manager}, nil
}

func (b Cleaner) Clean(name string) {
	id := fmt.Sprintf("v5.%s", name)

	b.manager.Decrement(id)
}
