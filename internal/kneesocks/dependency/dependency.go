package dependency

import (
	"github.com/sarulabs/di"
)

func Register(builder di.Builder) {
	registerConfig(builder)
	registerHandlers(builder)
	registerZeroLog(builder)
	registerLogger(builder)
	registerManagers(builder)
	registerProtocol(builder)
	registerServer(builder)
	registerTransfer(builder)
	registerUtils(builder)
}
