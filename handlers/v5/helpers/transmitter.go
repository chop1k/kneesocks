package helpers

import (
	"net"
	"socks/config/tree"
	v5 "socks/config/v5"
	"socks/managers"
	"socks/transfer"
)

type Transmitter interface {
	TransferConnect(config v5.Config, name string, client net.Conn, host net.Conn)
	TransferBind(config v5.Config, name string, client net.Conn, host net.Conn) error
}

type BaseTransmitter struct {
	connectHandler transfer.ConnectHandler
	bindHandler    transfer.BindHandler
	bindRate       managers.BindRateManager
}

func NewBaseTransmitter(
	connectHandler transfer.ConnectHandler,
	bindHandler transfer.BindHandler,
	bindRate managers.BindRateManager,
) (BaseTransmitter, error) {
	return BaseTransmitter{
		connectHandler: connectHandler,
		bindHandler:    bindHandler,
		bindRate:       bindRate,
	}, nil
}

func (b BaseTransmitter) TransferConnect(config v5.Config, name string, client net.Conn, host net.Conn) {
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

func (b BaseTransmitter) TransferBind(config v5.Config, name string, client net.Conn, host net.Conn) error {
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
