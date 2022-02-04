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

	udpAssociationManagerDef := di.Def{
		Name:  "udp_association_manager",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return managers.NewUdpAssociationManager(), nil
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
		udpAssociationManagerDef,
		whitelistManagerDef,
		blacklistManagerDef,
	)

	if err != nil {
		panic(err)
	}
}
