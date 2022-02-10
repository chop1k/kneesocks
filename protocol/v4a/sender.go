package v4a

import (
	"net"
	"socks/config/tcp"
	v4a2 "socks/config/v4a"
	"socks/protocol"
)

type Sender interface {
	SendFailAndClose(client net.Conn)
	SendSuccess(client net.Conn) error
	SendSuccessWithParameters(ip net.IP, port uint16, client net.Conn) error
}

type BaseSender struct {
	tcpConfig tcp.BindConfig
	config    v4a2.DeadlineConfig
	deadline  protocol.Deadline
	builder   Builder
}

func NewBaseSender(
	tcpConfig tcp.BindConfig,
	config v4a2.DeadlineConfig,
	deadline protocol.Deadline,
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
	deadline, configErr := b.config.GetResponseDeadline()

	if configErr != nil {
		return configErr
	}

	data, err := b.build(status, ip, port)

	if err != nil {
		return err
	}

	err = b.deadline.Write(deadline, data, client)

	if err == protocol.TimeoutError {
		return err
	}

	return nil
}

func (b BaseSender) SendFailAndClose(client net.Conn) {
	port, err := b.tcpConfig.GetPort()

	if err != nil {
		panic(err)
	}

	_ = b.send(91, net.IP{0, 0, 0, 0}, port, client)
	_ = client.Close()
}

func (b BaseSender) SendSuccess(client net.Conn) error {
	port, err := b.tcpConfig.GetPort()

	if err != nil {
		panic(err)
	}

	return b.send(90, net.IP{0, 0, 0, 0}, port, client)
}

func (b BaseSender) SendSuccessWithParameters(ip net.IP, port uint16, client net.Conn) error {
	return b.send(90, ip, port, client)
}