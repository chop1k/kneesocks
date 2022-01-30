package transfer

import (
	"net"
)

type StreamHandler interface {
	ClientToHost(client net.Conn, host net.Conn)
	HostToClient(client net.Conn, host net.Conn)
}

type BaseStreamHandler struct {
}

func NewBaseStreamHandler() *BaseStreamHandler {
	return &BaseStreamHandler{}
}

func (b BaseStreamHandler) ClientToHost(client, host net.Conn) {
	for {
		err := b.clientToHost(client, host)

		if err != nil {
			_ = client.Close()
			_ = host.Close()

			break
		}
	}
}

func (b BaseStreamHandler) HostToClient(client, host net.Conn) {
	for {
		err := b.hostToClient(client, host)

		if err != nil {
			_ = client.Close()
			_ = host.Close()

			break
		}
	}
}

func (b BaseStreamHandler) clientToHost(client, host net.Conn) error {
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

func (b BaseStreamHandler) hostToClient(client, host net.Conn) error {
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
