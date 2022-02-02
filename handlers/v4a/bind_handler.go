package v4a

import (
	"net"
	"socks/handlers/v4a/helpers"
	v4a2 "socks/logger/v4a"
	"socks/managers"
	"socks/transfer"
	"socks/utils"
)

type BindHandler interface {
	HandleBind(address string, client net.Conn)
}

type BaseBindHandler struct {
	logger        v4a2.Logger
	streamHandler transfer.StreamHandler
	bindManager   managers.BindManager
	utils         utils.AddressUtils
	sender        helpers.Sender
	errorHandler  ErrorHandler
	receiver      helpers.Receiver
}

func NewBaseBindHandler(
	logger v4a2.Logger,
	streamHandler transfer.StreamHandler,
	bindManager managers.BindManager,
	utils utils.AddressUtils,
	sender helpers.Sender,
	errorHandler ErrorHandler,
	receiver helpers.Receiver,
) (BaseBindHandler, error) {
	return BaseBindHandler{
		logger:        logger,
		streamHandler: streamHandler,
		bindManager:   bindManager,
		utils:         utils,
		sender:        sender,
		errorHandler:  errorHandler,
		receiver:      receiver,
	}, nil
}

func (b BaseBindHandler) HandleBind(address string, client net.Conn) {
	err := b.bindManager.Bind(address)

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

	b.bindManager.Remove(address)
}

func (b BaseBindHandler) bindWait(address string, client net.Conn) {
	host, err := b.receiver.ReceiveHost(address)

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

	err = b.bindManager.SendClient(address, client)

	if err != nil {
		b.errorHandler.HandleBindManagerSendClientError(err, address, client, host)

		return
	}

	b.logger.Bind.Successful(client.RemoteAddr().String(), address)

	b.streamHandler.ClientToHost(host, client)
}
