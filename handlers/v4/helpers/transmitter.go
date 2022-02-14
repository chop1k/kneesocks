package helpers

import (
	"net"
	v4 "socks/config/v4"
	"socks/managers"
	"socks/transfer"
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

func (b Transmitter) TransferConnect(config v4.Config, client net.Conn, host net.Conn) {
	b.connectHandler.HandleClient(config.Restrictions.Rate, client, host)
}

func (b Transmitter) TransferBind(config v4.Config, client net.Conn, host net.Conn) error {
	err := b.bindRate.Add(client.RemoteAddr().String(), config.Restrictions.Rate)

	if err != nil {
		return err
	}

	defer b.bindRate.Remove(client.RemoteAddr().String())

	b.bindHandler.HandleClient(client, host)

	return nil
}
