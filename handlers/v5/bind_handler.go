package v5

import (
	"net"
	"socks/handlers/v5/helpers"
	v52 "socks/logger/v5"
	"socks/protocol/v5"
	"socks/utils"
)

type BindHandler interface {
	HandleBind(name string, address string, client net.Conn)
}

type BaseBindHandler struct {
	utils        utils.AddressUtils
	logger       v52.Logger
	sender       v5.Sender
	errorHandler ErrorHandler
	binder       helpers.Binder
	transmitter  helpers.Transmitter
}

func NewBaseBindHandler(
	utils utils.AddressUtils,
	logger v52.Logger,
	sender v5.Sender,
	errorHandler ErrorHandler,
	binder helpers.Binder,
	transmitter helpers.Transmitter,
) (BaseBindHandler, error) {
	return BaseBindHandler{
		utils:        utils,
		logger:       logger,
		sender:       sender,
		errorHandler: errorHandler,
		binder:       binder,
		transmitter:  transmitter,
	}, nil
}

func (b BaseBindHandler) HandleBind(name string, address string, client net.Conn) {
	err := b.binder.Bind(address)

	if err != nil {
		b.errorHandler.HandleBindManagerBindError(err, address, client)

		return
	}

	b.logger.Bind.Successful(client.RemoteAddr().String(), address)

	b.bindSendFirstResponse(name, address, client)
}

func (b BaseBindHandler) bindSendFirstResponse(name string, address string, client net.Conn) {
	err := b.sender.SendSuccessWithTcpPort(client)

	if err != nil {
		b.errorHandler.HandleBindIOError(err, address, client)

		return
	}

	b.bindWait(name, address, client)

	b.binder.Remove(address)
}

func (b BaseBindHandler) bindWait(name string, address string, client net.Conn) {
	host, err := b.binder.Receive(address)

	if err != nil {
		b.errorHandler.HandleBindManagerReceiveHostError(err, address, client)

		return
	}

	b.bindCheckAddress(name, address, host, client)
}

func (b BaseBindHandler) bindCheckAddress(name string, address string, host, client net.Conn) {
	hostAddr, hostPort, parseErr := b.utils.ParseAddress(host.RemoteAddr().String())

	if parseErr != nil {
		b.errorHandler.HandleAddressParsingError(parseErr, address, client, host)

		return
	}

	addrType, determineErr := b.utils.DetermineAddressType(hostAddr)

	if determineErr != nil {
		b.errorHandler.HandleAddressDeterminationError(determineErr, address, client, host)

		return
	}

	b.sendSecondResponse(name, address, addrType, hostAddr, uint16(hostPort), host, client)
}

func (b BaseBindHandler) sendSecondResponse(name string, address string, addrType byte, hostAddress string, hostPort uint16, host, client net.Conn) {
	err := b.sender.SendSuccessWithParameters(addrType, hostAddress, hostPort, client)

	if err != nil {
		b.errorHandler.HandleBindIOErrorWithHost(err, address, client, host)

		return
	}

	err = b.binder.Send(address, client)

	if err != nil {
		b.errorHandler.HandleBindManagerSendClientError(err, address, client, host)

		return
	}

	b.logger.Bind.Bound(client.RemoteAddr().String(), host.RemoteAddr().String())

	b.transmitter.TransferConnect(name, client, host)
}
