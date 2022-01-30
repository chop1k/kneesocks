package v5

import (
	"net"
	"socks/config"
	"socks/logger"
	"socks/transfer"
	"socks/utils"
	"time"
)

type ConnectHandler interface {
	HandleConnect(name string, address string, client net.Conn)
}

type BaseConnectHandler struct {
	config        config.SocksV5Config
	streamHandler transfer.StreamHandler
	logger        logger.SocksV5Logger
	utils         utils.AddressUtils
	sender        Sender
	errorHandler  ErrorHandler
	whitelist     Whitelist
	blacklist     Blacklist
}

func NewBaseConnectHandler(
	config config.SocksV5Config,
	streamHandler transfer.StreamHandler,
	logger logger.SocksV5Logger,
	addressUtils utils.AddressUtils,
	sender Sender,
	errorHandler ErrorHandler,
	whitelist Whitelist,
	blacklist Blacklist,
) (BaseConnectHandler, error) {
	return BaseConnectHandler{
		config:        config,
		streamHandler: streamHandler,
		logger:        logger,
		utils:         addressUtils,
		sender:        sender,
		errorHandler:  errorHandler,
		whitelist:     whitelist,
		blacklist:     blacklist,
	}, nil
}

func (b BaseConnectHandler) HandleConnect(name string, address string, client net.Conn) {
	whitelisted := b.whitelist.IsWhitelisted(name, address)

	if whitelisted {
		b.sender.SendConnectionNotAllowedAndClose(client)

		b.logger.ConnectNotAllowedByWhitelist(client.RemoteAddr().String(), address)

		return
	}

	blacklisted := b.blacklist.IsBlacklisted(name, address)

	if blacklisted {
		b.sender.SendConnectionNotAllowedAndClose(client)

		b.logger.ConnectNotAllowedByBlacklist(client.RemoteAddr().String(), address)

		return
	}

	b.connect(address, client)
}

func (b BaseConnectHandler) connect(address string, client net.Conn) {
	deadline := time.Second * time.Duration(b.config.GetConnectDeadline())

	host, err := net.DialTimeout("tcp", address, deadline)

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

	b.logger.ConnectSuccessful(client.RemoteAddr().String(), host.RemoteAddr().String())

	go b.streamHandler.ClientToHost(client, host)
	b.streamHandler.HostToClient(client, host)
}
