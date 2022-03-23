package v4a

import (
	"net"
	v4a3 "socks/internal/kneesocks/config/v4a"
	"socks/internal/kneesocks/handlers/v4a/helpers"
	v4a2 "socks/internal/kneesocks/logger/v4a"
	"socks/internal/kneesocks/managers"
	"socks/pkg/protocol/v4a"
	"socks/pkg/utils"
)

type BindHandler struct {
	logger       v4a2.Logger
	utils        utils.AddressUtils
	sender       v4a.Sender
	errorHandler ErrorHandler
	bindManager  managers.BindManager
	transmitter  helpers.Transmitter
}

func NewBindHandler(
	logger v4a2.Logger,
	utils utils.AddressUtils,
	sender v4a.Sender,
	errorHandler ErrorHandler,
	bindManager managers.BindManager,
	transmitter helpers.Transmitter,
) (*BindHandler, error) {
	return &BindHandler{
		logger:       logger,
		utils:        utils,
		sender:       sender,
		errorHandler: errorHandler,
		bindManager:  bindManager,
		transmitter:  transmitter,
	}, nil
}

func (b BindHandler) HandleBind(config v4a3.Config, address string, client net.Conn) {
	err := b.bindManager.Bind(address)

	if err != nil {
		b.errorHandler.HandleBindManagerBindError(config, err, address, client)

		return
	}

	b.logger.Bind.Successful(client.RemoteAddr().String(), address)

	b.bindSendFirstResponse(config, address, client)
}

func (b BindHandler) bindSendFirstResponse(config v4a3.Config, address string, client net.Conn) {
	err := b.sender.SendSuccess(config, client)

	if err != nil {
		b.errorHandler.HandleBindIOError(config, err, address, client)

		return
	}

	b.bindWait(config, address, client)

	b.bindManager.Remove(address)
}

func (b BindHandler) bindWait(config v4a3.Config, address string, client net.Conn) {
	host, err := b.bindManager.ReceiveHost(address, config.Deadline.Bind)

	if err != nil {
		b.errorHandler.HandleBindManagerReceiveHostError(config, err, address, client)

		return
	}

	b.bindCheckAddress(config, address, host, client)
}

func (b BindHandler) bindCheckAddress(config v4a3.Config, address string, host net.Conn, client net.Conn) {
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

func (b BindHandler) bindSendSecondResponse(config v4a3.Config, address string, hostAddr string, hostPort uint16, host net.Conn, client net.Conn) {
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

	err = b.bindManager.SendClient(address, client)

	if err != nil {
		b.errorHandler.HandleBindManagerSendClientError(err, address, client, host)

		return
	}

	b.logger.Bind.Successful(client.RemoteAddr().String(), address)

	err = b.transmitter.TransferBind(config, client, host)

	if err != nil {
		b.errorHandler.HandleTransferError(err, client, host)

		return
	}
}
