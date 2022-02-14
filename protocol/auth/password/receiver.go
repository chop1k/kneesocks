package password

import (
	"net"
	v5 "socks/config/v5"
	"socks/utils"
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

func (b Receiver) ReceiveRequest(config v5.Config, conn net.Conn) (RequestChunk, error) {
	if config.Deadline.Password > 0 {
		deadlineErr := conn.SetReadDeadline(time.Now().Add(config.Deadline.Password))

		if deadlineErr != nil {
			return RequestChunk{}, deadlineErr
		}
	}

	data, err := b.buffer.Read(conn, 515)

	if err != nil {
		return RequestChunk{}, err
	}

	return b.parser.ParseRequest(data)
}
