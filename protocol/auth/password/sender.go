package password

import (
	"net"
	v5 "socks/config/v5"
	"time"
)

type Sender interface {
	SendResponse(config v5.Config, code byte, client net.Conn) error
}

type BaseSender struct {
	builder Builder
}

func NewBaseSender(
	builder Builder,
) (BaseSender, error) {
	return BaseSender{
		builder: builder,
	}, nil
}

func (b BaseSender) SendResponse(config v5.Config, code byte, client net.Conn) error {
	deadlineErr := client.SetWriteDeadline(time.Now().Add(config.Deadline.PasswordResponse))

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
