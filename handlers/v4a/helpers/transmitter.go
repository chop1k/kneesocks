package helpers

import (
	"net"
	"socks/config/v4a"
	"socks/managers"
	"socks/transfer"
)

type Transmitter interface {
	TransferConnect(client net.Conn, host net.Conn) error
	TransferBind(client net.Conn, host net.Conn) error
}

type BaseTransmitter struct {
	config         v4a.RestrictionsConfig
	connectHandler transfer.ConnectHandler
	bindHandler    transfer.BindHandler
	bindRate       managers.BindRateManager
}

func NewBaseTransmitter(
	config v4a.RestrictionsConfig,
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

func (b BaseTransmitter) TransferConnect(client net.Conn, host net.Conn) error {
	rate, err := b.config.GetRate()

	if err != nil {
		return err
	}

	b.connectHandler.HandleClient(rate, client, host)

	return nil
}

func (b BaseTransmitter) TransferBind(client net.Conn, host net.Conn) error {
	rate, err := b.config.GetRate()

	if err != nil {
		return err
	}

	err = b.bindRate.Add(client.RemoteAddr().String(), rate)

	if err != nil {
		return err
	}

	defer b.bindRate.Remove(client.RemoteAddr().String())

	b.bindHandler.HandleClient(client, host)

	return nil
}
