package v4

import (
	"net"
	"socks/config/v4"
	v42 "socks/logger/v4"
	"socks/transfer"
	"time"
)

type ConnectHandler interface {
	HandleConnect(address string, client net.Conn)
}

type BaseConnectHandler struct {
	config        v4.Config
	streamHandler transfer.StreamHandler
	logger        v42.Logger
	sender        Sender
	errorHandler  ErrorHandler
	whitelist     Whitelist
	blacklist     Blacklist
}

func NewBaseConnectHandler(
	config v4.Config,
	streamHandler transfer.StreamHandler,
	logger v42.Logger,
	sender Sender,
	errorHandler ErrorHandler,
	whitelist Whitelist,
	blacklist Blacklist,
) (BaseConnectHandler, error) {
	return BaseConnectHandler{
		config:        config,
		streamHandler: streamHandler,
		logger:        logger,
		sender:        sender,
		errorHandler:  errorHandler,
		whitelist:     whitelist,
		blacklist:     blacklist,
	}, nil
}

func (b BaseConnectHandler) HandleConnect(address string, client net.Conn) {
	whitelisted := b.whitelist.IsWhitelisted(address)

	if whitelisted {
		b.sender.SendFailAndClose(client)

		b.logger.Restrictions.NotAllowedByWhitelist(client.RemoteAddr().String(), address)

		return
	}

	blacklisted := b.blacklist.IsBlacklisted(address)

	if blacklisted {
		b.sender.SendFailAndClose(client)

		b.logger.Restrictions.NotAllowedByBlacklist(client.RemoteAddr().String(), address)

		return
	}

	b.connect(address, client)
}

func (b BaseConnectHandler) connect(address string, client net.Conn) {
	deadline := time.Second * time.Duration(b.config.GetConnectDeadline())

	host, err := net.DialTimeout("tcp4", address, deadline)

	if err != nil {
		b.errorHandler.HandleDialError(err, address, client)

		return
	}

	b.connectSendSuccess(address, host, client)
}

func (b BaseConnectHandler) connectSendSuccess(address string, host net.Conn, client net.Conn) {
	err := b.sender.SendSuccess(client)

	if err != nil {
		b.errorHandler.HandleConnectIOErrorWithHost(err, address, client, host)

		return
	}

	b.logger.Connect.ConnectSuccessful(client.RemoteAddr().String(), address)

	go b.streamHandler.ClientToHost(client, host)
	b.streamHandler.HostToClient(client, host)
}
