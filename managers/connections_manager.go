package managers

import "sync"

type ConnectionsManager struct {
	connections map[string]int
	mutex       sync.Mutex
}

func NewConnectionsManager() (*ConnectionsManager, error) {
	return &ConnectionsManager{
		connections: make(map[string]int),
	}, nil
}

func (c *ConnectionsManager) Increment(name string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	i, ok := c.connections[name]

	if !ok {
		c.connections[name] = 0
	}

	c.connections[name] = i + 1
}

func (c *ConnectionsManager) IsExceed(name string, limit int) bool {
	i, ok := c.connections[name]

	if !ok {
		return true
	}

	return i >= limit
}

func (c *ConnectionsManager) Decrement(name string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	i, ok := c.connections[name]

	if !ok {
		c.connections[name] = 0
	}

	if i <= 0 {
		c.connections[name] = 0
	} else {
		c.connections[name] = i - 1
	}
}
