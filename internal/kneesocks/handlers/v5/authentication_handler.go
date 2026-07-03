package v5

import (
	"errors"
	"net"
	v52 "socks/internal/kneesocks/config/v5"
	v53 "socks/pkg/protocol/v5"
)

var (
	NoAuthenticationMethodsProvidedError = errors.New("no authentication methods provided")
	MethodUnsupportedError               = errors.New("method unsupported")
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

		switch method {
		case "no-authentication":
			code = 0
		case "name/password":
			code = 2
		default:
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

	switch code {
	case 0:
		return b.noAuth.Authenticate(config, client)
	case 2:
		return b.password.Authenticate(config, client)
	default:
		_ = b.sender.SendMethodSelection(config, 255, client)

		return "", MethodUnsupportedError
	}
}
