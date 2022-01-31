package main

import (
	"github.com/sarulabs/di"
	"socks/server"
)

func main() {
	builder, err := di.NewBuilder()

	if err != nil {
		panic(err)
	}

	register(*builder)
}

func register(builder di.Builder) {
	registerConfig(builder)
	registerHandlers(builder)
	registerZeroLog(builder)
	registerLogger(builder)
	registerManagers(builder)
	registerProtocol(builder)
	registerServer(builder)
	registerTransfer(builder)
	registerUtils(builder)

	start(builder)
}

func start(builder di.Builder) {
	ctn := builder.Build()

	serv := ctn.Get("server").(server.Server)

	serv.Start()
}
