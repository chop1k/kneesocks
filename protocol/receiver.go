package protocol

import (
	"net"
	"socks/config/tcp"
	"socks/utils"
	"time"
)

type Receiver interface {
	ReceiveWelcome(conn net.Conn) ([]byte, error)
}

type BaseReceiver struct {
	config tcp.DeadlineConfig
	buffer utils.BufferReader
}

func NewBaseReceiver(
	config tcp.DeadlineConfig,
	buffer utils.BufferReader,
) (BaseReceiver, error) {
	return BaseReceiver{
		config: config,
		buffer: buffer,
	}, nil
}

func (b BaseReceiver) ReceiveWelcome(conn net.Conn) ([]byte, error) {
	deadline, configErr := b.config.GetWelcomeDeadline()

	if configErr != nil {
		return nil, configErr
	}

	deadlineErr := conn.SetReadDeadline(time.Now().Add(deadline))

	if deadlineErr != nil {
		return nil, deadlineErr
	}

	data, err := b.buffer.Read(conn, 263)

	if err != nil {
		return nil, err
	}

	return data, nil
}
