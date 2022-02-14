package protocol

import (
	"net"
	"socks/config/tcp"
	"socks/utils"
	"time"
)

type Receiver struct {
	buffer utils.BufferReader
}

func NewReceiver(
	buffer utils.BufferReader,
) (Receiver, error) {
	return Receiver{
		buffer: buffer,
	}, nil
}

func (b Receiver) ReceiveWelcome(config tcp.DeadlineConfig, conn net.Conn) ([]byte, error) {
	deadlineErr := conn.SetReadDeadline(time.Now().Add(config.Welcome))

	if deadlineErr != nil {
		return nil, deadlineErr
	}

	data, err := b.buffer.Read(conn, 263)

	if err != nil {
		return nil, err
	}

	return data, nil
}
