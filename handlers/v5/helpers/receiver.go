package helpers

import (
	"io"
	"net"
	v52 "socks/config/v5"
	"socks/managers"
	v5 "socks/protocol/v5"
	"time"
)

type Receiver interface {
	ReceiveRequest(reader io.Reader) (v5.RequestChunk, error)
	ReceiveHost(address string) (net.Conn, error)
}

type BaseReceiver struct {
	config      v52.DeadlineConfig
	deadline    managers.DeadlineManager
	parser      v5.Parser
	bindManager managers.BindManager
}

func NewBaseReceiver(
	config v52.DeadlineConfig,
	deadline managers.DeadlineManager,
	parser v5.Parser,
	bindManager managers.BindManager,
) (BaseReceiver, error) {
	return BaseReceiver{
		config:      config,
		deadline:    deadline,
		parser:      parser,
		bindManager: bindManager,
	}, nil
}

func (b BaseReceiver) ReceiveRequest(reader io.Reader) (v5.RequestChunk, error) {
	chunk, err := b.deadline.Read(b.config.GetRequestDeadline(), 263, reader)

	if err != nil {
		return v5.RequestChunk{}, err
	}

	return b.parser.ParseRequest(chunk)
}

func (b BaseReceiver) ReceiveHost(address string) (net.Conn, error) {
	deadline := time.Second * time.Duration(b.config.GetBindDeadline())

	return b.bindManager.ReceiveHost(address, deadline)
}
