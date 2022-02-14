package dependency

import (
	"github.com/sarulabs/di"
	"socks/internal/kneesocks/managers"
	"socks/internal/kneesocks/transfer"
)

func registerTransfer(builder di.Builder) {
	bindHandlerDef := di.Def{
		Name:  "transfer_bind_handler",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			bindRate := ctn.Get("bind_rate_manager").(managers.BindRateManager)
			handler := ctn.Get("transfer_handler").(transfer.Handler)

			return transfer.NewBindHandler(bindRate, handler)
		},
	}

	connectHandlerDef := di.Def{
		Name:  "transfer_connect_handler",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			handler := ctn.Get("transfer_handler").(transfer.Handler)

			return transfer.NewConnectHandler(handler)
		},
	}

	transferHandlerDef := di.Def{
		Name:  "transfer_handler",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return transfer.NewHandler()
		},
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
