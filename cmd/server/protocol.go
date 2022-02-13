package main

import (
	"github.com/sarulabs/di"
	"socks/config/tcp"
	"socks/config/udp"
	"socks/protocol"
	"socks/protocol/auth/password"
	v4 "socks/protocol/v4"
	"socks/protocol/v4a"
	v5 "socks/protocol/v5"
	"socks/utils"
)

func registerProtocol(builder di.Builder) {
	receiverDef := di.Def{
		Name:  "receiver",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			buffer := ctn.Get("buffer_reader").(utils.BufferReader)

			return protocol.NewBaseReceiver(buffer)
		},
	}

	err := builder.Add(
		receiverDef,
	)

	if err != nil {
		panic(err)
	}

	registerAuth(builder)
	registerV4Protocol(builder)
	registerV4aProtocol(builder)
	registerV5Protocol(builder)
}

func registerAuth(builder di.Builder) {
	registerPasswordAuth(builder)
}

func registerPasswordAuth(builder di.Builder) {
	parserDef := di.Def{
		Name:  "auth_password_parser",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return password.NewBaseParser(), nil
		},
	}

	builderDef := di.Def{
		Name:  "auth_password_builder",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return password.NewBaseBuilder()
		},
	}

	receiverDef := di.Def{
		Name:  "auth_password_receiver",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			parser := ctn.Get("auth_password_parser").(password.Parser)
			buffer := ctn.Get("buffer_reader").(utils.BufferReader)

			return password.NewBaseReceiver(parser, buffer)
		},
	}

	senderDef := di.Def{
		Name:  "auth_password_sender",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			builder := ctn.Get("auth_password_builder").(password.Builder)

			return password.NewBaseSender(builder)
		},
	}

	err := builder.Add(
		parserDef,
		builderDef,
		receiverDef,
		senderDef,
	)

	if err != nil {
		panic(err)
	}
}

func registerV4Protocol(builder di.Builder) {
	parserDef := di.Def{
		Name:  "v4_parser",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return v4.NewBaseParser(), nil
		},
	}

	builderDef := di.Def{
		Name:  "v4_builder",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return v4.NewBaseBuilder(), nil
		},
	}

	senderDef := di.Def{
		Name:  "v4_sender",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			tcpConfig := ctn.Get("tcp_base_config").(tcp.Config).Bind
			builder := ctn.Get("v4_builder").(v4.Builder)

			return v4.NewBaseSender(
				tcpConfig,
				builder,
			)
		},
	}

	err := builder.Add(
		parserDef,
		builderDef,
		senderDef,
	)

	if err != nil {
		panic(err)
	}
}

func registerV4aProtocol(builder di.Builder) {
	parserDef := di.Def{
		Name:  "v4a_parser",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return v4a.NewBaseParser(), nil
		},
	}

	builderDef := di.Def{
		Name:  "v4a_builder",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return v4a.NewBaseBuilder(), nil
		},
	}

	senderDef := di.Def{
		Name:  "v4a_sender",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			tcpConfig := ctn.Get("tcp_base_config").(tcp.Config).Bind
			builder := ctn.Get("v4a_builder").(v4a.Builder)

			return v4a.NewBaseSender(
				tcpConfig,
				builder,
			)
		},
	}

	err := builder.Add(
		parserDef,
		builderDef,
		senderDef,
	)

	if err != nil {
		panic(err)
	}
}

func registerV5Protocol(builder di.Builder) {
	parserDef := di.Def{
		Name:  "v5_parser",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			addressUtils := ctn.Get("address_utils").(utils.AddressUtils)

			return v5.NewBaseParser(addressUtils), nil
		},
	}

	builderDef := di.Def{
		Name:  "v5_builder",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return v5.NewBaseBuilder()
		},
	}

	receiverDef := di.Def{
		Name:  "v5_receiver",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			parser := ctn.Get("v5_parser").(v5.Parser)
			buffer := ctn.Get("buffer_reader").(utils.BufferReader)

			return v5.NewBaseReceiver(parser, buffer)
		},
	}

	senderDef := di.Def{
		Name:  "v5_sender",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			tcpConfig := ctn.Get("tcp_base_config").(tcp.Config).Bind
			udpConfig := ctn.Get("udp_base_config").(udp.Config).Bind
			builder := ctn.Get("v5_builder").(v5.Builder)

			return v5.NewBaseSender(tcpConfig, udpConfig, builder)
		},
	}

	err := builder.Add(
		parserDef,
		builderDef,
		receiverDef,
		senderDef,
	)

	if err != nil {
		panic(err)
	}
}
