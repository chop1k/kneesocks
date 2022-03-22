package build

import (
	"socks/internal/kneesocks/managers"
	"socks/internal/kneesocks/transfer"

	"github.com/sarulabs/di"
)

func TransferBindHandler(ctn di.Container) (interface{}, error) {
	bindRate := ctn.Get("bind_rate_manager").(managers.BindRateManager)
	handler := ctn.Get("transfer_handler").(transfer.Handler)

	return transfer.NewBindHandler(bindRate, handler)
}

func TransferConnectHandler(ctn di.Container) (interface{}, error) {
	handler := ctn.Get("transfer_handler").(transfer.Handler)

	return transfer.NewConnectHandler(handler)
}

func TransferHandler(ctn di.Container) (interface{}, error) {
	return transfer.NewHandler()
}
