package kneesocks

import (
	"socks/internal/kneesocks/dependency"
	"socks/internal/kneesocks/server"

	"github.com/sarulabs/di"
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

	_udp := ctn.Get("udp_server")

	if _udp != nil {
		udp := _udp.(server.UdpServer)

		go udp.Listen()
	}

	ctn.Get("tcp_server").(server.TcpServer).Listen()
}
