package managers

type UdpHostManager struct {
	hosts map[string]string
}

func NewUdpHostManager() (UdpHostManager, error) {
	return UdpHostManager{
		hosts: make(map[string]string),
	}, nil
}

func (u UdpHostManager) Add(host string, client string) error {
	_, ok := u.hosts[host]

	if ok {
		return AddressAlreadyBoundError
	} else {
		u.hosts[host] = client
	}

	return nil
}

func (u UdpHostManager) Get(host string) (string, error) {
	client, ok := u.hosts[host]

	if !ok {
		return "", AddressNotExistsError
	}

	return client, nil
}

func (u UdpHostManager) Remove(host string) {
	delete(u.hosts, host)
}
