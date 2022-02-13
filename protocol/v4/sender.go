package v4

import (
	"net"
	"socks/config/tcp"
	v42 "socks/config/v4"
	"time"
)

type Sender interface {
	SendFailAndClose(config v42.Config, client net.Conn)
	SendSuccess(config v42.Config, client net.Conn) error
	SendSuccessWithParameters(config v42.Config, ip net.IP, port uint16, client net.Conn) error
}

type BaseSender struct {
	bindConfig tcp.BindConfig
	builder    Builder
}

func NewBaseSender(
	bindConfig tcp.BindConfig,
	builder Builder,
) (BaseSender, error) {
	return BaseSender{
		bindConfig: bindConfig,
		builder:    builder,
	}, nil
}

func (b BaseSender) build(status byte, ip net.IP, port uint16) ([]byte, error) {
	return b.builder.BuildResponse(ResponseChunk{
		SocksVersion:    0,
		CommandCode:     status,
		DestinationPort: port,
		DestinationIp:   ip,
	})
}

func (b BaseSender) send(config v42.Config, status byte, ip net.IP, port uint16, client net.Conn) error {
	deadlineErr := client.SetWriteDeadline(time.Now().Add(config.Deadline.Response))

	if deadlineErr != nil {
		return deadlineErr
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

func (b BaseSender) SendFailAndClose(config v42.Config, client net.Conn) {
	_ = b.send(config, 91, net.IP{0, 0, 0, 0}, b.bindConfig.Port, client)
	_ = client.Close()
}

func (b BaseSender) SendSuccess(config v42.Config, client net.Conn) error {
	return b.send(config, 90, net.IP{0, 0, 0, 0}, b.bindConfig.Port, client)
}

func (b BaseSender) SendSuccessWithParameters(config v42.Config, ip net.IP, port uint16, client net.Conn) error {
	return b.send(config, 90, ip, port, client)
}
