package helpers

import (
	"net"
	"socks/config/v4a"
	"socks/managers"
	"time"
)

type Receiver interface {
	ReceiveHost(address string) (net.Conn, error)
}

type BaseReceiver struct {
	config      v4a.DeadlineConfig
	bindManager managers.BindManager
}

func NewBaseReceiver(config v4a.DeadlineConfig, bindManager managers.BindManager) (BaseReceiver, error) {
	return BaseReceiver{config: config, bindManager: bindManager}, nil
}

func (b BaseReceiver) ReceiveHost(address string) (net.Conn, error) {
	deadline := time.Second * time.Duration(b.config.GetBindDeadline())

	return b.bindManager.ReceiveHost(address, deadline)
}
