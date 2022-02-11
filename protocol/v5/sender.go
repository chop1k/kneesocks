package v5

import (
	"net"
	"socks/config/tcp"
	"socks/config/udp"
	v52 "socks/config/v5"
	"time"
)

type Sender interface {
	SendMethodSelection(method byte, client net.Conn) error
	SendConnectionRefusedAndClose(client net.Conn)
	SendTTLExpiredAndClose(client net.Conn)
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
	tcpConfig tcp.BindConfig
	udpConfig udp.BindConfig
	config    v52.DeadlineConfig
	builder   Builder
}

func NewBaseSender(
	tcpConfig tcp.BindConfig,
	udpConfig udp.BindConfig,
	config v52.DeadlineConfig,
	builder Builder,
) (BaseSender, error) {
	return BaseSender{
		tcpConfig: tcpConfig,
		udpConfig: udpConfig,
		config:    config,
		builder:   builder,
	}, nil
}

func (b BaseSender) SendMethodSelection(method byte, client net.Conn) error {
	deadline, configErr := b.config.GetSelectionDeadline()

	if configErr != nil {
		return configErr
	}

	deadlineErr := client.SetWriteDeadline(time.Now().Add(deadline))

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

func (b BaseSender) responseWithCode(code byte, addrType byte, addr string, port uint16, client net.Conn) error {
	deadline, configErr := b.config.GetResponseDeadline()

	if configErr != nil {
		return configErr
	}

	deadlineErr := client.SetWriteDeadline(time.Now().Add(deadline))

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

func (b BaseSender) responseWithTTLExpired(addrType byte, addr string, port uint16, client net.Conn) error {
	return b.responseWithCode(6, addrType, addr, port, client)
}

func (b BaseSender) responseWithCommandNotSupported(addrType byte, addr string, port uint16, client net.Conn) error {
	return b.responseWithCode(7, addrType, addr, port, client)
}

func (b BaseSender) responseWithAddressNotSupported(addrType byte, addr string, port uint16, client net.Conn) error {
	return b.responseWithCode(8, addrType, addr, port, client)
}

func (b BaseSender) SendConnectionRefusedAndClose(client net.Conn) {
	port, err := b.tcpConfig.GetPort()

	if err != nil {
		panic(err)
	}

	_ = b.responseWithConnectionRefused(1, "0.0.0.0", port, client)
	_ = client.Close()
}

func (b BaseSender) SendTTLExpiredAndClose(client net.Conn) {
	port, err := b.tcpConfig.GetPort()

	if err != nil {
		panic(err)
	}

	_ = b.responseWithTTLExpired(1, "0.0.0.0", port, client)
	_ = client.Close()
}

func (b BaseSender) SendNetworkUnreachableAndClose(client net.Conn) {
	port, err := b.tcpConfig.GetPort()

	if err != nil {
		panic(err)
	}

	_ = b.responseWithNetworkUnreachable(1, "0.0.0.0", port, client)
	_ = client.Close()
}

func (b BaseSender) SendHostUnreachableAndClose(client net.Conn) {
	port, err := b.tcpConfig.GetPort()

	if err != nil {
		panic(err)
	}

	_ = b.responseWithHostUnreachable(1, "0.0.0.0", port, client)
	_ = client.Close()
}

func (b BaseSender) SendFailAndClose(client net.Conn) {
	port, err := b.tcpConfig.GetPort()

	if err != nil {
		panic(err)
	}

	_ = b.responseWithFail(1, "0.0.0.0", port, client)
	_ = client.Close()
}

func (b BaseSender) SendCommandNotSupportedAndClose(client net.Conn) {
	port, err := b.tcpConfig.GetPort()

	if err != nil {
		panic(err)
	}

	_ = b.responseWithCommandNotSupported(1, "0.0.0.0", port, client)
	_ = client.Close()
}

func (b BaseSender) SendConnectionNotAllowedAndClose(client net.Conn) {
	port, err := b.tcpConfig.GetPort()

	if err != nil {
		panic(err)
	}

	_ = b.responseWithNotAllowed(1, "0.0.0.0", port, client)
	_ = client.Close()
}

func (b BaseSender) SendAddressNotSupportedAndClose(client net.Conn) {
	port, err := b.tcpConfig.GetPort()

	if err != nil {
		panic(err)
	}

	_ = b.responseWithAddressNotSupported(1, "0.0.0.0", port, client)
	_ = client.Close()
}

func (b BaseSender) SendSuccessWithTcpPort(client net.Conn) error {
	port, err := b.tcpConfig.GetPort()

	if err != nil {
		panic(err)
	}

	return b.responseWithSuccess(1, "0.0.0.0", port, client)
}

func (b BaseSender) SendSuccessWithUdpPort(client net.Conn) error {
	port, err := b.udpConfig.GetPort()

	if err != nil {
		panic(err)
	}

	return b.responseWithSuccess(1, "0.0.0.0", port, client)
}

func (b BaseSender) SendSuccessWithParameters(addressType byte, address string, port uint16, client net.Conn) error {
	return b.responseWithSuccess(addressType, address, port, client)
}
