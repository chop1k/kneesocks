package managers

import (
	"errors"
	"net"
	"time"
)

var (
	AddressNotExistsError    = errors.New("Address does not exists. ")
	HostChannelClosedError   = errors.New("Host channel is closed. ")
	ClientChannelClosedError = errors.New("Client channel is closed. ")
	AddressAlreadyBoundError = errors.New("Address already bound. ")
	TimeoutError             = errors.New("Timeout exceeded. ")
)

type bundle struct {
	client chan net.Conn
	host   chan net.Conn
}

type BindManager struct {
	addresses map[string]bundle
}

func NewBindManager() BindManager {
	return BindManager{
		addresses: make(map[string]bundle),
	}
}

func (m BindManager) IsBound(addr string) bool {
	_, is := m.addresses[addr]

	return is
}

func (m BindManager) Bind(addr string) error {
	_, ok := m.addresses[addr]

	if ok {
		return AddressAlreadyBoundError
	}

	channels := bundle{
		client: make(chan net.Conn),
		host:   make(chan net.Conn),
	}

	m.addresses[addr] = channels

	return nil
}

func (m BindManager) Remove(addr string) {
	channel, is := m.addresses[addr]

	if is == false {
		return
	}

	close(channel.client)
	close(channel.host)
	delete(m.addresses, addr)
}

func (m BindManager) SendHost(address string, host net.Conn) error { // TODO: add deadline
	channel, ok := m.addresses[address]

	if !ok {
		return AddressNotExistsError
	}

	channel.host <- host

	return nil
}

func (m BindManager) SendClient(address string, client net.Conn) error { // TODO: add deadline
	channel, ok := m.addresses[address]

	if !ok {
		return AddressNotExistsError
	}

	channel.client <- client

	return nil
}

func (m BindManager) ReceiveClient(address string, deadline time.Duration) (net.Conn, error) {
	channel, ok := m.addresses[address]

	if !ok {
		return nil, AddressNotExistsError
	}

	timer := time.NewTimer(deadline)

	var client net.Conn
	var ko bool

	select {
	case client, ko = <-channel.client:
		timer.Stop()

		break
	case _ = <-timer.C:
		timer.Stop()

		return nil, TimeoutError
	}

	if !ko {
		return nil, ClientChannelClosedError
	}

	return client, nil
}

func (m BindManager) ReceiveHost(address string, deadline time.Duration) (net.Conn, error) {
	channel, ok := m.addresses[address]

	if !ok {
		return nil, AddressNotExistsError
	}

	timer := time.NewTimer(deadline)

	var host net.Conn
	var ko bool

	select {
	case host, ko = <-channel.host:
		timer.Stop()

		break
	case _ = <-timer.C:
		timer.Stop()

		return nil, TimeoutError
	}

	if !ko {
		return nil, HostChannelClosedError
	}

	return host, nil
}
