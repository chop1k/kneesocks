package protocol

import (
	"net"
	"socks/internal/kneesocks/config/tcp"
	"socks/pkg/utils"
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
	if config.Welcome > 0 {
		deadlineErr := conn.SetReadDeadline(time.Now().Add(config.Welcome))

		if deadlineErr != nil {
			return nil, deadlineErr
		}
	}

	data, err := b.buffer.Read(conn, 263)

	if err != nil {
		return nil, err
	}

	return data, nil
}
