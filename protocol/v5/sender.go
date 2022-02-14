package v5

import (
	"net"
	"socks/config/tcp"
	"socks/config/udp"
	v52 "socks/config/v5"
	"time"
)

type Sender struct {
	tcpConfig tcp.BindConfig
	udpConfig udp.BindConfig
	builder   Builder
}

func NewSender(
	tcpConfig tcp.BindConfig,
	udpConfig udp.BindConfig,
	builder Builder,
) (Sender, error) {
	return Sender{
		tcpConfig: tcpConfig,
		udpConfig: udpConfig,
		builder:   builder,
	}, nil
}

func (b Sender) SendMethodSelection(config v52.Config, method byte, client net.Conn) error {
	if config.Deadline.Selection > 0 {
		deadlineErr := client.SetWriteDeadline(time.Now().Add(config.Deadline.Selection))

		if deadlineErr != nil {
			return deadlineErr
		}
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

func (b Sender) responseWithCode(config v52.Config, code byte, addrType byte, addr string, port uint16, client net.Conn) error {
	if config.Deadline.Response > 0 {
		deadlineErr := client.SetWriteDeadline(time.Now().Add(config.Deadline.Response))

		if deadlineErr != nil {
			return deadlineErr
		}
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

func (b Sender) responseWithSuccess(config v52.Config, addrType byte, addr string, port uint16, client net.Conn) error {
	return b.responseWithCode(config, 0, addrType, addr, port, client)
}

func (b Sender) responseWithFail(config v52.Config, addrType byte, addr string, port uint16, client net.Conn) error {
	return b.responseWithCode(config, 1, addrType, addr, port, client)
}

func (b Sender) responseWithNotAllowed(config v52.Config, addrType byte, addr string, port uint16, client net.Conn) error {
	return b.responseWithCode(config, 2, addrType, addr, port, client)
}

func (b Sender) responseWithNetworkUnreachable(config v52.Config, addrType byte, addr string, port uint16, client net.Conn) error {
	return b.responseWithCode(config, 3, addrType, addr, port, client)
}

func (b Sender) responseWithHostUnreachable(config v52.Config, addrType byte, addr string, port uint16, client net.Conn) error {
	return b.responseWithCode(config, 4, addrType, addr, port, client)
}

func (b Sender) responseWithConnectionRefused(config v52.Config, addrType byte, addr string, port uint16, client net.Conn) error {
	return b.responseWithCode(config, 5, addrType, addr, port, client)
}

func (b Sender) responseWithTTLExpired(config v52.Config, addrType byte, addr string, port uint16, client net.Conn) error {
	return b.responseWithCode(config, 6, addrType, addr, port, client)
}

func (b Sender) responseWithCommandNotSupported(config v52.Config, addrType byte, addr string, port uint16, client net.Conn) error {
	return b.responseWithCode(config, 7, addrType, addr, port, client)
}

func (b Sender) responseWithAddressNotSupported(config v52.Config, addrType byte, addr string, port uint16, client net.Conn) error {
	return b.responseWithCode(config, 8, addrType, addr, port, client)
}

func (b Sender) SendConnectionRefusedAndClose(config v52.Config, client net.Conn) {
	_ = b.responseWithConnectionRefused(config, 1, "0.0.0.0", b.tcpConfig.Port, client)
	_ = client.Close()
}

func (b Sender) SendTTLExpiredAndClose(config v52.Config, client net.Conn) {
	_ = b.responseWithTTLExpired(config, 1, "0.0.0.0", b.tcpConfig.Port, client)
	_ = client.Close()
}

func (b Sender) SendNetworkUnreachableAndClose(config v52.Config, client net.Conn) {
	_ = b.responseWithNetworkUnreachable(config, 1, "0.0.0.0", b.tcpConfig.Port, client)
	_ = client.Close()
}

func (b Sender) SendHostUnreachableAndClose(config v52.Config, client net.Conn) {
	_ = b.responseWithHostUnreachable(config, 1, "0.0.0.0", b.tcpConfig.Port, client)
	_ = client.Close()
}

func (b Sender) SendFailAndClose(config v52.Config, client net.Conn) {
	_ = b.responseWithFail(config, 1, "0.0.0.0", b.tcpConfig.Port, client)
	_ = client.Close()
}

func (b Sender) SendCommandNotSupportedAndClose(config v52.Config, client net.Conn) {
	_ = b.responseWithCommandNotSupported(config, 1, "0.0.0.0", b.tcpConfig.Port, client)
	_ = client.Close()
}

func (b Sender) SendConnectionNotAllowedAndClose(config v52.Config, client net.Conn) {
	_ = b.responseWithNotAllowed(config, 1, "0.0.0.0", b.tcpConfig.Port, client)
	_ = client.Close()
}

func (b Sender) SendAddressNotSupportedAndClose(config v52.Config, client net.Conn) {
	_ = b.responseWithAddressNotSupported(config, 1, "0.0.0.0", b.tcpConfig.Port, client)
	_ = client.Close()
}

func (b Sender) SendSuccessWithTcpPort(config v52.Config, client net.Conn) error {
	return b.responseWithSuccess(config, 1, "0.0.0.0", b.tcpConfig.Port, client)
}

func (b Sender) SendSuccessWithUdpPort(config v52.Config, client net.Conn) error {
	return b.responseWithSuccess(config, 1, "0.0.0.0", b.udpConfig.Port, client)
}

func (b Sender) SendSuccessWithParameters(config v52.Config, addressType byte, address string, port uint16, client net.Conn) error {
	return b.responseWithSuccess(config, addressType, address, port, client)
}
