package v5

import (
	"net"
	"socks/config/tcp"
	"socks/config/udp"
	v52 "socks/config/v5"
	"time"
)

type Sender interface {
	SendMethodSelection(config v52.Config, method byte, client net.Conn) error
	SendConnectionRefusedAndClose(config v52.Config, client net.Conn)
	SendTTLExpiredAndClose(config v52.Config, client net.Conn)
	SendNetworkUnreachableAndClose(config v52.Config, client net.Conn)
	SendHostUnreachableAndClose(config v52.Config, client net.Conn)
	SendCommandNotSupportedAndClose(config v52.Config, client net.Conn)
	SendAddressNotSupportedAndClose(config v52.Config, client net.Conn)
	SendConnectionNotAllowedAndClose(config v52.Config, client net.Conn)
	SendFailAndClose(config v52.Config, client net.Conn)
	SendSuccessWithTcpPort(config v52.Config, client net.Conn) error
	SendSuccessWithUdpPort(config v52.Config, client net.Conn) error
	SendSuccessWithParameters(config v52.Config, addressType byte, address string, port uint16, client net.Conn) error
}

type BaseSender struct {
	tcpConfig tcp.BindConfig
	udpConfig udp.BindConfig
	builder   Builder
}

func NewBaseSender(
	tcpConfig tcp.BindConfig,
	udpConfig udp.BindConfig,
	builder Builder,
) (BaseSender, error) {
	return BaseSender{
		tcpConfig: tcpConfig,
		udpConfig: udpConfig,
		builder:   builder,
	}, nil
}

func (b BaseSender) SendMethodSelection(config v52.Config, method byte, client net.Conn) error {
	deadlineErr := client.SetWriteDeadline(time.Now().Add(config.Deadline.Selection))

	if deadlineErr != nil {
		return deadlineErr
	}

	selection := MethodSelectionChunk{
		SocksVersion: 5,
		Method:       method,
	}

	response, err := b.builder.BuildMethodSelection(selection)

	if err != nil {
		return err
	}

	_, err = client.Write(response)

	return err
}

func (b BaseSender) responseWithCode(config v52.Config, code byte, addrType byte, addr string, port uint16, client net.Conn) error {
	deadlineErr := client.SetWriteDeadline(time.Now().Add(config.Deadline.Response))

	if deadlineErr != nil {
		return deadlineErr
	}

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

	_, err = client.Write(response)

	return err
}

func (b BaseSender) responseWithSuccess(config v52.Config, addrType byte, addr string, port uint16, client net.Conn) error {
	return b.responseWithCode(config, 0, addrType, addr, port, client)
}

func (b BaseSender) responseWithFail(config v52.Config, addrType byte, addr string, port uint16, client net.Conn) error {
	return b.responseWithCode(config, 1, addrType, addr, port, client)
}

func (b BaseSender) responseWithNotAllowed(config v52.Config, addrType byte, addr string, port uint16, client net.Conn) error {
	return b.responseWithCode(config, 2, addrType, addr, port, client)
}

func (b BaseSender) responseWithNetworkUnreachable(config v52.Config, addrType byte, addr string, port uint16, client net.Conn) error {
	return b.responseWithCode(config, 3, addrType, addr, port, client)
}

func (b BaseSender) responseWithHostUnreachable(config v52.Config, addrType byte, addr string, port uint16, client net.Conn) error {
	return b.responseWithCode(config, 4, addrType, addr, port, client)
}

func (b BaseSender) responseWithConnectionRefused(config v52.Config, addrType byte, addr string, port uint16, client net.Conn) error {
	return b.responseWithCode(config, 5, addrType, addr, port, client)
}

func (b BaseSender) responseWithTTLExpired(config v52.Config, addrType byte, addr string, port uint16, client net.Conn) error {
	return b.responseWithCode(config, 6, addrType, addr, port, client)
}

func (b BaseSender) responseWithCommandNotSupported(config v52.Config, addrType byte, addr string, port uint16, client net.Conn) error {
	return b.responseWithCode(config, 7, addrType, addr, port, client)
}

func (b BaseSender) responseWithAddressNotSupported(config v52.Config, addrType byte, addr string, port uint16, client net.Conn) error {
	return b.responseWithCode(config, 8, addrType, addr, port, client)
}

func (b BaseSender) SendConnectionRefusedAndClose(config v52.Config, client net.Conn) {
	_ = b.responseWithConnectionRefused(config, 1, "0.0.0.0", b.tcpConfig.Port, client)
	_ = client.Close()
}

func (b BaseSender) SendTTLExpiredAndClose(config v52.Config, client net.Conn) {
	_ = b.responseWithTTLExpired(config, 1, "0.0.0.0", b.tcpConfig.Port, client)
	_ = client.Close()
}

func (b BaseSender) SendNetworkUnreachableAndClose(config v52.Config, client net.Conn) {
	_ = b.responseWithNetworkUnreachable(config, 1, "0.0.0.0", b.tcpConfig.Port, client)
	_ = client.Close()
}

func (b BaseSender) SendHostUnreachableAndClose(config v52.Config, client net.Conn) {
	_ = b.responseWithHostUnreachable(config, 1, "0.0.0.0", b.tcpConfig.Port, client)
	_ = client.Close()
}

func (b BaseSender) SendFailAndClose(config v52.Config, client net.Conn) {
	_ = b.responseWithFail(config, 1, "0.0.0.0", b.tcpConfig.Port, client)
	_ = client.Close()
}

func (b BaseSender) SendCommandNotSupportedAndClose(config v52.Config, client net.Conn) {
	_ = b.responseWithCommandNotSupported(config, 1, "0.0.0.0", b.tcpConfig.Port, client)
	_ = client.Close()
}

func (b BaseSender) SendConnectionNotAllowedAndClose(config v52.Config, client net.Conn) {
	_ = b.responseWithNotAllowed(config, 1, "0.0.0.0", b.tcpConfig.Port, client)
	_ = client.Close()
}

func (b BaseSender) SendAddressNotSupportedAndClose(config v52.Config, client net.Conn) {
	_ = b.responseWithAddressNotSupported(config, 1, "0.0.0.0", b.tcpConfig.Port, client)
	_ = client.Close()
}

func (b BaseSender) SendSuccessWithTcpPort(config v52.Config, client net.Conn) error {
	return b.responseWithSuccess(config, 1, "0.0.0.0", b.tcpConfig.Port, client)
}

func (b BaseSender) SendSuccessWithUdpPort(config v52.Config, client net.Conn) error {
	return b.responseWithSuccess(config, 1, "0.0.0.0", b.udpConfig.Port, client)
}

func (b BaseSender) SendSuccessWithParameters(config v52.Config, addressType byte, address string, port uint16, client net.Conn) error {
	return b.responseWithSuccess(config, addressType, address, port, client)
}
