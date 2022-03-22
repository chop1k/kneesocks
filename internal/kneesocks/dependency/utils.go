package dependency

import (
	"socks/internal/kneesocks/dependency/build"

	"github.com/sarulabs/di"
)

func registerUtils(builder di.Builder) {
	addressUtilsDef := di.Def{
		Name:  "address_utils",
		Scope: di.App,
		Build: build.AddressUtils,
	}

	bufferReaderDef := di.Def{
		Name:  "buffer_reader",
		Scope: di.App,
		Build: build.BufferReader,
	}

	errorUtils := di.Def{
		Name:  "error_utils",
		Scope: di.App,
		Build: build.ErrorUtils,
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
