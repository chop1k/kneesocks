package helpers

import (
	"net"
	v4 "socks/config/v4"
	"time"
)

type Dialer interface {
	Dial(address string) (net.Conn, error)
}

type BaseDialer struct {
	config v4.DeadlineConfig
}

func NewBaseDialer(config v4.DeadlineConfig) (BaseDialer, error) {
	return BaseDialer{config: config}, nil
}

func (b BaseDialer) Dial(address string) (net.Conn, error) {
	deadline := time.Second * time.Duration(b.config.GetConnectDeadline())

	return net.DialTimeout("tcp4", address, deadline)
}
