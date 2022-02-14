package authenticator

import (
	"net"
	"socks/internal/kneesocks/config/v5"
)

type NoAuthAuthenticator struct {
}

func NewNoAuthAuthenticator() (NoAuthAuthenticator, error) {
	return NoAuthAuthenticator{}, nil
}

func (b NoAuthAuthenticator) Authenticate(_ v5.Config, _ net.Conn) (string, error) {
	return "anonymous", nil
}
