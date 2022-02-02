package helpers

import (
	"io"
	"net"
	"socks/config/tcp"
	"socks/managers"
	"time"
)

type Receiver interface {
	ReceiveWelcome(reader io.Reader) ([]byte, error)
	ReceiveClient(address string) (net.Conn, error)
}

type BaseReceiver struct {
	config      tcp.DeadlineConfig
	deadline    managers.DeadlineManager
	bindManager managers.BindManager
}

func NewBaseReceiver(
	config tcp.DeadlineConfig,
	deadline managers.DeadlineManager,
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

func (b BaseReceiver) ReceiveClient(address string) (net.Conn, error) {
	deadline := time.Second * time.Duration(b.config.GetExchangeDeadline())

	return b.bindManager.ReceiveClient(address, deadline)
}
