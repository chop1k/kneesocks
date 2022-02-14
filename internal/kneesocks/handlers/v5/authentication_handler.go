package v5

import (
	"errors"
	"net"
	v52 "socks/internal/kneesocks/config/v5"
	v53 "socks/pkg/protocol/v5"
)

var (
	NoAuthenticationMethodsProvidedError = errors.New("No authentication methods provided. ")
	MethodUnsupportedError               = errors.New("Method unsupported. ")
)

type Authenticator interface {
	Authenticate(config v52.Config, client net.Conn) (string, error)
}

type AuthenticationHandler struct {
	errorHandler ErrorHandler
	password     Authenticator
	noAuth       Authenticator
	sender       v53.Sender
}

func NewAuthenticationHandler(
	errorHandler ErrorHandler,
	password Authenticator,
	noAuth Authenticator,
	sender v53.Sender,
) AuthenticationHandler {
	return AuthenticationHandler{
		errorHandler: errorHandler,
		password:     password,
		noAuth:       noAuth,
		sender:       sender,
	}
}

func (b AuthenticationHandler) HandleAuthentication(config v52.Config, methods v53.MethodsChunk, client net.Conn) (string, error) {
	_methods := config.AuthenticationMethodsAllowed

	for _, method := range _methods {
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

			return b.selectMethod(config, code, client)
		}
	}

	_ = b.sender.SendMethodSelection(config, 255, client)

	return "", NoAuthenticationMethodsProvidedError
}

func (b AuthenticationHandler) selectMethod(config v52.Config, code byte, client net.Conn) (string, error) {
	err := b.sender.SendMethodSelection(config, code, client)

	if err != nil {
		b.errorHandler.HandleMethodSelectionError(config, err, client)

		return "", err
	}

	if code == 0 {
		return b.noAuth.Authenticate(config, client)
	} else if code == 2 {
		return b.password.Authenticate(config, client)
	} else {
		_ = b.sender.SendMethodSelection(config, 255, client)

		return "", MethodUnsupportedError
	}
}
