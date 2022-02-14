package v4

import (
	"net"
	"socks/internal/kneesocks/config/tcp"
	v42 "socks/internal/kneesocks/config/v4"
	"time"
)

type Sender struct {
	bindConfig tcp.BindConfig
	builder    Builder
}

func NewSender(
	bindConfig tcp.BindConfig,
	builder Builder,
) (Sender, error) {
	return Sender{
		bindConfig: bindConfig,
		builder:    builder,
	}, nil
}

func (b Sender) build(status byte, ip net.IP, port uint16) ([]byte, error) {
	return b.builder.BuildResponse(ResponseChunk{
		SocksVersion:    0,
		CommandCode:     status,
		DestinationPort: port,
		DestinationIp:   ip,
	})
}

func (b Sender) send(config v42.Config, status byte, ip net.IP, port uint16, client net.Conn) error {
	if config.Deadline.Response > 0 {
		deadlineErr := client.SetWriteDeadline(time.Now().Add(config.Deadline.Response))

		if deadlineErr != nil {
			return deadlineErr
		}
	}

	data, err := b.build(status, ip, port)

	if err != nil {
		return err
	}

	_, err = client.Write(data)

	if err != nil {
		return err
	}

	return nil
}

func (b Sender) SendFailAndClose(config v42.Config, client net.Conn) {
	_ = b.send(config, 91, net.IP{0, 0, 0, 0}, b.bindConfig.Port, client)
	_ = client.Close()
}

func (b Sender) SendSuccess(config v42.Config, client net.Conn) error {
	return b.send(config, 90, net.IP{0, 0, 0, 0}, b.bindConfig.Port, client)
}

func (b Sender) SendSuccessWithParameters(config v42.Config, ip net.IP, port uint16, client net.Conn) error {
	return b.send(config, 90, ip, port, client)
}
