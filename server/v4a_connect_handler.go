package server

import (
	"net"
	"socks/config"
	"socks/logger"
	"time"
)

type V4aConnectHandler interface {
	HandleV4aConnect(address string, client net.Conn)
}

type BaseV4aConnectHandler struct {
	config        config.SocksV4aConfig
	streamHandler StreamHandler
	logger        logger.SocksV4aLogger
	sender        V4aSender
	errorHandler  V4aErrorHandler
	whitelist     WhitelistManager
	blacklist     BlacklistManager
}

func NewBaseV4aConnectHandler(
	config config.SocksV4aConfig,
	streamHandler StreamHandler,
	logger logger.SocksV4aLogger,
	sender V4aSender,
	errorHandler V4aErrorHandler,
	whitelist WhitelistManager,
	blacklist BlacklistManager,
) (BaseV4aConnectHandler, error) {
	return BaseV4aConnectHandler{
		config:        config,
		streamHandler: streamHandler,
		logger:        logger,
		sender:        sender,
		errorHandler:  errorHandler,
		whitelist:     whitelist,
		blacklist:     blacklist,
	}, nil
}

func (b BaseV4aConnectHandler) HandleV4aConnect(address string, client net.Conn) {
	whitelisted := b.whitelist.IsWhitelisted(address)

	if whitelisted {
		b.sender.SendFailAndClose(client)

		b.logger.ConnectNotAllowedByWhitelist(client.RemoteAddr().String(), address)

		return
	}

	blacklisted := b.blacklist.IsBlacklisted(address)

	if blacklisted {
		b.sender.SendFailAndClose(client)

		b.logger.ConnectNotAllowedByBlacklist(client.RemoteAddr().String(), address)

		return
	}

	b.connect(address, client)
}

func (b BaseV4aConnectHandler) connect(address string, client net.Conn) {
	deadline := time.Second * time.Duration(b.config.GetConnectDeadline())

	host, err := net.DialTimeout("tcp", address, deadline)

	if err != nil {
		b.errorHandler.HandleV4aNetworkError(err, address, client)

		return
	}

	b.connectSendSuccess(address, host, client)
}

func (b BaseV4aConnectHandler) connectSendSuccess(address string, host net.Conn, client net.Conn) {
	err := b.sender.SendSuccess(client)

	if err != nil {
		b.sender.SendFailAndClose(client)

		_ = host.Close()

		b.logger.ConnectFailed(client.RemoteAddr().String(), address)

		return
	}

	b.logger.ConnectSuccessful(client.RemoteAddr().String(), address)

	go b.streamHandler.ClientToHost(client, host)
	b.streamHandler.HostToClient(client, host)
}
