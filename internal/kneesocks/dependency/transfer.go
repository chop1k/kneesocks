package dependency

import (
	"socks/internal/kneesocks/dependency/build"

	"github.com/sarulabs/di"
)

func registerTransfer(builder di.Builder) {
	bindHandlerDef := di.Def{
		Name:  "transfer_bind_handler",
		Scope: di.App,
		Build: build.TransferBindHandler,
	}

	connectHandlerDef := di.Def{
		Name:  "transfer_connect_handler",
		Scope: di.App,
		Build: build.TransferConnectHandler,
	}

	transferHandlerDef := di.Def{
		Name:  "transfer_handler",
		Scope: di.App,
		Build: build.TransferHandler,
	}

	err := builder.Add(
		bindHandlerDef,
		connectHandlerDef,
		transferHandlerDef,
	)

	if err != nil {
		panic(err)
	}
}
