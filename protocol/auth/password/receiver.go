package password

import (
	"net"
	v5 "socks/config/v5"
	"socks/protocol"
)

type Receiver interface {
	ReceiveRequest(client net.Conn) (RequestChunk, error)
}

type BaseReceiver struct {
	config   v5.DeadlineConfig
	deadline protocol.Deadline
	parser   Parser
}

func NewBaseReceiver(
	config v5.DeadlineConfig,
	deadline protocol.Deadline,
	parser Parser,
) (BaseReceiver, error) {
	return BaseReceiver{
		config:   config,
		deadline: deadline,
		parser:   parser,
	}, nil
}

func (b BaseReceiver) ReceiveRequest(client net.Conn) (RequestChunk, error) {
	data, err := b.deadline.Read(b.config.GetPasswordDeadline(), 515, client)

	if err != nil {
		return RequestChunk{}, err
	}

	return b.parser.ParseRequest(data)
}
