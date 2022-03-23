package v5

import (
	"net"
	v53 "socks/internal/kneesocks/config/v5"
	"socks/internal/kneesocks/handlers/v5/helpers"
	v52 "socks/internal/kneesocks/logger/v5"
	v5 "socks/pkg/protocol/v5"
	"socks/pkg/utils"
)

type ConnectHandler struct {
	logger       v52.Logger
	utils        utils.AddressUtils
	sender       v5.Sender
	errorHandler ErrorHandler
	transmitter  helpers.Transmitter
}

func NewConnectHandler(
	logger v52.Logger,
	addressUtils utils.AddressUtils,
	sender v5.Sender,
	errorHandler ErrorHandler,
	transmitter helpers.Transmitter,
) (*ConnectHandler, error) {
	return &ConnectHandler{
		logger:       logger,
		utils:        addressUtils,
		sender:       sender,
		errorHandler: errorHandler,
		transmitter:  transmitter,
	}, nil
}

func (b ConnectHandler) HandleConnect(config v53.Config, name string, address string, client net.Conn) {
	var host net.Conn
	var err error

	if config.Deadline.Connect > 0 {
		host, err = net.DialTimeout("tcp", address, config.Deadline.Connect)
	} else {
		host, err = net.Dial("tcp", address)
	}

	if err != nil {
		b.errorHandler.HandleDialError(config, err, address, client)

		return
	}

	b.connectSendResponse(config, name, address, host, client)
}

func (b ConnectHandler) connectSendResponse(config v53.Config, name string, address string, host, client net.Conn) {
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
