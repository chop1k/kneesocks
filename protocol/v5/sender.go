package v5

import (
	"net"
	"socks/config/tcp"
	"socks/config/udp"
	v52 "socks/config/v5"
	"socks/managers"
)

type Sender interface {
	SendMethodSelection(method byte, client net.Conn) error
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
	tcpConfig tcp.Config
	udpConfig udp.Config
	config    v52.DeadlineConfig
	deadline  managers.DeadlineManager
	builder   Builder
}

func NewBaseSender(
	tcpConfig tcp.Config,
	udpConfig udp.Config,
	config v52.DeadlineConfig,
	deadline managers.DeadlineManager,
	builder Builder,
) (BaseSender, error) {
	return BaseSender{
		tcpConfig: tcpConfig,
		udpConfig: udpConfig,
		config:    config,
		deadline:  deadline,
		builder:   builder,
	}, nil
}

func (b BaseSender) SendMethodSelection(method byte, client net.Conn) error {
	selection := MethodSelectionChunk{
		SocksVersion: 5,
		Method:       method,
	}

	response, err := b.builder.BuildMethodSelection(selection)

	if err != nil {
		return err
	}

	return b.deadline.Write(b.config.GetSelectionDeadline(), response, client)
}

func (b BaseSender) responseWithCode(code byte, addrType byte, addr string, port uint16, client net.Conn) error {
	chunk := ResponseChunk{
		SocksVersion: 5,
		ReplyCode:    code,
		AddressType:  addrType,
		Address:      addr,
		Port:         port,
	}

	response, err := b.builder.BuildResponse(chunk)

	if err != nil {
		return err
	}

	err = b.deadline.Write(b.config.GetResponseDeadline(), response, client)

	if err != nil {
		return err
	}

	return nil
}

func (b BaseSender) responseWithSuccess(addrType byte, addr string, port uint16, client net.Conn) error {
	return b.responseWithCode(0, addrType, addr, port, client)
}

func (b BaseSender) responseWithFail(addrType byte, addr string, port uint16, client net.Conn) error {
	return b.responseWithCode(1, addrType, addr, port, client)
}

func (b BaseSender) responseWithNotAllowed(addrType byte, addr string, port uint16, client net.Conn) error {
	return b.responseWithCode(2, addrType, addr, port, client)
}

func (b BaseSender) responseWithNetworkUnreachable(addrType byte, addr string, port uint16, client net.Conn) error {
	return b.responseWithCode(3, addrType, addr, port, client)
}

func (b BaseSender) responseWithHostUnreachable(addrType byte, addr string, port uint16, client net.Conn) error {
	return b.responseWithCode(4, addrType, addr, port, client)
}

func (b BaseSender) responseWithConnectionRefused(addrType byte, addr string, port uint16, client net.Conn) error {
	return b.responseWithCode(5, addrType, addr, port, client)
}

func (b BaseSender) responseWithCommandNotSupported(addrType byte, addr string, port uint16, client net.Conn) error {
	return b.responseWithCode(7, addrType, addr, port, client)
}

func (b BaseSender) responseWithAddressNotSupported(addrType byte, addr string, port uint16, client net.Conn) error {
	return b.responseWithCode(8, addrType, addr, port, client)
}

func (b BaseSender) SendConnectionRefusedAndClose(client net.Conn) {
	_ = b.responseWithConnectionRefused(1, "0.0.0.0", b.tcpConfig.GetBindPort(), client)
	_ = client.Close()
}

func (b BaseSender) SendNetworkUnreachableAndClose(client net.Conn) {
	_ = b.responseWithNetworkUnreachable(1, "0.0.0.0", b.tcpConfig.GetBindPort(), client)
	_ = client.Close()
}

func (b BaseSender) SendHostUnreachableAndClose(client net.Conn) {
	_ = b.responseWithHostUnreachable(1, "0.0.0.0", b.tcpConfig.GetBindPort(), client)
	_ = client.Close()
}

func (b BaseSender) SendFailAndClose(client net.Conn) {
	_ = b.responseWithFail(1, "0.0.0.0", b.tcpConfig.GetBindPort(), client)
	_ = client.Close()
}

func (b BaseSender) SendCommandNotSupportedAndClose(client net.Conn) {
	_ = b.responseWithCommandNotSupported(1, "0.0.0.0", b.tcpConfig.GetBindPort(), client)
	_ = client.Close()
}

func (b BaseSender) SendConnectionNotAllowedAndClose(client net.Conn) {
	_ = b.responseWithNotAllowed(1, "0.0.0.0", b.tcpConfig.GetBindPort(), client)
	_ = client.Close()
}

func (b BaseSender) SendAddressNotSupportedAndClose(client net.Conn) {
	_ = b.responseWithAddressNotSupported(1, "0.0.0.0", b.tcpConfig.GetBindPort(), client)
	_ = client.Close()
}

func (b BaseSender) SendSuccessWithTcpPort(client net.Conn) error {
	return b.responseWithSuccess(1, "0.0.0.0", b.tcpConfig.GetBindPort(), client)
}

func (b BaseSender) SendSuccessWithUdpPort(client net.Conn) error {
	return b.responseWithSuccess(1, "0.0.0.0", b.udpConfig.GetBindPort(), client)
}

func (b BaseSender) SendSuccessWithParameters(addressType byte, address string, port uint16, client net.Conn) error {
	return b.responseWithSuccess(addressType, address, port, client)
}
