package dependency

import (
	"socks/internal/kneesocks/config/tcp"
	"socks/internal/kneesocks/config/udp"
	"socks/internal/kneesocks/handlers"
	tcpLogger "socks/internal/kneesocks/logger/tcp"
	udpLogger "socks/internal/kneesocks/logger/udp"
	"socks/internal/kneesocks/server"

	"github.com/sarulabs/di"
)

func registerServer(builder di.Builder) {
	tcpServerDef := di.Def{
		Name:  "tcp_server",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			connectionHandler := ctn.Get("connection_handler").(handlers.ConnectionHandler)
			config := ctn.Get("tcp_base_config").(tcp.Config).Bind
			logger := ctn.Get("tcp_logger").(tcpLogger.Logger)

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
			_packetHandler := ctn.Get("packet_handler")

			if _packetHandler == nil {
				return nil, nil
			}

			packetHandler := _packetHandler.(handlers.PacketHandler)

			config := ctn.Get("udp_base_config").(udp.Config).Bind
			logger := ctn.Get("udp_logger").(udpLogger.Logger)
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
