package build

import (
	"socks/internal/kneesocks/config/tcp"
	"socks/internal/kneesocks/config/udp"
	"socks/pkg/protocol"
	"socks/pkg/protocol/auth/password"
	v4 "socks/pkg/protocol/v4"
	"socks/pkg/protocol/v4a"
	v5 "socks/pkg/protocol/v5"
	"socks/pkg/utils"

	"github.com/sarulabs/di"
)

func Receiver(ctn di.Container) (interface{}, error) {
	buffer := ctn.Get("buffer_reader").(utils.BufferReader)

	return protocol.NewReceiver(buffer)
}

func AuthPasswordParser(ctn di.Container) (interface{}, error) {
	return password.NewParser(), nil
}

func AuthPasswordBuilder(ctn di.Container) (interface{}, error) {
	return password.NewBuilder()
}

func AuthPasswordReceiver(ctn di.Container) (interface{}, error) {
	parser := ctn.Get("auth_password_parser").(password.Parser)
	buffer := ctn.Get("buffer_reader").(utils.BufferReader)

	return password.NewReceiver(parser, buffer)
}

func AuthPasswordSender(ctn di.Container) (interface{}, error) {
	builder := ctn.Get("auth_password_builder").(password.Builder)

	return password.NewSender(builder)
}

func V4Parser(ctn di.Container) (interface{}, error) {
	return v4.NewParser(), nil
}

func V4Builder(ctn di.Container) (interface{}, error) {
	return v4.NewBuilder(), nil
}

func V4Sender(ctn di.Container) (interface{}, error) {
	bind := ctn.Get("tcp_base_config").(tcp.Config).Bind
	builder := ctn.Get("v4_builder").(v4.Builder)

	return v4.NewSender(
		bind,
		builder,
	)
}

func V4aParser(ctn di.Container) (interface{}, error) {
	return v4a.NewParser(), nil
}

func V4aBuilder(ctn di.Container) (interface{}, error) {
	return v4a.NewBuilder(), nil
}

func V4aSender(ctn di.Container) (interface{}, error) {
	bind := ctn.Get("tcp_base_config").(tcp.Config).Bind
	builder := ctn.Get("v4a_builder").(v4a.Builder)

	return v4a.NewSender(
		bind,
		builder,
	)
}

func V5Parser(ctn di.Container) (interface{}, error) {
	addressUtils := ctn.Get("address_utils").(utils.AddressUtils)

	return v5.NewParser(addressUtils), nil
}

func V5Builder(ctn di.Container) (interface{}, error) {
	return v5.NewBuilder()
}

func V5Receiver(ctn di.Container) (interface{}, error) {
	parser := ctn.Get("v5_parser").(v5.Parser)
	buffer := ctn.Get("buffer_reader").(utils.BufferReader)

	return v5.NewReceiver(parser, buffer)
}

func V5Sender(ctn di.Container) (interface{}, error) {
	tcpConfig := ctn.Get("tcp_base_config").(tcp.Config).Bind

	_udpConfig := ctn.Get("udp_base_config")

	var udpConfig *udp.BindConfig

	if _udpConfig == nil {
		udpConfig = nil
	} else {
		udpConfig = &_udpConfig.(*udp.Config).Bind
	}

	builder := ctn.Get("v5_builder").(v5.Builder)

	return v5.NewSender(tcpConfig, udpConfig, builder)
}
