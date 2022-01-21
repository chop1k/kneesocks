package main

import "github.com/sarulabs/di"

func main() {
	builder, err := di.NewBuilder()

	if err != nil {
		panic(err)
	}

	register(*builder)
}

func register(builder di.Builder) {

}

func start(ctn di.Container) {
	server := ctn.Get("server").(Server)

	server.Start()
}
