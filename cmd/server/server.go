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
	serverDef := di.Def{
		Name:  "server",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			connectionHandler := ctn.Get("connection_handler").(handlers.ConnectionHandler)
			packetHandler := ctn.Get("packet_handler").(handlers.PacketHandler)
			tcpConfig := ctn.Get("tcp_config").(tcp.Config)
			tcpLogger := ctn.Get("tcp_logger").(tcp2.Logger)
			udpConfig := ctn.Get("udp_config").(udp.Config)
			udpLogger := ctn.Get("udp_logger").(udp2.Logger)

			return server.NewServer(
				connectionHandler,
				packetHandler,
				tcpLogger,
				tcpConfig,
				udpLogger,
				udpConfig,
			)
		},
	}

	err := builder.Add(
		serverDef,
	)

	if err != nil {
		panic(err)
	}
}
