package password

import (
	"net"
	v5 "socks/config/v5"
	"socks/utils"
	"time"
)

type Receiver interface {
	ReceiveRequest(conn net.Conn) (RequestChunk, error)
}

type BaseReceiver struct {
	config v5.DeadlineConfig
	parser Parser
	buffer utils.BufferReader
}

func NewBaseReceiver(
	config v5.DeadlineConfig,
	parser Parser,
	buffer utils.BufferReader,
) (BaseReceiver, error) {
	return BaseReceiver{
		config: config,
		parser: parser,
		buffer: buffer,
	}, nil
}

func (b BaseReceiver) ReceiveRequest(conn net.Conn) (RequestChunk, error) {
	deadline, configErr := b.config.GetPasswordDeadline()

	if configErr != nil {
		return RequestChunk{}, configErr
	}

	deadlineErr := conn.SetReadDeadline(time.Now().Add(deadline))

	if deadlineErr != nil {
		return RequestChunk{}, deadlineErr
	}

	data, err := b.buffer.Read(conn, 515)

	if err != nil {
		return RequestChunk{}, err
	}

	return b.parser.ParseRequest(data)
}
