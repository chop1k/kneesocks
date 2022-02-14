package dependency

import (
	"github.com/sarulabs/di"
	"socks/pkg/utils"
)

func registerUtils(builder di.Builder) {
	addressUtilsDef := di.Def{
		Name:  "address_utils",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return utils.NewUtils()
		},
	}

	bufferReaderDef := di.Def{
		Name:  "buffer_reader",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return utils.NewBufferReader()
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
		bufferReaderDef,
		errorUtils,
	)

	if err != nil {
		panic(err)
	}
}
