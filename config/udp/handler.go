package udp

import (
	"socks/config/tree"
	"time"
)

type Handler struct {
}

func NewHandler() (Handler, error) {
	return Handler{}, nil
}

func (h Handler) Handle(raw tree.UdpConfig) Config {
	return Config{
		Bind: BindConfig{
			Address: raw.Bind.Address,
			Port:    raw.Bind.Port,
		},
		Buffer: BufferConfig{
			PacketSize: raw.Buffer.PacketSize,
		},
		Deadline: DeadlineConfig{
			Read: time.Second * time.Duration(raw.Deadline.Read),
		},
	}
}
