package password

import (
	"net"
	v5 "socks/config/v5"
	"time"
)

type Sender interface {
	SendResponse(code byte, client net.Conn) error
}

type BaseSender struct {
	config  v5.DeadlineConfig
	builder Builder
}

func NewBaseSender(
	config v5.DeadlineConfig,
	builder Builder,
) (BaseSender, error) {
	return BaseSender{
		config:  config,
		builder: builder,
	}, nil
}

func (b BaseSender) SendResponse(code byte, client net.Conn) error {
	deadline, configErr := b.config.GetPasswordResponseDeadline()

	if configErr != nil {
		return configErr
	}

	deadlineErr := client.SetWriteDeadline(time.Now().Add(deadline))

	if deadlineErr != nil {
		return deadlineErr
	}

	chunk, err := b.builder.BuildResponse(ResponseChunk{
		Version: 1,
		Status:  code,
	})

	if err != nil {
		return err
	}

	_, err = client.Write(chunk)

	return err
}
