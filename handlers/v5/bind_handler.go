package v5

import (
	"net"
	"socks/config/v5"
	v52 "socks/logger/v5"
	"socks/managers"
	"socks/transfer"
	"socks/utils"
	"time"
)

type BindHandler interface {
	HandleBind(name string, address string, client net.Conn)
}

type BaseBindHandler struct {
	bindManager   managers.BindManager
	config        v5.Config
	streamHandler transfer.StreamHandler
	utils         utils.AddressUtils
	logger        v52.Logger
	sender        Sender
	whitelist     Whitelist
	blacklist     Blacklist
	errorHandler  ErrorHandler
}

func NewBaseBindHandler(
	bindManager managers.BindManager,
	config v5.Config,
	streamHandler transfer.StreamHandler,
	utils utils.AddressUtils,
	logger v52.Logger,
	sender Sender,
	whitelist Whitelist,
	blacklist Blacklist,
	errorHandler ErrorHandler,
) (BaseBindHandler, error) {
	return BaseBindHandler{
		bindManager:   bindManager,
		config:        config,
		streamHandler: streamHandler,
		utils:         utils,
		logger:        logger,
		sender:        sender,
		whitelist:     whitelist,
		blacklist:     blacklist,
		errorHandler:  errorHandler,
	}, nil
}

func (b BaseBindHandler) HandleBind(name string, address string, client net.Conn) {
	whitelisted := b.whitelist.IsWhitelisted(name, address)

	if whitelisted {
		b.sender.SendConnectionNotAllowedAndClose(client)

		b.logger.Restrictions.NotAllowedByWhitelist(client.RemoteAddr().String(), address)

		return
	}

	blacklisted := b.blacklist.IsBlacklisted(name, address)

	if blacklisted {
		b.sender.SendConnectionNotAllowedAndClose(client)

		b.logger.Restrictions.NotAllowedByBlacklist(client.RemoteAddr().String(), address)

		return
	}

	b.bind(address, client)
}

func (b BaseBindHandler) bind(address string, client net.Conn) {
	err := b.bindManager.Bind(address)

	if err != nil {
		b.errorHandler.HandleBindManagerBindError(err, address, client)

		return
	}

	b.logger.Bind.BindSuccessful(client.RemoteAddr().String(), address)

	b.bindSendFirstResponse(address, client)
}

func (b BaseBindHandler) bindSendFirstResponse(address string, client net.Conn) {
	err := b.sender.SendSuccessWithTcpPort(client)

	if err != nil {
		b.errorHandler.HandleBindIOError(err, address, client)

		return
	}

	b.bindWait(address, client)

	b.bindManager.Remove(address)
}

func (b BaseBindHandler) bindWait(address string, client net.Conn) {
	deadline := time.Second * time.Duration(b.config.GetBindDeadline())

	host, err := b.bindManager.ReceiveHost(address, deadline)

	if err != nil {
		b.errorHandler.HandleBindManagerReceiveHostError(err, address, client)

		return
	}

	b.bindCheckAddress(address, host, client)
}

func (b BaseBindHandler) bindCheckAddress(address string, host, client net.Conn) {
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

	b.sendSecondResponse(address, addrType, hostAddr, uint16(hostPort), host, client)
}

func (b BaseBindHandler) sendSecondResponse(address string, addrType byte, hostAddress string, hostPort uint16, host, client net.Conn) {
	err := b.sender.SendSuccessWithParameters(addrType, hostAddress, hostPort, client)

	if err != nil {
		b.errorHandler.HandleBindIOErrorWithHost(err, address, client, host)

		return
	}

	err = b.bindManager.SendClient(address, client)

	if err != nil {
		b.errorHandler.HandleBindManagerSendClientError(err, address, client, host)

		return
	}

	b.logger.Bind.Bound(client.RemoteAddr().String(), host.RemoteAddr().String())

	b.streamHandler.ClientToHost(host, client)
}
