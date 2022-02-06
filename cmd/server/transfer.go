package main

import (
	"github.com/sarulabs/di"
	"socks/managers"
	"socks/transfer"
)

func registerTransfer(builder di.Builder) {
	bindHandlerDef := di.Def{
		Name:  "transfer_bind_handler",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			bindRate := ctn.Get("bind_rate_manager").(managers.BindRateManager)
			handler := ctn.Get("transfer_handler").(transfer.BaseHandler)

			return transfer.NewBaseBindHandler(bindRate, handler)
		},
	}

	connectHandlerDef := di.Def{
		Name:  "transfer_connect_handler",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			handler := ctn.Get("transfer_handler").(transfer.BaseHandler)

			return transfer.NewBaseConnectHandler(handler)
		},
	}

	transferHandlerDef := di.Def{
		Name:  "transfer_handler",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return transfer.NewBaseHandler()
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
