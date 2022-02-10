package password

import (
	"net"
	v5 "socks/config/v5"
	"socks/protocol"
)

type Sender interface {
	SendResponse(code byte, client net.Conn) error
}

type BaseSender struct {
	config   v5.DeadlineConfig
	deadline protocol.Deadline
	builder  Builder
}

func NewBaseSender(
	config v5.DeadlineConfig,
	deadline protocol.Deadline,
	builder Builder,
) (BaseSender, error) {
	return BaseSender{
		config:   config,
		deadline: deadline,
		builder:  builder,
	}, nil
}

func (b BaseSender) SendResponse(code byte, client net.Conn) error {
	deadline, configErr := b.config.GetPasswordResponseDeadline()

	if configErr != nil {
		return configErr
	}

	chunk, err := b.builder.BuildResponse(ResponseChunk{
		Version: 1,
		Status:  code,
	})

	if err != nil {
		return err
	}

	return b.deadline.Write(deadline, chunk, client)
}
