package managers

import "fmt"

type UdpBindManager struct {
	bound map[string]bool
}

func NewUdpBindManager() (UdpBindManager, error) {
	return UdpBindManager{
		bound: make(map[string]bool),
	}, nil
}

func (u UdpBindManager) Bind(client string, host string) error {
	address := fmt.Sprintf("%s:%s", client, host)

	_, ok := u.bound[address]

	if ok {
		return AddressAlreadyBoundError
	} else {
		u.bound[address] = true
	}

	return nil
}

func (u UdpBindManager) IsBound(client string, host string) bool {
	address := fmt.Sprintf("%s:%s", client, host)

	_, ok := u.bound[address]

	return ok
}

func (u UdpBindManager) Unbind(client string, host string) {
	address := fmt.Sprintf("%s:%s", client, host)

	delete(u.bound, address)
}
