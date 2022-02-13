package v5

import (
	"net"
	v53 "socks/config/v5"
	"socks/handlers/v5/helpers"
	v52 "socks/logger/v5"
	"socks/protocol/v5"
	"socks/utils"
)

type ConnectHandler interface {
	HandleConnect(config v53.Config, name string, address string, client net.Conn)
}

type BaseConnectHandler struct {
	logger       v52.Logger
	utils        utils.AddressUtils
	sender       v5.Sender
	errorHandler ErrorHandler
	transmitter  helpers.Transmitter
}

func NewBaseConnectHandler(
	logger v52.Logger,
	addressUtils utils.AddressUtils,
	sender v5.Sender,
	errorHandler ErrorHandler,
	transmitter helpers.Transmitter,
) (BaseConnectHandler, error) {
	return BaseConnectHandler{
		logger:       logger,
		utils:        addressUtils,
		sender:       sender,
		errorHandler: errorHandler,
		transmitter:  transmitter,
	}, nil
}

func (b BaseConnectHandler) HandleConnect(config v53.Config, name string, address string, client net.Conn) {
	host, err := net.DialTimeout("tcp", address, config.Deadline.Connect)

	if err != nil {
		b.errorHandler.HandleDialError(config, err, address, client)

		return
	}

	b.connectSendResponse(config, name, address, host, client)
}

func (b BaseConnectHandler) connectSendResponse(config v53.Config, name string, address string, host, client net.Conn) {
	addr, port, parseErr := b.utils.ParseAddress(host.RemoteAddr().String())

	if parseErr != nil {
		b.errorHandler.HandleAddressParsingError(config, parseErr, address, client, host)

		return
	}

	addrType, determineErr := b.utils.DetermineAddressType(addr)

	if determineErr != nil {
		b.errorHandler.HandleAddressDeterminationError(config, determineErr, address, client, host)

		return
	}

	responseErr := b.sender.SendSuccessWithParameters(config, addrType, addr, uint16(port), client)

	if responseErr != nil {
		b.errorHandler.HandleConnectIOErrorWithHost(config, responseErr, address, client, host)

		return
	}

	b.logger.Connect.Successful(client.RemoteAddr().String(), host.RemoteAddr().String())

	b.transmitter.TransferConnect(config, name, client, host)
}
