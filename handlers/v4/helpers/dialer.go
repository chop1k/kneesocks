package helpers

import (
	"net"
	v4 "socks/config/v4"
)

type Dialer interface {
	Dial(config v4.Config, address string) (net.Conn, error)
}

type BaseDialer struct {
}

func NewBaseDialer() (BaseDialer, error) {
	return BaseDialer{}, nil
}

func (b BaseDialer) Dial(config v4.Config, address string) (net.Conn, error) {
	return net.DialTimeout("tcp4", address, config.Deadline.Connect)
}
