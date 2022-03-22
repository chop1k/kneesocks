package build

import (
	"socks/pkg/utils"

	"github.com/sarulabs/di"
)

func AddressUtils(ctn di.Container) (interface{}, error) {
	return utils.NewUtils()
}

func BufferReader(ctn di.Container) (interface{}, error) {
	return utils.NewBufferReader()
}

func ErrorUtils(ctn di.Container) (interface{}, error) {
	return utils.NewErrorUtils()
}
