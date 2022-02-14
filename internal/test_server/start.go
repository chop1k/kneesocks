package test_server

import (
	"github.com/sarulabs/di"
	"socks/internal/test_server/dependency"
	"socks/internal/test_server/server"
)

func Start() {
	builder, err := di.NewBuilder()

	if err != nil {
		panic(err)
	}

	dependency.Register(*builder)

	start(*builder)
}

func start(builder di.Builder) {
	ctn := builder.Build()

	ctn.Get("server").(server.Server).Start()
}
