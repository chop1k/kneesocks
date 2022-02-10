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
	deadline, configErr := b.config.GetRequestDeadline()

	if configErr != nil {
		return RequestChunk{}, configErr
	}

	chunk, err := b.deadline.Read(deadline, 263, reader)

	if err != nil {
		return RequestChunk{}, err
	}

	return b.parser.ParseRequest(chunk)
}
