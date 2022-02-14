package main

import (
	"github.com/sarulabs/di"
	"socks/config/tcp"
	"socks/config/udp"
	"socks/handlers"
	tcp2 "socks/logger/tcp"
	udp2 "socks/logger/udp"
	"socks/server"
)

func registerServer(builder di.Builder) {
	tcpServerDef := di.Def{
		Name:  "tcp_server",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			connectionHandler := ctn.Get("connection_handler").(handlers.ConnectionHandler)
			config := ctn.Get("tcp_base_config").(tcp.Config).Bind
			logger := ctn.Get("tcp_logger").(tcp2.Logger)

			return server.NewTcpServer(
				connectionHandler,
				logger,
				config,
			)
		},
	}

	udpServerDef := di.Def{
		Name:  "udp_server",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			packetHandler := ctn.Get("packet_handler").(handlers.PacketHandler)
			config := ctn.Get("udp_base_config").(udp.Config).Bind
			logger := ctn.Get("udp_logger").(udp2.Logger)
			replicator := ctn.Get("udp_config_replicator").(udp.ConfigReplicator)

			return server.NewUdpServer(
				logger,
				packetHandler,
				config,
				replicator,
			)
		},
	}

	err := builder.Add(
		tcpServerDef,
		udpServerDef,
	)

	if err != nil {
		panic(err)
	}
}
