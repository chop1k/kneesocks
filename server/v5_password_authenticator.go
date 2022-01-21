package server

import (
	"errors"
	"net"
	"socks/config"
	"socks/protocol/auth/password"
)

var (
	UserNotFoundError = errors.New("User not found. ")
)

type BaseV5PasswordAuthenticator struct {
	password     password.Password
	users        config.UsersConfig
	errorHandler V5ErrorHandler
}

func NewBaseV5PasswordAuthenticator(password password.Password, users config.UsersConfig, errorHandler V5ErrorHandler) (BaseV5PasswordAuthenticator, error) {
	return BaseV5PasswordAuthenticator{
		password:     password,
		users:        users,
		errorHandler: errorHandler,
	}, nil
}

func (b BaseV5PasswordAuthenticator) Authenticate(client net.Conn) (string, error) {
	request, err := b.password.ReceiveRequest(client)

	if err != nil {
		//b.errorHandler.HandleV5PasswordReceiveRequestError(err, client)

		return "", err
	}

	for _, user := range b.users.GetUsers() {
		if user.Name == request.Name && user.Password == request.Password {
			err := b.password.ResponseWith(0, client)

			if err != nil {
				//b.errorHandler.HandleV5PasswordResponseError(err, user.Name, client)

				return "", err
			}

			return user.Name, nil
		}
	}

	return "", UserNotFoundError
}
