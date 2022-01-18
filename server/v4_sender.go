package server

import (
	"net"
	"socks/config"
	v4 "socks/protocol/v4"
)

type V4Sender interface {
	SendFailAndClose(client net.Conn)
	SendSuccess(client net.Conn) error
	SendSuccessWithParameters(ip net.IP, port uint16, client net.Conn) error
}

type BaseV4Sender struct {
	protocol  v4.Protocol
	tcpConfig config.TcpConfig
}

func NewBaseV4Sender(protocol v4.Protocol, tcpConfig config.TcpConfig) (BaseV4Sender, error) {
	return BaseV4Sender{
		protocol:  protocol,
		tcpConfig: tcpConfig,
	}, nil
}

func (b BaseV4Sender) SendFailAndClose(client net.Conn) {
	_ = b.protocol.ResponseWithFail(uint16(b.tcpConfig.GetBindPort()), net.IP{0, 0, 0, 0}, client)
	_ = client.Close()
}

func (b BaseV4Sender) SendSuccess(client net.Conn) error {
	return b.protocol.ResponseWithSuccess(uint16(b.tcpConfig.GetBindPort()), net.IP{0, 0, 0, 0}, client)
}

func (b BaseV4Sender) SendSuccessWithParameters(ip net.IP, port uint16, client net.Conn) error {
	return b.protocol.ResponseWithSuccess(port, ip, client)
}
