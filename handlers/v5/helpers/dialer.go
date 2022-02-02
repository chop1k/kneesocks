package helpers

import (
	"net"
	v5 "socks/config/v5"
	"time"
)

type Dialer interface {
	Dial(address string) (net.Conn, error)
}

type BaseDialer struct {
	config v5.DeadlineConfig
}

func NewBaseDialer(config v5.DeadlineConfig) (BaseDialer, error) {
	return BaseDialer{config: config}, nil
}

func (b BaseDialer) Dial(address string) (net.Conn, error) {
	deadline := time.Second * time.Duration(b.config.GetConnectDeadline())

	return net.DialTimeout("tcp", address, deadline)
}
