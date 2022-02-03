package v4

import (
	"net"
	"socks/config/tcp"
	v42 "socks/config/v4"
	"socks/managers"
)

type Sender interface {
	SendFailAndClose(client net.Conn)
	SendSuccess(client net.Conn) error
	SendSuccessWithParameters(ip net.IP, port uint16, client net.Conn) error
}

type BaseSender struct {
	tcpConfig tcp.Config
	config    v42.DeadlineConfig
	deadline  managers.DeadlineManager
	builder   Builder
}

func NewBaseSender(
	tcpConfig tcp.Config,
	config v42.DeadlineConfig,
	deadline managers.DeadlineManager,
	builder Builder,
) (BaseSender, error) {
	return BaseSender{
		tcpConfig: tcpConfig,
		config:    config,
		deadline:  deadline,
		builder:   builder,
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

func (b BaseSender) send(status byte, ip net.IP, port uint16, client net.Conn) error {
	data, err := b.build(status, ip, port)

	if err != nil {
		return err
	}

	err = b.deadline.Write(b.config.GetResponseDeadline(), data, client)

	if err == managers.TimeoutError {
		return err
	}

	return nil
}

func (b BaseSender) SendFailAndClose(client net.Conn) {
	_ = b.send(91, net.IP{0, 0, 0, 0}, b.tcpConfig.GetBindPort(), client)
	_ = client.Close()
}

func (b BaseSender) SendSuccess(client net.Conn) error {
	return b.send(90, net.IP{0, 0, 0, 0}, b.tcpConfig.GetBindPort(), client)
}

func (b BaseSender) SendSuccessWithParameters(ip net.IP, port uint16, client net.Conn) error {
	return b.send(90, ip, port, client)
}
