package authenticator

import (
	"errors"
	"net"
	v52 "socks/config/v5"
	"socks/handlers/v5"
	"socks/protocol/auth/password"
)

var (
	UserNotFoundError = errors.New("User not found. ")
)

type BasePasswordAuthenticator struct {
	password     password.Password
	config       v52.Config
	errorHandler v5.ErrorHandler
}

func NewBasePasswordAuthenticator(password password.Password, config v52.Config, errorHandler v5.ErrorHandler) (BasePasswordAuthenticator, error) {
	return BasePasswordAuthenticator{
		password:     password,
		config:       config,
		errorHandler: errorHandler,
	}, nil
}

func (b BasePasswordAuthenticator) Authenticate(client net.Conn) (string, error) {
	request, err := b.password.ReceiveRequest(client)

	if err != nil {
		//b.errorHandler.HandlePasswordReceiveRequestError(err, client)

		return "", err
	}

	for name, user := range b.config.GetUsers() {
		if name == request.Name && user.Password == request.Password {
			err := b.password.ResponseWith(0, client)

			if err != nil {
				//b.errorHandler.HandlePasswordResponseError(err, user.Name, client)

				return "", err
			}

			return name, nil
		}
	}

	return "", UserNotFoundError
}
