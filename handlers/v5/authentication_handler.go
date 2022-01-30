package v5

import (
	"errors"
	"net"
	"socks/config"
	v5 "socks/protocol/v5"
)

var (
	NoAuthenticationMethodsProvidedError = errors.New("No authentication methods provided. ")
	MethodUnsupportedError               = errors.New("Method unsupported. ")
)

type Authenticator interface {
	Authenticate(client net.Conn) (string, error)
}

type AuthenticationHandler interface {
	HandleAuthentication(methods v5.MethodsChunk, client net.Conn) (string, error)
}

type BaseAuthenticationHandler struct {
	protocol     v5.Protocol
	config       config.SocksV5Config
	errorHandler ErrorHandler
	password     Authenticator
	noAuth       Authenticator
}

func NewBaseAuthenticationHandler(
	protocol v5.Protocol,
	config config.SocksV5Config,
	errorHandler ErrorHandler,
	password Authenticator,
	noAuth Authenticator,
) BaseAuthenticationHandler {
	return BaseAuthenticationHandler{
		protocol:     protocol,
		config:       config,
		errorHandler: errorHandler,
		password:     password,
		noAuth:       noAuth,
	}
}

func (b BaseAuthenticationHandler) HandleAuthentication(methods v5.MethodsChunk, client net.Conn) (string, error) {
	for _, method := range b.config.GetAuthenticationMethods() {
		code := byte(255)

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

			return b.selectMethod(code, client)
		}
	}

	_ = b.protocol.SelectMethod(255, client)

	return "", NoAuthenticationMethodsProvidedError
}

func (b BaseAuthenticationHandler) selectMethod(code byte, client net.Conn) (string, error) {
	err := b.protocol.SelectMethod(code, client)

	if err != nil {
		b.errorHandler.HandleSelectMethodsError(err, client)

		return "", err
	}

	if code == 0 {
		return b.noAuth.Authenticate(client)
	} else if code == 2 {
		return b.password.Authenticate(client)
	} else {
		_ = b.protocol.SelectMethod(255, client)

		return "", MethodUnsupportedError
	}
}
