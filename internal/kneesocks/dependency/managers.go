package dependency

import (
	"github.com/sarulabs/di"
	"socks/internal/kneesocks/managers"
)

func registerManagers(builder di.Builder) {
	bindManagerDef := di.Def{
		Name:  "bind_manager",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return managers.NewBindManager(), nil
		},
	}

	bindRateManagerDef := di.Def{
		Name:  "bind_rate_manager",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return managers.NewBindRateManager()
		},
	}

	connectionsManagerDef := di.Def{
		Name:  "connections_manager",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return managers.NewConnectionsManager()
		},
	}

	udpClientManagerDef := di.Def{
		Name:  "udp_client_manager",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return managers.NewUdpClientManager()
		},
	}

	udpHostManagerDef := di.Def{
		Name:  "udp_host_manager",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return managers.NewUdpHostManager()
		},
	}

	whitelistManagerDef := di.Def{
		Name:  "whitelist_manager",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return managers.NewWhitelistManager()
		},
	}

	blacklistManagerDef := di.Def{
		Name:  "blacklist_manager",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return managers.NewBlacklistManager()
		},
	}

	err := builder.Add(
		bindManagerDef,
		bindRateManagerDef,
		connectionsManagerDef,
		udpClientManagerDef,
		udpHostManagerDef,
		whitelistManagerDef,
		blacklistManagerDef,
	)

	if err != nil {
		panic(err)
	}
}
