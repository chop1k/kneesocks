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
	config       v52.UsersConfig
	errorHandler v5.ErrorHandler
	sender       password.Sender
	receiver     password.Receiver
}

func NewBasePasswordAuthenticator(
	config v52.UsersConfig,
	errorHandler v5.ErrorHandler,
	sender password.Sender,
	receiver password.Receiver,
) (BasePasswordAuthenticator, error) {
	return BasePasswordAuthenticator{
		config:       config,
		errorHandler: errorHandler,
		sender:       sender,
		receiver:     receiver,
	}, nil
}

func (b BasePasswordAuthenticator) Authenticate(client net.Conn) (string, error) {
	request, err := b.receiver.ReceiveRequest(client)

	if err != nil {
		b.errorHandler.HandlePasswordReceiveRequestError(err, client)

		return "", err
	}

	user, configErr := b.config.GetUser(request.Name)

	if configErr != nil {
		return "", configErr
	}

	if user.Password == request.Password {
		err := b.sender.SendResponse(0, client)

		if err != nil {
			b.errorHandler.HandlePasswordResponseError(err, request.Name, client)

			return "", err
		}

		return request.Name, nil
	}

	return "", UserNotFoundError
}
