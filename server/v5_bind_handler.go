package server

import (
	"net"
	"socks/config"
	"socks/logger"
	"socks/utils"
	"time"
)

type V5BindHandler interface {
	HandleV5Bind(address string, client net.Conn)
}

type BaseV5BindHandler struct {
	bindManager   BindManager
	config        config.SocksV5Config
	streamHandler StreamHandler
	utils         utils.AddressUtils
	logger        logger.SocksV5Logger
	sender        V5Sender
	whitelist     WhitelistManager
	blacklist     BlacklistManager
}

func NewBaseV5BindHandler(
	bindManager BindManager,
	config config.SocksV5Config,
	streamHandler StreamHandler,
	utils utils.AddressUtils,
	logger logger.SocksV5Logger,
	sender V5Sender,
	whitelist WhitelistManager,
	blacklist BlacklistManager,
) (BaseV5BindHandler, error) {
	return BaseV5BindHandler{
		bindManager:   bindManager,
		config:        config,
		streamHandler: streamHandler,
		utils:         utils,
		logger:        logger,
		sender:        sender,
		whitelist:     whitelist,
		blacklist:     blacklist,
	}, nil
}

func (b BaseV5BindHandler) HandleV5Bind(address string, client net.Conn) {
	whitelisted := b.whitelist.IsWhitelisted(address)

	if whitelisted {
		b.sender.SendConnectionNotAllowedAndClose(client)

		b.logger.BindNotAllowedByWhitelist(client.RemoteAddr().String(), address)

		return
	}

	blacklisted := b.blacklist.IsBlacklisted(address)

	if blacklisted {
		b.sender.SendConnectionNotAllowedAndClose(client)

		b.logger.BindNotAllowedByBlacklist(client.RemoteAddr().String(), address)

		return
	}

	b.bind(address, client)
}

func (b BaseV5BindHandler) bind(address string, client net.Conn) {
	err := b.bindManager.Bind(address)

	if err != nil {
		b.sender.SendConnectionNotAllowedAndClose(client)

		b.logger.BindFailed(client.RemoteAddr().String(), address)

		return
	}

	b.logger.BindSuccessful(client.RemoteAddr().String(), address)

	b.bindSendFirstResponse(address, client)
}

func (b BaseV5BindHandler) bindSendFirstResponse(address string, client net.Conn) {
	err := b.sender.SendSuccessWithTcpPort(client)

	if err != nil {
		_ = client.Close()

		b.logger.BindFailed(client.RemoteAddr().String(), address)

		return
	}

	b.bindWait(address, client)

	b.bindManager.Remove(address)
}

func (b BaseV5BindHandler) bindWait(address string, client net.Conn) {
	deadline := time.Second * time.Duration(b.config.GetBindDeadline())

	host, err := b.bindManager.ReceiveHost(address, deadline)

	if err != nil {
		_ = client.Close()

		b.logger.BindFailed(client.RemoteAddr().String(), address)

		return
	}

	b.bindCheckAddress(address, host, client)
}

func (b BaseV5BindHandler) bindCheckAddress(address string, host, client net.Conn) {
	hostAddr, hostPort, parseErr := b.utils.ParseAddress(host.RemoteAddr().String())

	if parseErr != nil {
		b.sender.SendFailAndClose(client)

		_ = host.Close()

		b.logger.BindFailed(client.RemoteAddr().String(), address)

		return
	}

	addrType, determineErr := b.utils.DetermineAddressType(hostAddr)

	if determineErr != nil {
		b.sender.SendFailAndClose(client)

		_ = host.Close()

		b.logger.BindFailed(client.RemoteAddr().String(), address)

		return
	}

	b.sendSecondResponse(address, addrType, hostAddr, uint16(hostPort), host, client)
}

func (b BaseV5BindHandler) sendSecondResponse(address string, addrType byte, hostAddress string, hostPort uint16, host, client net.Conn) {
	err := b.sender.SendSuccessWithParameters(addrType, hostAddress, hostPort, client)

	if err != nil {
		b.sender.SendFailAndClose(client)

		_ = host.Close()

		b.logger.BindFailed(client.RemoteAddr().String(), address)

		return
	}

	err = b.bindManager.SendClient(address, client)

	if err != nil {
		b.sender.SendFailAndClose(client)

		_ = host.Close()

		b.logger.BindFailed(client.RemoteAddr().String(), address)

		return
	}

	b.logger.Bound(client.RemoteAddr().String(), host.RemoteAddr().String())

	b.streamHandler.ClientToHost(host, client)
}
