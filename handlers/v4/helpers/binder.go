package helpers

import (
	"net"
	v4 "socks/config/v4"
	"socks/managers"
)

type Binder interface {
	Bind(address string) error
	Receive(config v4.Config, address string) (net.Conn, error)
	Send(address string, conn net.Conn) error
	Remove(address string)
}

type BaseBinder struct {
	bindManager managers.BindManager
}

func NewBaseBinder(
	bindManager managers.BindManager,
) (BaseBinder, error) {
	return BaseBinder{
		bindManager: bindManager,
	}, nil
}

func (b BaseBinder) Bind(address string) error {
	return b.bindManager.Bind(address)
}

func (b BaseBinder) Receive(config v4.Config, address string) (net.Conn, error) {
	return b.bindManager.ReceiveHost(address, config.Deadline.Bind)
}

func (b BaseBinder) Send(address string, conn net.Conn) error {
	return b.bindManager.SendClient(address, conn)
}

func (b BaseBinder) Remove(address string) {
	b.bindManager.Remove(address)
}
