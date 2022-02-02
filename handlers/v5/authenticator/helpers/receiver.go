package helpers

import (
	"io"
	v5 "socks/config/v5"
	"socks/managers"
	"socks/protocol/auth/password"
)

type Receiver interface {
	ReceivePassword(reader io.Reader) (password.RequestChunk, error)
}

type BaseReceiver struct {
	config   v5.DeadlineConfig
	deadline managers.DeadlineManager
	parser   password.Parser
}

func NewBaseReceiver(
	config v5.DeadlineConfig,
	deadline managers.DeadlineManager,
	parser password.Parser,
) (BaseReceiver, error) {
	return BaseReceiver{
		config:   config,
		deadline: deadline,
		parser:   parser,
	}, nil
}

func (b BaseReceiver) ReceivePassword(reader io.Reader) (password.RequestChunk, error) {
	data, err := b.deadline.Read(b.config.GetPasswordDeadline(), 515, reader)

	if err != nil {
		return password.RequestChunk{}, err
	}

	return b.parser.ParseRequest(data)
}
