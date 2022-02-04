package protocol

import (
	"io"
	"socks/config/tcp"
	"socks/managers"
)

type Receiver interface {
	ReceiveWelcome(reader io.Reader) ([]byte, error)
}

type BaseReceiver struct {
	config      tcp.DeadlineConfig
	deadline    Deadline
	bindManager managers.BindManager
}

func NewBaseReceiver(
	config tcp.DeadlineConfig,
	deadline Deadline,
	bindManager managers.BindManager,
) (BaseReceiver, error) {
	return BaseReceiver{
		config:      config,
		deadline:    deadline,
		bindManager: bindManager,
	}, nil
}

func (b BaseReceiver) ReceiveWelcome(reader io.Reader) ([]byte, error) {
	data, err := b.deadline.Read(b.config.GetWelcomeDeadline(), 263, reader)

	if err != nil {
		return nil, err
	}

	return data, nil
}
