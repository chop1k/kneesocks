package helpers

import (
	"net"
	"socks/config/tree"
	v5 "socks/config/v5"
	"socks/managers"
	"socks/transfer"
)

type Transmitter interface {
	TransferConnect(name string, client net.Conn, host net.Conn)
	TransferBind(name string, client net.Conn, host net.Conn)
}

type BaseTransmitter struct {
	config         v5.UsersConfig
	connectHandler transfer.ConnectHandler
	bindHandler    transfer.BindHandler
	bindRate       managers.BindRateManager
}

func NewBaseTransmitter(
	config v5.UsersConfig,
	connectHandler transfer.ConnectHandler,
	bindHandler transfer.BindHandler,
	bindRate managers.BindRateManager,
) (BaseTransmitter, error) {
	return BaseTransmitter{
		config:         config,
		connectHandler: connectHandler,
		bindHandler:    bindHandler,
		bindRate:       bindRate,
	}, nil
}

func (b BaseTransmitter) TransferConnect(name string, client net.Conn, host net.Conn) {
	rate, err := b.config.GetRate(name)

	if err != nil && err == v5.UserNotExistsError {
		rate = tree.RateRestrictions{
			MaxSimultaneousConnections:  -1,
			HostReadBuffersPerSecond:    -1,
			HostWriteBuffersPerSecond:   -1,
			ClientReadBuffersPerSecond:  -1,
			ClientWriteBuffersPerSecond: -1,
		}
	}

	b.connectHandler.HandleClient(rate, client, host)
}

func (b BaseTransmitter) TransferBind(name string, client net.Conn, host net.Conn) {
	rate, err := b.config.GetRate(name)

	if err != nil && err == v5.UserNotExistsError {
		rate = tree.RateRestrictions{
			MaxSimultaneousConnections:  -1,
			HostReadBuffersPerSecond:    -1,
			HostWriteBuffersPerSecond:   -1,
			ClientReadBuffersPerSecond:  -1,
			ClientWriteBuffersPerSecond: -1,
		}
	}

	err = b.bindRate.Add(client.RemoteAddr().String(), rate)

	if err != nil {
		panic(err) // TODO: fix
	}

	defer b.bindRate.Remove(client.RemoteAddr().String())

	b.bindHandler.HandleClient(client, host)
}
