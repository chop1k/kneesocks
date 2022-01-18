package server

import (
	"net"
	"socks/config"
	"socks/logger"
	"socks/utils"
	"time"
)

type V5ConnectHandler interface {
	HandleV5Connect(address string, client net.Conn)
}

type BaseV5ConnectHandler struct {
	config        config.SocksV5Config
	streamHandler StreamHandler
	logger        logger.SocksV5Logger
	utils         utils.AddressUtils
	sender        V5Sender
	errorHandler  V5ErrorHandler
	whitelist     WhitelistManager
	blacklist     BlacklistManager
}

func NewBaseV5ConnectHandler(
	config config.SocksV5Config,
	streamHandler StreamHandler,
	logger logger.SocksV5Logger,
	addressUtils utils.AddressUtils,
	sender V5Sender,
	errorHandler V5ErrorHandler,
	whitelist WhitelistManager,
	blacklist BlacklistManager,
) (BaseV5ConnectHandler, error) {
	return BaseV5ConnectHandler{
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

func (b BaseV5ConnectHandler) HandleV5Connect(address string, client net.Conn) {
	whitelisted := b.whitelist.IsWhitelisted(address)

	if whitelisted {
		b.sender.SendConnectionNotAllowedAndClose(client)

		b.logger.ConnectNotAllowedByWhitelist(client.RemoteAddr().String(), address)

		return
	}

	blacklisted := b.blacklist.IsBlacklisted(address)

	if blacklisted {
		b.sender.SendConnectionNotAllowedAndClose(client)

		b.logger.ConnectNotAllowedByBlacklist(client.RemoteAddr().String(), address)

		return
	}

	deadline := time.Second * time.Duration(b.config.GetConnectDeadline())

	host, err := net.DialTimeout("tcp", address, deadline)

	if err != nil {
		b.errorHandler.HandleV5NetworkError(err, address, client)

		return
	}

	b.connectSendResponse(host, client)
}

func (b BaseV5ConnectHandler) connectSendResponse(host, client net.Conn) {
	addr, port, parseErr := b.utils.ParseAddress(host.RemoteAddr().String())

	if parseErr != nil {
		b.sender.SendFailAndClose(client)

		_ = host.Close()

		//b.errors.ParseAddressError(client.RemoteAddr().String(), host.RemoteAddr().String())

		return
	}

	addrType, determineErr := b.utils.DetermineAddressType(addr)

	if determineErr != nil {
		b.sender.SendFailAndClose(client)

		_ = host.Close()

		//b.errors.DetermineAddressError(client.RemoteAddr().String(), addr)

		return
	}

	responseErr := b.sender.SendSuccessWithParameters(addrType, addr, uint16(port), client)

	if responseErr != nil {
		b.errorHandler.HandleV5NetworkError(responseErr, host.RemoteAddr().String(), client)

		_ = host.Close()

		return
	}

	b.logger.ConnectSuccessful(client.RemoteAddr().String(), host.RemoteAddr().String())

	go b.streamHandler.ClientToHost(client, host)
	b.streamHandler.HostToClient(client, host)
}
