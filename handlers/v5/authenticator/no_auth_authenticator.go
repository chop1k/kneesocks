package authenticator

import (
	"net"
	v5 "socks/config/v5"
)

type BaseNoAuthAuthenticator struct {
}

func NewBaseNoAuthAuthenticator() (BaseNoAuthAuthenticator, error) {
	return BaseNoAuthAuthenticator{}, nil
}

func (b BaseNoAuthAuthenticator) Authenticate(_ v5.Config, _ net.Conn) (string, error) {
	return "anonymous", nil
}
