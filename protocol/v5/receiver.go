package v5

import (
	"io"
	v52 "socks/config/v5"
	"socks/protocol"
)

type Receiver interface {
	ReceiveRequest(reader io.Reader) (RequestChunk, error)
}

type BaseReceiver struct {
	config   v52.DeadlineConfig
	deadline protocol.Deadline
	parser   Parser
}

func NewBaseReceiver(
	config v52.DeadlineConfig,
	deadline protocol.Deadline,
	parser Parser,
) (BaseReceiver, error) {
	return BaseReceiver{
		config:   config,
		deadline: deadline,
		parser:   parser,
	}, nil
}

func (b BaseReceiver) ReceiveRequest(reader io.Reader) (RequestChunk, error) {
	chunk, err := b.deadline.Read(b.config.GetRequestDeadline(), 263, reader)

	if err != nil {
		return RequestChunk{}, err
	}

	return b.parser.ParseRequest(chunk)
}
