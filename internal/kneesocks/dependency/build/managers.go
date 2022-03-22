package build

import (
	"socks/internal/kneesocks/managers"

	"github.com/sarulabs/di"
)

func BindManager(ctn di.Container) (interface{}, error) {
	return managers.NewBindManager(), nil
}

func BindRateManager(ctn di.Container) (interface{}, error) {
	return managers.NewBindRateManager()
}

func ConnectionsManager(ctn di.Container) (interface{}, error) {
	return managers.NewConnectionsManager()
}

func UdpClientManager(ctn di.Container) (interface{}, error) {
	return managers.NewUdpClientManager()
}

func UdpHostManager(ctn di.Container) (interface{}, error) {
	return managers.NewUdpHostManager()
}

func WhitelistManager(ctn di.Container) (interface{}, error) {
	return managers.NewWhitelistManager()
}

func BlacklistManager(ctn di.Container) (interface{}, error) {
	return managers.NewBlacklistManager()
}
