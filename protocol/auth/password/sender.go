package password

import (
	"net"
	v5 "socks/config/v5"
	"time"
)

type Sender struct {
	builder Builder
}

func NewSender(
	builder Builder,
) (Sender, error) {
	return Sender{
		builder: builder,
	}, nil
}

func (b Sender) SendResponse(config v5.Config, code byte, client net.Conn) error {
	if config.Deadline.PasswordResponse > 0 {
		deadlineErr := client.SetWriteDeadline(time.Now().Add(config.Deadline.PasswordResponse))

		if deadlineErr != nil {
			return deadlineErr
		}
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
