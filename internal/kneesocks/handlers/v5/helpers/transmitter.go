package helpers

import (
	"net"
	"socks/internal/kneesocks/config/tree"
	"socks/internal/kneesocks/config/v5"
	"socks/internal/kneesocks/managers"
	"socks/internal/kneesocks/transfer"
)

type Transmitter struct {
	connectHandler transfer.ConnectHandler
	bindHandler    transfer.BindHandler
	bindRate       managers.BindRateManager
}

func NewTransmitter(
	connectHandler transfer.ConnectHandler,
	bindHandler transfer.BindHandler,
	bindRate managers.BindRateManager,
) (Transmitter, error) {
	return Transmitter{
		connectHandler: connectHandler,
		bindHandler:    bindHandler,
		bindRate:       bindRate,
	}, nil
}

func (b Transmitter) TransferConnect(config v5.Config, name string, client net.Conn, host net.Conn) {
	var rate tree.RateRestrictions

	user, ok := config.Users[name]

	if !ok {
		rate = tree.RateRestrictions{
			MaxSimultaneousConnections:  -1,
			HostReadBuffersPerSecond:    -1,
			HostWriteBuffersPerSecond:   -1,
			ClientReadBuffersPerSecond:  -1,
			ClientWriteBuffersPerSecond: -1,
		}
	} else {
		rate = user.Restrictions.Rate
	}

	b.connectHandler.HandleClient(rate, client, host)
}

func (b Transmitter) TransferBind(config v5.Config, name string, client net.Conn, host net.Conn) error {
	var rate tree.RateRestrictions

	user, ok := config.Users[name]

	if !ok {
		rate = tree.RateRestrictions{
			MaxSimultaneousConnections:  -1,
			HostReadBuffersPerSecond:    -1,
			HostWriteBuffersPerSecond:   -1,
			ClientReadBuffersPerSecond:  -1,
			ClientWriteBuffersPerSecond: -1,
		}
	} else {
		rate = user.Restrictions.Rate
	}

	err := b.bindRate.Add(client.RemoteAddr().String(), rate)

	if err != nil {
		return err
	}

	defer b.bindRate.Remove(client.RemoteAddr().String())

	b.bindHandler.HandleClient(client, host)

	return nil
}
