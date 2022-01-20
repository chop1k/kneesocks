package server

import (
	"net"
	"socks/config"
	"socks/protocol/auth/password"
	v5 "socks/protocol/v5"
)

type V5AuthenticationHandler interface {
	HandleAuthentication(methods v5.MethodsChunk, client net.Conn) bool
}

type BaseV5AuthenticationHandler struct {
	password     password.Password
	protocol     v5.Protocol
	config       config.SocksV5Config
	users        config.UsersConfig
	errorHandler V5ErrorHandler
}

func NewBaseAuthenticationHandler(
	password password.Password,
	protocol v5.Protocol,
	config config.SocksV5Config,
	users config.UsersConfig,
	errorHandler V5ErrorHandler,
) BaseV5AuthenticationHandler {
	return BaseV5AuthenticationHandler{
		password:     password,
		protocol:     protocol,
		config:       config,
		users:        users,
		errorHandler: errorHandler,
	}
}

func (b BaseV5AuthenticationHandler) HandleAuthentication(methods v5.MethodsChunk, client net.Conn) bool {
	for _, method := range b.config.GetAuthenticationMethods() {
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
				b.errorHandler.HandleV5SelectMethodsError(err, client)

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

func (b BaseV5AuthenticationHandler) handleNoAuthentication() bool {
	return true
}

func (b BaseV5AuthenticationHandler) handlePassword(client net.Conn) bool {
	request, err := b.password.ReceiveRequest(client)

	if err != nil {
		b.errorHandler.HandleV5PasswordReceiveRequestError(err, client)

		return false
	}

	for _, user := range b.users.GetUsers() {
		if user.Name == request.Name && user.Password == request.Password {
			err := b.password.ResponseWith(0, client)

			if err != nil {
				b.errorHandler.HandleV5PasswordResponseError(err, user.Name, client)

				return false
			}

			return true
		}
	}

	return false
}
