package v4

import (
	"net"
	v43 "socks/config/v4"
	"socks/handlers/v4/helpers"
	v42 "socks/logger/v4"
	"socks/protocol/v4"
	"socks/utils"
)

type BindHandler interface {
	HandleBind(config v43.Config, address string, client net.Conn)
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

func (b BaseBindHandler) HandleBind(config v43.Config, address string, client net.Conn) {
	err := b.binder.Bind(address)

	if err != nil {
		b.errorHandler.HandleBindManagerBindError(config, err, address, client)

		return
	}

	b.logger.Bind.Successful(client.RemoteAddr().String(), address)

	b.bindSendFirstResponse(config, address, client)
}

func (b BaseBindHandler) bindSendFirstResponse(config v43.Config, address string, client net.Conn) {
	err := b.sender.SendSuccess(config, client)

	if err != nil {
		b.errorHandler.HandleBindIOError(config, err, address, client)

		return
	}

	b.bindWait(config, address, client)

	b.binder.Remove(address)
}

func (b BaseBindHandler) bindWait(config v43.Config, address string, client net.Conn) {
	host, err := b.binder.Receive(config, address)

	if err != nil {
		b.errorHandler.HandleBindManagerReceiveHostError(config, err, address, client)

		return
	}

	b.bindCheckAddress(config, address, host, client)
}

func (b BaseBindHandler) bindCheckAddress(config v43.Config, address string, host net.Conn, client net.Conn) {
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

	if addrType != 1 {
		b.errorHandler.HandleInvalidAddressTypeError(config, address, client, host)

		return
	}

	b.bindSendSecondResponse(config, address, hostAddr, uint16(hostPort), host, client)
}

func (b BaseBindHandler) bindSendSecondResponse(config v43.Config, address string, hostAddr string, hostPort uint16, host net.Conn, client net.Conn) {
	ip := net.ParseIP(hostAddr)

	if ip == nil {
		b.errorHandler.HandleInvalidAddressTypeError(config, address, client, host)

		return
	}

	ip = ip.To4()

	if ip == nil {
		b.errorHandler.HandleInvalidAddressTypeError(config, address, client, host)

		return
	}

	err := b.sender.SendSuccessWithParameters(config, ip, hostPort, client)

	if err != nil {
		b.errorHandler.HandleBindIOErrorWithHost(config, err, address, client, host)

		return
	}

	err = b.binder.Send(address, client)

	if err != nil {
		b.errorHandler.HandleBindManagerSendClientError(err, address, client, host)

		return
	}

	b.logger.Bind.Bound(client.RemoteAddr().String(), host.RemoteAddr().String())

	err = b.transmitter.TransferBind(config, client, host)

	if err != nil {
		b.errorHandler.HandleTransferError(err, client, host)

		return
	}
}
