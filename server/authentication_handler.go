package server

import (
	"net"
	config2 "socks/config/tree"
	"socks/protocol/auth/password"
	v5 "socks/protocol/v5"
)

type AuthenticationHandler interface {
	HandleAuthentication(methods v5.MethodsChunk, client net.Conn) bool
}

type BaseAuthenticationHandler struct {
	config   config2.Config
	password password.Password
	protocol v5.Protocol
}

func NewBaseAuthenticationHandler(
	config config2.Config,
	password password.Password,
	protocol v5.Protocol,
) BaseAuthenticationHandler {
	return BaseAuthenticationHandler{
		config:   config,
		password: password,
		protocol: protocol,
	}
}

func (b BaseAuthenticationHandler) HandleAuthentication(methods v5.MethodsChunk, client net.Conn) bool {
	for _, method := range b.config.SocksV5.AuthenticationMethodsAllowed {
		var code byte

		if method == "no-authentication" {
			code = 0
		} else if method == "name/password" {
			code = 2
		} else {
			continue
		}

		for _, method := range methods.Methods {
			if code != method {
				continue
			}

			err := b.protocol.SelectMethod(code, client)

			if err != nil {
				return false
			}

			if code == 0 {
				return b.handleNoAuthentication()
			} else if code == 2 {
				return b.handlePassword(client)
			}

			return false
		}
	}

	_ = b.protocol.SelectMethod(255, client)

	return false
}

func (b BaseAuthenticationHandler) handleNoAuthentication() bool {
	return true
}

func (b BaseAuthenticationHandler) handlePassword(client net.Conn) bool {
	request, err := b.password.ReceiveRequest(client)

	if err != nil {
		return false
	}

	for _, user := range b.config.Users {
		if user.Name == request.Name && user.Password == request.Password {
			err := b.password.ResponseWith(0, client)

			if err != nil {
				return false
			}

			return true
		}
	}

	return false
}
