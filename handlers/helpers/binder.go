package helpers

import (
	"net"
	"socks/config/tcp"
	"socks/managers"
	"time"
)

type Binder interface {
	IsBound(address string) bool
	Bind(address string) error
	Receive(address string) (net.Conn, error)
	Send(address string, conn net.Conn) error
	Remove(address string)
}

type BaseBinder struct {
	config      tcp.DeadlineConfig
	bindManager managers.BindManager
}

func NewBaseBinder(
	config tcp.DeadlineConfig,
	bindManager managers.BindManager,
) (BaseBinder, error) {
	return BaseBinder{
		config:      config,
		bindManager: bindManager,
	}, nil
}

func (b BaseBinder) IsBound(address string) bool {
	return b.bindManager.IsBound(address)
}

func (b BaseBinder) Bind(address string) error {
	return b.bindManager.Bind(address)
}

func (b BaseBinder) Receive(address string) (net.Conn, error) {
	deadline := time.Second * time.Duration(b.config.GetExchangeDeadline())

	return b.bindManager.ReceiveClient(address, deadline)
}

func (b BaseBinder) Send(address string, conn net.Conn) error {
	return b.bindManager.SendHost(address, conn)
}

func (b BaseBinder) Remove(address string) {
	b.bindManager.Remove(address)
}
