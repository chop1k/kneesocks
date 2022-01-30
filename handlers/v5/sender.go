package v5

import (
	"net"
	"socks/config"
	v5 "socks/protocol/v5"
)

type Sender interface {
	SendConnectionRefusedAndClose(client net.Conn)
	SendNetworkUnreachableAndClose(client net.Conn)
	SendHostUnreachableAndClose(client net.Conn)
	SendCommandNotSupportedAndClose(client net.Conn)
	SendAddressNotSupportedAndClose(client net.Conn)
	SendConnectionNotAllowedAndClose(client net.Conn)
	SendFailAndClose(client net.Conn)
	SendSuccessWithTcpPort(client net.Conn) error
	SendSuccessWithUdpPort(client net.Conn) error
	SendSuccessWithParameters(addressType byte, address string, port uint16, client net.Conn) error
}

type BaseSender struct {
	protocol  v5.Protocol
	tcpConfig config.TcpConfig
	udpConfig config.UdpConfig
}

func NewBaseSender(
	protocol v5.Protocol,
	tcpConfig config.TcpConfig,
	udpConfig config.UdpConfig,
) (BaseSender, error) {
	return BaseSender{
		protocol:  protocol,
		tcpConfig: tcpConfig,
		udpConfig: udpConfig,
	}, nil
}

func (b BaseSender) SendConnectionRefusedAndClose(client net.Conn) {
	_ = b.protocol.ResponseWithConnectionRefused(1, "0.0.0.0", b.tcpConfig.GetBindPort(), client)
	_ = client.Close()
}

func (b BaseSender) SendNetworkUnreachableAndClose(client net.Conn) {
	_ = b.protocol.ResponseWithNetworkUnreachable(1, "0.0.0.0", b.tcpConfig.GetBindPort(), client)
	_ = client.Close()
}

func (b BaseSender) SendHostUnreachableAndClose(client net.Conn) {
	_ = b.protocol.ResponseWithHostUnreachable(1, "0.0.0.0", b.tcpConfig.GetBindPort(), client)
	_ = client.Close()
}

func (b BaseSender) SendFailAndClose(client net.Conn) {
	_ = b.protocol.ResponseWithFail(1, "0.0.0.0", b.tcpConfig.GetBindPort(), client)
	_ = client.Close()
}

func (b BaseSender) SendCommandNotSupportedAndClose(client net.Conn) {
	_ = b.protocol.ResponseWithCommandNotSupported(1, "0.0.0.0", b.tcpConfig.GetBindPort(), client)
	_ = client.Close()
}

func (b BaseSender) SendConnectionNotAllowedAndClose(client net.Conn) {
	_ = b.protocol.ResponseWithNotAllowed(1, "0.0.0.0", b.tcpConfig.GetBindPort(), client)
	_ = client.Close()
}

func (b BaseSender) SendAddressNotSupportedAndClose(client net.Conn) {
	_ = b.protocol.ResponseWithAddressNotSupported(1, "0.0.0.0", b.tcpConfig.GetBindPort(), client)
	_ = client.Close()
}

func (b BaseSender) SendSuccessWithTcpPort(client net.Conn) error {
	return b.protocol.ResponseWithSuccess(1, "0.0.0.0", b.tcpConfig.GetBindPort(), client)
}

func (b BaseSender) SendSuccessWithUdpPort(client net.Conn) error {
	return b.protocol.ResponseWithSuccess(1, "0.0.0.0", b.udpConfig.GetBindPort(), client)
}

func (b BaseSender) SendSuccessWithParameters(addressType byte, address string, port uint16, client net.Conn) error {
	return b.protocol.ResponseWithSuccess(addressType, address, port, client)
}
