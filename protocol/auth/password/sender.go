package password

import (
	"net"
	v5 "socks/config/v5"
	"socks/managers"
)

type Sender interface {
	SendResponse(code byte, client net.Conn) error
}

type BaseSender struct {
	config   v5.DeadlineConfig
	deadline managers.DeadlineManager
	builder  Builder
}

func NewBaseSender(
	config v5.DeadlineConfig,
	deadline managers.DeadlineManager,
	builder Builder,
) (BaseSender, error) {
	return BaseSender{
		config:   config,
		deadline: deadline,
		builder:  builder,
	}, nil
}

func (b BaseSender) SendResponse(code byte, client net.Conn) error {
	chunk, err := b.builder.BuildResponse(ResponseChunk{
		Version: 1,
		Status:  code,
	})

	if err != nil {
		return err
	}

	return b.deadline.Write(b.config.GetPasswordResponseDeadline(), chunk, client)
}
