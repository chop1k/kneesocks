package authenticator

import "net"

type BaseNoAuthAuthenticator struct {
}

func NewBaseNoAuthAuthenticator() (BaseNoAuthAuthenticator, error) {
	return BaseNoAuthAuthenticator{}, nil
}

func (b BaseNoAuthAuthenticator) Authenticate(_ net.Conn) (string, error) {
	return "anonymous", nil
}
