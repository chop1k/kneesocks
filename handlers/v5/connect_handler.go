package v5

import (
	"net"
	"socks/handlers/v5/helpers"
	v52 "socks/logger/v5"
	"socks/protocol/v5"
	"socks/transfer"
	"socks/utils"
)

type ConnectHandler interface {
	HandleConnect(name string, address string, client net.Conn)
}

type BaseConnectHandler struct {
	streamHandler transfer.StreamHandler
	logger        v52.Logger
	utils         utils.AddressUtils
	sender        v5.Sender
	errorHandler  ErrorHandler
	dialer        helpers.Dialer
}

func NewBaseConnectHandler(
	streamHandler transfer.StreamHandler,
	logger v52.Logger,
	addressUtils utils.AddressUtils,
	sender v5.Sender,
	errorHandler ErrorHandler,
	dialer helpers.Dialer,
) (BaseConnectHandler, error) {
	return BaseConnectHandler{
		streamHandler: streamHandler,
		logger:        logger,
		utils:         addressUtils,
		sender:        sender,
		errorHandler:  errorHandler,
		dialer:        dialer,
	}, nil
}

func (b BaseConnectHandler) HandleConnect(name string, address string, client net.Conn) {
	host, err := b.dialer.Dial(address)

	if err != nil {
		b.errorHandler.HandleDialError(err, address, client)

		return
	}

	b.connectSendResponse(address, host, client)
}

func (b BaseConnectHandler) connectSendResponse(address string, host, client net.Conn) {
	addr, port, parseErr := b.utils.ParseAddress(host.RemoteAddr().String())

	if parseErr != nil {
		b.errorHandler.HandleAddressParsingError(parseErr, address, client, host)

		return
	}

	addrType, determineErr := b.utils.DetermineAddressType(addr)

	if determineErr != nil {
		b.errorHandler.HandleAddressDeterminationError(determineErr, address, client, host)

		return
	}

	responseErr := b.sender.SendSuccessWithParameters(addrType, addr, uint16(port), client)

	if responseErr != nil {
		b.errorHandler.HandleConnectIOErrorWithHost(responseErr, address, client, host)

		return
	}

	b.logger.Connect.Successful(client.RemoteAddr().String(), host.RemoteAddr().String())

	go b.streamHandler.ClientToHost(client, host)
	b.streamHandler.HostToClient(client, host)
}
