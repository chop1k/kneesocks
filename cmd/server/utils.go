package main

import (
	"github.com/sarulabs/di"
	"socks/utils"
)

func registerUtils(builder di.Builder) {
	addressUtilsDef := di.Def{
		Name:  "address_utils",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return utils.NewUtils()
		},
	}

	errorUtils := di.Def{
		Name:  "error_utils",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return utils.NewErrorUtils()
		},
	}

	err := builder.Add(
		addressUtilsDef,
		errorUtils,
	)

	if err != nil {
		panic(err)
	}
}
