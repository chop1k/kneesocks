package helpers

import (
	"net"
	"socks/config/v4a"
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
	deadline, err := b.config.GetConnectDeadline()

	if err != nil {
		return nil, err
	}

	return net.DialTimeout("tcp", address, deadline)
}
