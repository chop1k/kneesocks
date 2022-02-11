package v5

import (
	"net"
	v52 "socks/config/v5"
	"socks/utils"
	"time"
)

type Receiver interface {
	ReceiveRequest(conn net.Conn) (RequestChunk, error)
}

type BaseReceiver struct {
	config v52.DeadlineConfig
	parser Parser
	buffer utils.BufferReader
}

func NewBaseReceiver(
	config v52.DeadlineConfig,
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
	deadline, configErr := b.config.GetRequestDeadline()

	if configErr != nil {
		return RequestChunk{}, configErr
	}

	deadlineErr := conn.SetReadDeadline(time.Now().Add(deadline))

	if deadlineErr != nil {
		return RequestChunk{}, deadlineErr
	}

	chunk, err := b.buffer.Read(conn, 263)

	if err != nil {
		return RequestChunk{}, err
	}

	return b.parser.ParseRequest(chunk)
}
