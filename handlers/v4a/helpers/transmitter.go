package helpers

import (
	"net"
	"socks/config/v4a"
	"socks/managers"
	"socks/transfer"
)

type Transmitter interface {
	TransferConnect(config v4a.Config, client net.Conn, host net.Conn)
	TransferBind(config v4a.Config, client net.Conn, host net.Conn) error
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

func (b BaseTransmitter) TransferConnect(config v4a.Config, client net.Conn, host net.Conn) {
	b.connectHandler.HandleClient(config.Restrictions.Rate, client, host)
}

func (b BaseTransmitter) TransferBind(config v4a.Config, client net.Conn, host net.Conn) error {
	err := b.bindRate.Add(client.RemoteAddr().String(), config.Restrictions.Rate)

	if err != nil {
		return err
	}

	defer b.bindRate.Remove(client.RemoteAddr().String())

	b.bindHandler.HandleClient(client, host)

	return nil
}
