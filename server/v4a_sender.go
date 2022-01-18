package server

import (
	"net"
	"socks/config"
	"socks/protocol/v4a"
)

type V4aSender interface {
	SendFailAndClose(client net.Conn)
	SendSuccess(client net.Conn) error
	SendSuccessWithParameters(ip net.IP, port uint16, client net.Conn) error
}

type BaseV4aSender struct {
	protocol  v4a.Protocol
	tcpConfig config.TcpConfig
}

func NewBaseV4aSender(protocol v4a.Protocol, tcpConfig config.TcpConfig) (BaseV4aSender, error) {
	return BaseV4aSender{
		protocol:  protocol,
		tcpConfig: tcpConfig,
	}, nil
}

func (b BaseV4aSender) SendFailAndClose(client net.Conn) {
	_ = b.protocol.ResponseWithFail(uint16(b.tcpConfig.GetBindPort()), net.IP{0, 0, 0, 0}, client)
	_ = client.Close()
}

func (b BaseV4aSender) SendSuccess(client net.Conn) error {
	return b.protocol.ResponseWithSuccess(uint16(b.tcpConfig.GetBindPort()), net.IP{0, 0, 0, 0}, client)
}

func (b BaseV4aSender) SendSuccessWithParameters(ip net.IP, port uint16, client net.Conn) error {
	return b.protocol.ResponseWithSuccess(port, ip, client)
}
