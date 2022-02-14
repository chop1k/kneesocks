package tcp

import (
	"socks/internal/kneesocks/config/tree"
	"time"
)

type Handler struct {
}

func NewHandler() (Handler, error) {
	return Handler{}, nil
}

func (h Handler) Handle(raw tree.TcpConfig) Config {
	return Config{
		Bind: BindConfig{
			Address: raw.Bind.Address,
			Port:    raw.Bind.Port,
		},
		Buffer: BufferConfig{
			ClientSize: raw.Buffer.ClientSize,
			HostSize:   raw.Buffer.HostSize,
		},
		Deadline: DeadlineConfig{
			Welcome:  time.Second * time.Duration(raw.Deadline.Welcome),
			Exchange: time.Second * time.Duration(raw.Deadline.Exchange),
		},
	}
}
