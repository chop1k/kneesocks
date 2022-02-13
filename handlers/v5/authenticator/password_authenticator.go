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
	errorHandler v5.ErrorHandler
	sender       password.Sender
	receiver     password.Receiver
}

func NewBasePasswordAuthenticator(
	errorHandler v5.ErrorHandler,
	sender password.Sender,
	receiver password.Receiver,
) (BasePasswordAuthenticator, error) {
	return BasePasswordAuthenticator{
		errorHandler: errorHandler,
		sender:       sender,
		receiver:     receiver,
	}, nil
}

func (b BasePasswordAuthenticator) Authenticate(config v52.Config, client net.Conn) (string, error) {
	request, err := b.receiver.ReceiveRequest(config, client)

	if err != nil {
		b.errorHandler.HandlePasswordReceiveRequestError(err, client)

		return "", err
	}

	user, ok := config.Users[request.Name]

	if !ok {
		return "", UserNotFoundError
	}

	if user.Password == request.Password {
		err := b.sender.SendResponse(config, 0, client)

		if err != nil {
			b.errorHandler.HandlePasswordResponseError(err, request.Name, client)

			return "", err
		}

		return request.Name, nil
	}

	return "", UserNotFoundError
}
