package dependency

import (
	"socks/internal/kneesocks/dependency/build"

	"github.com/sarulabs/di"
)

func registerManagers(builder di.Builder) {
	bindManagerDef := di.Def{
		Name:  "bind_manager",
		Scope: di.App,
		Build: build.BindManager,
	}

	bindRateManagerDef := di.Def{
		Name:  "bind_rate_manager",
		Scope: di.App,
		Build: build.BindRateManager,
	}

	connectionsManagerDef := di.Def{
		Name:  "connections_manager",
		Scope: di.App,
		Build: build.ConnectionsManager,
	}

	udpClientManagerDef := di.Def{
		Name:  "udp_client_manager",
		Scope: di.App,
		Build: build.UdpClientManager,
	}

	udpHostManagerDef := di.Def{
		Name:  "udp_host_manager",
		Scope: di.App,
		Build: build.UdpHostManager,
	}

	whitelistManagerDef := di.Def{
		Name:  "whitelist_manager",
		Scope: di.App,
		Build: build.WhitelistManager,
	}

	blacklistManagerDef := di.Def{
		Name:  "blacklist_manager",
		Scope: di.App,
		Build: build.BlacklistManager,
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
