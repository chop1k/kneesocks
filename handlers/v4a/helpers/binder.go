package helpers

import (
	"net"
	"socks/config/v4a"
	"socks/managers"
	"time"
)

type Binder interface {
	Bind(address string) error
	Receive(address string) (net.Conn, error)
	Send(address string, conn net.Conn) error
	Remove(address string)
}

type BaseBinder struct {
	config      v4a.DeadlineConfig
	bindManager managers.BindManager
}

func NewBaseBinder(
	config v4a.DeadlineConfig,
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
	deadline := time.Second * time.Duration(b.config.GetBindDeadline())

	return b.bindManager.ReceiveHost(address, deadline)
}

func (b BaseBinder) Send(address string, conn net.Conn) error {
	return b.bindManager.SendClient(address, conn)
}

func (b BaseBinder) Remove(address string) {
	b.bindManager.Remove(address)
}
