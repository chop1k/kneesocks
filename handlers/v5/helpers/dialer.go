package helpers

import (
	"net"
	v5 "socks/config/v5"
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
	deadline, err := b.config.GetConnectDeadline()

	if err != nil {
		return nil, err
	}

	return net.DialTimeout("tcp", address, deadline)
}
