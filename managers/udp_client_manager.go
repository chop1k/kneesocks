package managers

import (
	"errors"
	"github.com/emirpasic/gods/sets/hashset"
)

var (
	ClientNotExistsError = errors.New("Client is not exists. ")
	NilPointerError      = errors.New("Got nil instead of set. ")
)

type UdpClientManager struct {
	clients map[string]*hashset.Set
}

func NewUdpClientManager() (UdpClientManager, error) {
	return UdpClientManager{
		clients: make(map[string]*hashset.Set),
	}, nil
}

func (u UdpClientManager) Add(client string) error {
	_, ok := u.clients[client]

	if ok {
		return AddressAlreadyBoundError
	} else {
		u.clients[client] = hashset.New()
	}

	return nil
}

func (u UdpClientManager) Bind(client string, host string) error {
	_, ok := u.clients[client]

	if ok {
		return AddressAlreadyBoundError
	} else {
		set := u.clients[client]

		if set == nil {
			return NilPointerError
		}

		set.Add(host)
	}

	return nil
}

func (u UdpClientManager) IsBound(address string) bool {
	_, ok := u.clients[address]

	return ok
}

func (u UdpClientManager) Remove(client string) {
	delete(u.clients, client)
}
