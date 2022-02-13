package protocol

import (
	"net"
	"socks/config/tcp"
	"socks/utils"
	"time"
)

type Receiver interface {
	ReceiveWelcome(config tcp.DeadlineConfig, conn net.Conn) ([]byte, error)
}

type BaseReceiver struct {
	buffer utils.BufferReader
}

func NewBaseReceiver(
	buffer utils.BufferReader,
) (BaseReceiver, error) {
	return BaseReceiver{
		buffer: buffer,
	}, nil
}

func (b BaseReceiver) ReceiveWelcome(config tcp.DeadlineConfig, conn net.Conn) ([]byte, error) {
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
