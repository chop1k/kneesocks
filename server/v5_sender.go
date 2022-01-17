package server

import (
	"net"
	"socks/config"
	v5 "socks/protocol/v5"
)

type V5Sender interface {
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

type BaseV5Sender struct {
	protocol  v5.Protocol
	tcpConfig config.TcpConfig
	udpConfig config.UdpConfig
}

func NewBaseV5Sender(
	protocol v5.Protocol,
	tcpConfig config.TcpConfig,
	udpConfig config.UdpConfig,
) (BaseV5Sender, error) {
	return BaseV5Sender{
		protocol:  protocol,
		tcpConfig: tcpConfig,
		udpConfig: udpConfig,
	}, nil
}

func (b BaseV5Sender) SendConnectionRefusedAndClose(client net.Conn) {
	_ = b.protocol.ResponseWithConnectionRefused(1, "0.0.0.0", uint16(b.tcpConfig.GetBindPort()), client)
	_ = client.Close()
}

func (b BaseV5Sender) SendNetworkUnreachableAndClose(client net.Conn) {
	_ = b.protocol.ResponseWithNetworkUnreachable(1, "0.0.0.0", uint16(b.tcpConfig.GetBindPort()), client)
	_ = client.Close()
}

func (b BaseV5Sender) SendHostUnreachableAndClose(client net.Conn) {
	_ = b.protocol.ResponseWithHostUnreachable(1, "0.0.0.0", uint16(b.tcpConfig.GetBindPort()), client)
	_ = client.Close()
}

func (b BaseV5Sender) SendFailAndClose(client net.Conn) {
	_ = b.protocol.ResponseWithFail(1, "0.0.0.0", uint16(b.tcpConfig.GetBindPort()), client)
	_ = client.Close()
}

func (b BaseV5Sender) SendCommandNotSupportedAndClose(client net.Conn) {
	_ = b.protocol.ResponseWithCommandNotSupported(1, "0.0.0.0", uint16(b.tcpConfig.GetBindPort()), client)
	_ = client.Close()
}

func (b BaseV5Sender) SendConnectionNotAllowedAndClose(client net.Conn) {
	_ = b.protocol.ResponseWithNotAllowed(1, "0.0.0.0", uint16(b.tcpConfig.GetBindPort()), client)
	_ = client.Close()
}

func (b BaseV5Sender) SendAddressNotSupportedAndClose(client net.Conn) {
	_ = b.protocol.ResponseWithAddressNotSupported(1, "0.0.0.0", uint16(b.tcpConfig.GetBindPort()), client)
	_ = client.Close()
}

func (b BaseV5Sender) SendSuccessWithTcpPort(client net.Conn) error {
	return b.protocol.ResponseWithSuccess(1, "0.0.0.0", uint16(b.tcpConfig.GetBindPort()), client)
}

func (b BaseV5Sender) SendSuccessWithUdpPort(client net.Conn) error {
	return b.protocol.ResponseWithSuccess(1, "0.0.0.0", uint16(b.udpConfig.GetBindPort()), client)
}

func (b BaseV5Sender) SendSuccessWithParameters(addressType byte, address string, port uint16, client net.Conn) error {
	return b.protocol.ResponseWithSuccess(addressType, address, port, client)
}
