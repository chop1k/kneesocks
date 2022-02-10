package helpers

import (
	"net"
	v4 "socks/config/v4"
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
	deadline, err := b.config.GetConnectDeadline()

	if err != nil {
		return nil, err
	}

	return net.DialTimeout("tcp4", address, deadline)
}
