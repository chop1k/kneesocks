package dependency

import "github.com/sarulabs/di"

func Register(builder di.Builder) {
	registerCli(builder)
}
