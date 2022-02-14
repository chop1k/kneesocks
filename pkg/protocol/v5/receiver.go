package v5

import (
	"net"
	v52 "socks/internal/kneesocks/config/v5"
	"socks/pkg/utils"
	"time"
)

type Receiver struct {
	parser Parser
	buffer utils.BufferReader
}

func NewReceiver(
	parser Parser,
	buffer utils.BufferReader,
) (Receiver, error) {
	return Receiver{
		parser: parser,
		buffer: buffer,
	}, nil
}

func (b Receiver) ReceiveRequest(config v52.Config, conn net.Conn) (RequestChunk, error) {
	if config.Deadline.Request > 0 {
		deadlineErr := conn.SetReadDeadline(time.Now().Add(config.Deadline.Request))

		if deadlineErr != nil {
			return RequestChunk{}, deadlineErr
		}
	}

	chunk, err := b.buffer.Read(conn, 263)

	if err != nil {
		return RequestChunk{}, err
	}

	return b.parser.ParseRequest(chunk)
}
