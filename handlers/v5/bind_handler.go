package v5

import (
	"net"
	v53 "socks/config/v5"
	"socks/handlers/v5/helpers"
	v52 "socks/logger/v5"
	"socks/managers"
	"socks/protocol/v5"
	"socks/utils"
)

type BindHandler interface {
	HandleBind(config v53.Config, name string, address string, client net.Conn)
}

type BaseBindHandler struct {
	utils        utils.AddressUtils
	logger       v52.Logger
	sender       v5.Sender
	errorHandler ErrorHandler
	bindManager  managers.BindManager
	transmitter  helpers.Transmitter
}

func NewBaseBindHandler(
	utils utils.AddressUtils,
	logger v52.Logger,
	sender v5.Sender,
	errorHandler ErrorHandler,
	bindManager managers.BindManager,
	transmitter helpers.Transmitter,
) (BaseBindHandler, error) {
	return BaseBindHandler{
		utils:        utils,
		logger:       logger,
		sender:       sender,
		errorHandler: errorHandler,
		bindManager:  bindManager,
		transmitter:  transmitter,
	}, nil
}

func (b BaseBindHandler) HandleBind(config v53.Config, name string, address string, client net.Conn) {
	err := b.bindManager.Bind(address)

	if err != nil {
		b.errorHandler.HandleBindManagerBindError(config, err, address, client)

		return
	}

	b.logger.Bind.Successful(client.RemoteAddr().String(), address)

	b.bindSendFirstResponse(config, name, address, client)
}

func (b BaseBindHandler) bindSendFirstResponse(config v53.Config, name string, address string, client net.Conn) {
	err := b.sender.SendSuccessWithTcpPort(config, client)

	if err != nil {
		b.errorHandler.HandleBindIOError(config, err, address, client)

		return
	}

	b.bindWait(config, name, address, client)

	b.bindManager.Remove(address)
}

func (b BaseBindHandler) bindWait(config v53.Config, name string, address string, client net.Conn) {
	host, err := b.bindManager.ReceiveHost(address, config.Deadline.Bind)

	if err != nil {
		b.errorHandler.HandleBindManagerReceiveHostError(config, err, address, client)

		return
	}

	b.bindCheckAddress(config, name, address, host, client)
}

func (b BaseBindHandler) bindCheckAddress(config v53.Config, name string, address string, host, client net.Conn) {
	hostAddr, hostPort, parseErr := b.utils.ParseAddress(host.RemoteAddr().String())

	if parseErr != nil {
		b.errorHandler.HandleAddressParsingError(config, parseErr, address, client, host)

		return
	}

	addrType, determineErr := b.utils.DetermineAddressType(hostAddr)

	if determineErr != nil {
		b.errorHandler.HandleAddressDeterminationError(config, determineErr, address, client, host)

		return
	}

	b.sendSecondResponse(config, name, address, addrType, hostAddr, uint16(hostPort), host, client)
}

func (b BaseBindHandler) sendSecondResponse(config v53.Config, name string, address string, addrType byte, hostAddress string, hostPort uint16, host, client net.Conn) {
	err := b.sender.SendSuccessWithParameters(config, addrType, hostAddress, hostPort, client)

	if err != nil {
		b.errorHandler.HandleBindIOErrorWithHost(config, err, address, client, host)

		return
	}

	err = b.bindManager.SendClient(address, client)

	if err != nil {
		b.errorHandler.HandleBindManagerSendClientError(config, err, address, client, host)

		return
	}

	b.logger.Bind.Bound(client.RemoteAddr().String(), host.RemoteAddr().String())

	err = b.transmitter.TransferBind(config, name, client, host)

	if err != nil {
		b.errorHandler.HandleTransferError(err, client, host)

		return
	}
}
