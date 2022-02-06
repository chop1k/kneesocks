package main

import (
	"github.com/sarulabs/di"
	"socks/logger"
	"socks/managers"
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

	udpBindManagerDef := di.Def{
		Name:  "udp_bind_manager",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return managers.NewUdpBindManager()
		},
	}

	whitelistManagerDef := di.Def{
		Name:  "whitelist_manager",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			serverLogger := ctn.Get("server_logger").(logger.ServerLogger)

			return managers.NewBaseWhitelistManager(serverLogger)
		},
	}

	blacklistManagerDef := di.Def{
		Name:  "blacklist_manager",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			serverLogger := ctn.Get("server_logger").(logger.ServerLogger)

			return managers.NewBaseBlacklistManager(serverLogger)
		},
	}

	err := builder.Add(
		bindManagerDef,
		bindRateManagerDef,
		udpClientManagerDef,
		udpHostManagerDef,
		udpBindManagerDef,
		whitelistManagerDef,
		blacklistManagerDef,
	)

	if err != nil {
		panic(err)
	}
}
