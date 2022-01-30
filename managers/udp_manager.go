package managers

import (
	"errors"
	"net"
)

var (
	ClientNotExistsError = errors.New("Client is not exists. ")
)

type UdpAssociationManager struct {
	clients     map[string]net.Addr
	clientHosts map[string][]string
	hosts       map[string]net.Addr
}

func NewUdpAssociationManager() UdpAssociationManager {
	return UdpAssociationManager{
		clients:     make(map[string]net.Addr),
		clientHosts: make(map[string][]string),
		hosts:       make(map[string]net.Addr),
	}
}

func (u UdpAssociationManager) Set(client string, addr net.Addr) {
	u.clients[client] = addr
	u.clientHosts[client] = []string{}
}

func (u UdpAssociationManager) Get(client string) (net.Addr, bool) {
	addr, ok := u.clients[client]

	return addr, ok
}

func (u UdpAssociationManager) AddHost(host, clientAddr string, client net.Addr) error {
	hosts, ok := u.clientHosts[clientAddr]

	if !ok {
		return ClientNotExistsError
	}

	hosts = append(hosts, host)

	u.clientHosts[clientAddr] = hosts

	u.hosts[host] = client

	return nil
}

func (u UdpAssociationManager) GetHost(host string) (net.Addr, bool) {
	addr, ok := u.hosts[host]

	return addr, ok
}

func (u UdpAssociationManager) Remove(client string) {
	hosts, ok := u.clientHosts[client]

	if ok {
		for _, host := range hosts {
			delete(u.hosts, host)
		}

		delete(u.clientHosts, client)
	}

	delete(u.clients, client)
}
