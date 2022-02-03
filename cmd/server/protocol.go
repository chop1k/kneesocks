package main

import (
	"github.com/sarulabs/di"
	"socks/config/tcp"
	"socks/config/udp"
	v53 "socks/config/v5"
	"socks/managers"
	"socks/protocol/auth/password"
	v4 "socks/protocol/v4"
	"socks/protocol/v4a"
	v5 "socks/protocol/v5"
	"socks/utils"
)

func registerProtocol(builder di.Builder) {
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
			cfg := ctn.Get("v5_deadline_config").(v53.DeadlineConfig)
			deadlineManager := ctn.Get("deadline_manager").(managers.DeadlineManager)
			parser := ctn.Get("auth_password_parser").(password.Parser)

			return password.NewBaseReceiver(cfg, deadlineManager, parser)
		},
	}

	senderDef := di.Def{
		Name:  "auth_password_sender",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("v5_deadline_config").(v53.DeadlineConfig)
			deadlineManager := ctn.Get("deadline_manager").(managers.DeadlineManager)
			builder := ctn.Get("auth_password_builder").(password.Builder)

			return password.NewBaseSender(cfg, deadlineManager, builder)
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

	err := builder.Add(
		parserDef,
		builderDef,
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

	err := builder.Add(
		parserDef,
		builderDef,
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
			cfg := ctn.Get("v5_deadline_config").(v53.DeadlineConfig)
			deadline := ctn.Get("deadline_manager").(managers.DeadlineManager)
			parser := ctn.Get("v5_parser").(v5.Parser)

			return v5.NewBaseReceiver(cfg, deadline, parser)
		},
	}

	senderDef := di.Def{
		Name:  "v5_sender",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("v5_deadline_config").(v53.DeadlineConfig)
			tcpConfig := ctn.Get("tcp_config").(tcp.Config)
			udpConfig := ctn.Get("udp_config").(udp.Config)
			deadline := ctn.Get("deadline_manager").(managers.DeadlineManager)
			builder := ctn.Get("v5_builder").(v5.Builder)

			return v5.NewBaseSender(tcpConfig, udpConfig, cfg, deadline, builder)
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
