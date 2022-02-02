package main

import (
	"github.com/sarulabs/di"
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

	protocolDef := di.Def{
		Name:  "auth_password",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			parser := ctn.Get("auth_password_parser").(password.Parser)
			builder := ctn.Get("auth_password_builder").(password.Builder)

			return password.NewPassword(parser, builder), nil
		},
	}

	err := builder.Add(
		parserDef,
		builderDef,
		protocolDef,
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

	protocolDef := di.Def{
		Name:  "v4a",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			builder := ctn.Get("v4a_builder").(v4a.Builder)

			return v4a.NewProtocol(builder), nil
		},
	}

	err := builder.Add(
		parserDef,
		builderDef,
		protocolDef,
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

	protocolDef := di.Def{
		Name:  "v5",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			builder := ctn.Get("v5_builder").(v5.Builder)
			parser := ctn.Get("v5_parser").(v5.Parser)

			return v5.NewProtocol(builder, parser), nil
		},
	}

	err := builder.Add(
		parserDef,
		builderDef,
		protocolDef,
	)

	if err != nil {
		panic(err)
	}
}
