package authenticator

import (
	"errors"
	"net"
	v52 "socks/internal/kneesocks/config/v5"
	"socks/internal/kneesocks/handlers/v5"
	password2 "socks/pkg/protocol/auth/password"
)

var (
	UserNotFoundError = errors.New("User not found. ")
)

type PasswordAuthenticator struct {
	errorHandler v5.ErrorHandler
	sender       password2.Sender
	receiver     password2.Receiver
}

func NewPasswordAuthenticator(
	errorHandler v5.ErrorHandler,
	sender password2.Sender,
	receiver password2.Receiver,
) (PasswordAuthenticator, error) {
	return PasswordAuthenticator{
		errorHandler: errorHandler,
		sender:       sender,
		receiver:     receiver,
	}, nil
}

func (b PasswordAuthenticator) Authenticate(config v52.Config, client net.Conn) (string, error) {
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
