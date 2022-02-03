package v5

import (
	"net"
	"socks/handlers/v5/helpers"
	v52 "socks/logger/v5"
	"socks/protocol/v5"
	"socks/transfer"
	"socks/utils"
)

type BindHandler interface {
	HandleBind(name string, address string, client net.Conn)
}

type BaseBindHandler struct {
	streamHandler transfer.StreamHandler
	utils         utils.AddressUtils
	logger        v52.Logger
	sender        v5.Sender
	errorHandler  ErrorHandler
	binder        helpers.Binder
}

func NewBaseBindHandler(
	streamHandler transfer.StreamHandler,
	utils utils.AddressUtils,
	logger v52.Logger,
	sender v5.Sender,
	errorHandler ErrorHandler,
	binder helpers.Binder,
) (BaseBindHandler, error) {
	return BaseBindHandler{
		streamHandler: streamHandler,
		utils:         utils,
		logger:        logger,
		sender:        sender,
		errorHandler:  errorHandler,
		binder:        binder,
	}, nil
}

func (b BaseBindHandler) HandleBind(name string, address string, client net.Conn) {
	err := b.binder.Bind(address)

	if err != nil {
		b.errorHandler.HandleBindManagerBindError(err, address, client)

		return
	}

	b.logger.Bind.Successful(client.RemoteAddr().String(), address)

	b.bindSendFirstResponse(address, client)
}

func (b BaseBindHandler) bindSendFirstResponse(address string, client net.Conn) {
	err := b.sender.SendSuccessWithTcpPort(client)

	if err != nil {
		b.errorHandler.HandleBindIOError(err, address, client)

		return
	}

	b.bindWait(address, client)

	b.binder.Remove(address)
}

func (b BaseBindHandler) bindWait(address string, client net.Conn) {
	host, err := b.binder.Receive(address)

	if err != nil {
		b.errorHandler.HandleBindManagerReceiveHostError(err, address, client)

		return
	}

	b.bindCheckAddress(address, host, client)
}

func (b BaseBindHandler) bindCheckAddress(address string, host, client net.Conn) {
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

	b.sendSecondResponse(address, addrType, hostAddr, uint16(hostPort), host, client)
}

func (b BaseBindHandler) sendSecondResponse(address string, addrType byte, hostAddress string, hostPort uint16, host, client net.Conn) {
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

	b.streamHandler.ClientToHost(host, client)
}
