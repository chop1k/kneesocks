package v4a

import (
	"net"
	"socks/config/tcp"
	"socks/protocol/v4a"
)

type Sender interface {
	SendFailAndClose(client net.Conn)
	SendSuccess(client net.Conn) error
	SendSuccessWithParameters(ip net.IP, port uint16, client net.Conn) error
}

type BaseSender struct {
	protocol  v4a.Protocol
	tcpConfig tcp.Config
}

func NewBaseSender(protocol v4a.Protocol, tcpConfig tcp.Config) (BaseSender, error) {
	return BaseSender{
		protocol:  protocol,
		tcpConfig: tcpConfig,
	}, nil
}

func (b BaseSender) SendFailAndClose(client net.Conn) {
	_ = b.protocol.ResponseWithFail(b.tcpConfig.GetBindPort(), net.IP{0, 0, 0, 0}, client)
	_ = client.Close()
}

func (b BaseSender) SendSuccess(client net.Conn) error {
	return b.protocol.ResponseWithSuccess(b.tcpConfig.GetBindPort(), net.IP{0, 0, 0, 0}, client)
}

func (b BaseSender) SendSuccessWithParameters(ip net.IP, port uint16, client net.Conn) error {
	return b.protocol.ResponseWithSuccess(port, ip, client)
}