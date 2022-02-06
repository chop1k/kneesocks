package transfer

import (
	"net"
)

type BaseHandler struct {
}

func NewBaseHandler() (BaseHandler, error) {
	return BaseHandler{}, nil
}

func (b BaseHandler) TransferToHost(client net.Conn, host net.Conn) {
	for {
		err := b.clientToHost(client, host)

		if err != nil {
			_ = client.Close()
			_ = host.Close()

			break
		}
	}
}

func (b BaseHandler) TransferToClient(client net.Conn, host net.Conn) {
	for {
		err := b.hostToClient(client, host)

		if err != nil {
			_ = client.Close()
			_ = host.Close()

			break
		}
	}
}

func (b BaseHandler) clientToHost(client, host net.Conn) error {
	buffer := make([]byte, 512)

	i, readErr := client.Read(buffer)

	if readErr != nil && i == 0 {
		return readErr
	}

	_, writeErr := host.Write(buffer[:i])

	if writeErr != nil {
		return writeErr
	}

	return readErr
}

func (b BaseHandler) hostToClient(client, host net.Conn) error {
	buffer := make([]byte, 512)

	i, readErr := host.Read(buffer)

	if readErr != nil && i == 0 {
		return readErr
	}

	_, writeErr := client.Write(buffer[:i])

	if writeErr != nil {
		return writeErr
	}

	return readErr
}
