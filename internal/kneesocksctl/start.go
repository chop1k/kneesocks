package kneesocksctl

import (
	"github.com/sarulabs/di"
	"github.com/urfave/cli"
	"os"
	"socks/internal/kneesocksctl/dependency"
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

	app := ctn.Get("cli_app").(cli.App)

	err := app.Run(os.Args)

	if err != nil {
		panic(err)
	}
}
