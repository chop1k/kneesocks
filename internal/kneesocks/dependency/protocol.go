package dependency

import (
	"socks/internal/kneesocks/dependency/build"

	"github.com/sarulabs/di"
)

func registerProtocol(builder di.Builder) {
	receiverDef := di.Def{
		Name:  "receiver",
		Scope: di.App,
		Build: build.Receiver,
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
		Build: build.AuthPasswordParser,
	}

	builderDef := di.Def{
		Name:  "auth_password_builder",
		Scope: di.App,
		Build: build.AuthPasswordBuilder,
	}

	receiverDef := di.Def{
		Name:  "auth_password_receiver",
		Scope: di.App,
		Build: build.AuthPasswordReceiver,
	}

	senderDef := di.Def{
		Name:  "auth_password_sender",
		Scope: di.App,
		Build: build.AuthPasswordSender,
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
		Build: build.V4Parser,
	}

	builderDef := di.Def{
		Name:  "v4_builder",
		Scope: di.App,
		Build: build.V4Builder,
	}

	senderDef := di.Def{
		Name:  "v4_sender",
		Scope: di.App,
		Build: build.V4Sender,
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
		Build: build.V4aParser,
	}

	builderDef := di.Def{
		Name:  "v4a_builder",
		Scope: di.App,
		Build: build.V4aBuilder,
	}

	senderDef := di.Def{
		Name:  "v4a_sender",
		Scope: di.App,
		Build: build.V4aSender,
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
		Build: build.V5Parser,
	}

	builderDef := di.Def{
		Name:  "v5_builder",
		Scope: di.App,
		Build: build.V5Builder,
	}

	receiverDef := di.Def{
		Name:  "v5_receiver",
		Scope: di.App,
		Build: build.V5Receiver,
	}

	senderDef := di.Def{
		Name:  "v5_sender",
		Scope: di.App,
		Build: build.V5Sender,
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
