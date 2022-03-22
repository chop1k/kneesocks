package dependency

import (
	"socks/internal/kneesocks/dependency/build"

	"github.com/sarulabs/di"
)

func registerServer(builder di.Builder) {
	tcpServerDef := di.Def{
		Name:  "tcp_server",
		Scope: di.App,
		Build: build.TcpServer,
	}

	udpServerDef := di.Def{
		Name:  "udp_server",
		Scope: di.App,
		Build: build.UdpServer,
	}

	err := builder.Add(
		tcpServerDef,
		udpServerDef,
	)

	if err != nil {
		panic(err)
	}
}
