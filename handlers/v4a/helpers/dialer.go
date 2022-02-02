package helpers

import (
	"net"
	"socks/config/v4a"
	"time"
)

type Dialer interface {
	Dial(address string) (net.Conn, error)
}

type BaseDialer struct {
	config v4a.DeadlineConfig
}

func NewBaseDialer(config v4a.DeadlineConfig) (BaseDialer, error) {
	return BaseDialer{config: config}, nil
}

func (b BaseDialer) Dial(address string) (net.Conn, error) {
	deadline := time.Second * time.Duration(b.config.GetConnectDeadline())

	return net.DialTimeout("tcp", address, deadline)
}
