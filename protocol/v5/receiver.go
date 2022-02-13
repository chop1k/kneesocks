package v5

import (
	"net"
	v52 "socks/config/v5"
	"socks/utils"
	"time"
)

type Receiver interface {
	ReceiveRequest(config v52.Config, conn net.Conn) (RequestChunk, error)
}

type BaseReceiver struct {
	parser Parser
	buffer utils.BufferReader
}

func NewBaseReceiver(
	parser Parser,
	buffer utils.BufferReader,
) (BaseReceiver, error) {
	return BaseReceiver{
		parser: parser,
		buffer: buffer,
	}, nil
}

func (b BaseReceiver) ReceiveRequest(config v52.Config, conn net.Conn) (RequestChunk, error) {
	deadlineErr := conn.SetReadDeadline(time.Now().Add(config.Deadline.Request))

	if deadlineErr != nil {
		return RequestChunk{}, deadlineErr
	}

	chunk, err := b.buffer.Read(conn, 263)

	if err != nil {
		return RequestChunk{}, err
	}

	return b.parser.ParseRequest(chunk)
}
