package helpers

import (
	"net"
	v4 "socks/config/v4"
	"socks/managers"
)

type Binder interface {
	Bind(address string) error
	Receive(address string) (net.Conn, error)
	Send(address string, conn net.Conn) error
	Remove(address string)
}

type BaseBinder struct {
	config      v4.DeadlineConfig
	bindManager managers.BindManager
}

func NewBaseBinder(
	config v4.DeadlineConfig,
	bindManager managers.BindManager,
) (BaseBinder, error) {
	return BaseBinder{
		config:      config,
		bindManager: bindManager,
	}, nil
}

func (b BaseBinder) Bind(address string) error {
	return b.bindManager.Bind(address)
}

func (b BaseBinder) Receive(address string) (net.Conn, error) {
	deadline, err := b.config.GetBindDeadline()

	if err != nil {
		return nil, err
	}

	return b.bindManager.ReceiveHost(address, deadline)
}

func (b BaseBinder) Send(address string, conn net.Conn) error {
	return b.bindManager.SendClient(address, conn)
}

func (b BaseBinder) Remove(address string) {
	b.bindManager.Remove(address)
}