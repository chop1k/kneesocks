package kneesocks

import (
	"github.com/sarulabs/di"
	"socks/internal/kneesocks/dependency"
	"socks/internal/kneesocks/server"
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

	go ctn.Get("udp_server").(server.UdpServer).Listen()

	ctn.Get("tcp_server").(server.TcpServer).Listen()
}
