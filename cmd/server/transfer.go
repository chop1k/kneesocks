package main

import (
	"github.com/sarulabs/di"
	"socks/transfer"
)

func registerTransfer(builder di.Builder) {
	streamHandlerDef := di.Def{
		Name:  "stream_handler",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return transfer.NewBaseStreamHandler(), nil
		},
	}

	err := builder.Add(
		streamHandlerDef,
	)

	if err != nil {
		panic(err)
	}
}
