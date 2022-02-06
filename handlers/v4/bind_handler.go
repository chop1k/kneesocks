package v4

import (
	"net"
	"socks/handlers/v4/helpers"
	v42 "socks/logger/v4"
	"socks/protocol/v4"
	"socks/utils"
)

type BindHandler interface {
	HandleBind(address string, client net.Conn)
}

type BaseBindHandler struct {
	logger       v42.Logger
	utils        utils.AddressUtils
	sender       v4.Sender
	errorHandler ErrorHandler
	binder       helpers.Binder
	transmitter  helpers.Transmitter
}

func NewBaseBindHandler(
	logger v42.Logger,
	utils utils.AddressUtils,
	sender v4.Sender,
	errorHandler ErrorHandler,
	binder helpers.Binder,
	transmitter helpers.Transmitter,
) (BaseBindHandler, error) {
	return BaseBindHandler{
		logger:       logger,
		utils:        utils,
		sender:       sender,
		errorHandler: errorHandler,
		binder:       binder,
		transmitter:  transmitter,
	}, nil
}

func (b BaseBindHandler) HandleBind(address string, client net.Conn) {
	err := b.binder.Bind(address)

	if err != nil {
		b.errorHandler.HandleBindManagerBindError(err, address, client)

		return
	}

	b.logger.Bind.Successful(client.RemoteAddr().String(), address)

	b.bindSendFirstResponse(address, client)
}

func (b BaseBindHandler) bindSendFirstResponse(address string, client net.Conn) {
	err := b.sender.SendSuccess(client)

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

func (b BaseBindHandler) bindCheckAddress(address string, host net.Conn, client net.Conn) {
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

	if addrType != 1 {
		b.errorHandler.HandleInvalidAddressTypeError(address, client, host)

		return
	}

	b.bindSendSecondResponse(address, hostAddr, uint16(hostPort), host, client)
}

func (b BaseBindHandler) bindSendSecondResponse(address string, hostAddr string, hostPort uint16, host net.Conn, client net.Conn) {
	ip := net.ParseIP(hostAddr)

	if ip == nil {
		b.errorHandler.HandleInvalidAddressTypeError(address, client, host)

		return
	}

	ip = ip.To4()

	if ip == nil {
		b.errorHandler.HandleInvalidAddressTypeError(address, client, host)

		return
	}

	err := b.sender.SendSuccessWithParameters(ip, hostPort, client)

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

	b.transmitter.TransferBind(client, host)
}
