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
			cfg := ctn.Get("tcp_config").(tcp.BindConfig)
			logger := ctn.Get("tcp_logger").(tcp2.Logger)

			return server.NewTcpServer(
				connectionHandler,
				logger,
				cfg,
			)
		},
	}

	udpServerDef := di.Def{
		Name:  "udp_server",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			packetHandler := ctn.Get("packet_handler").(handlers.PacketHandler)
			cfg := ctn.Get("udp_config").(udp.BindConfig)
			logger := ctn.Get("udp_logger").(udp2.Logger)
			buffer := ctn.Get("udp_buffer_config").(udp.BufferConfig)

			return server.NewUdpServer(
				cfg,
				logger,
				packetHandler,
				buffer,
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
