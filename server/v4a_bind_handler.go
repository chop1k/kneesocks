package server

import (
	"net"
	"socks/config"
	"socks/logger"
	"socks/utils"
	"time"
)

type V4aBindHandler interface {
	HandleV4aBind(address string, client net.Conn)
}

type BaseV4aBindHandler struct {
	config        config.SocksV4aConfig
	logger        logger.SocksV4aLogger
	streamHandler StreamHandler
	bindManager   BindManager
	utils         utils.AddressUtils
	sender        V4aSender
	whitelist     WhitelistManager
	blacklist     BlacklistManager
	errorHandler  V4aErrorHandler
}

func NewBaseV4aBindHandler(
	config config.SocksV4Config,
	logger logger.SocksV4aLogger,
	streamHandler StreamHandler,
	bindManager BindManager,
	utils utils.AddressUtils,
	sender V4aSender,
	whitelist WhitelistManager,
	blacklist BlacklistManager,
	errorHandler V4aErrorHandler,
) (BaseV4aBindHandler, error) {
	return BaseV4aBindHandler{
		config:        config,
		logger:        logger,
		streamHandler: streamHandler,
		bindManager:   bindManager,
		utils:         utils,
		sender:        sender,
		whitelist:     whitelist,
		blacklist:     blacklist,
		errorHandler:  errorHandler,
	}, nil
}

func (b BaseV4aBindHandler) HandleV4aBind(address string, client net.Conn) {
	whitelisted := b.whitelist.IsWhitelisted(address)

	if whitelisted {
		b.sender.SendFailAndClose(client)

		b.logger.BindNotAllowedByWhitelist(client.RemoteAddr().String(), address)

		return
	}

	blacklisted := b.blacklist.IsBlacklisted(address)

	if blacklisted {
		b.sender.SendFailAndClose(client)

		b.logger.BindNotAllowedByBlacklist(client.RemoteAddr().String(), address)

		return
	}

	b.bind(address, client)
}

func (b BaseV4aBindHandler) bind(address string, client net.Conn) {
	err := b.bindManager.Bind(address)

	if err != nil {
		b.sender.SendFailAndClose(client)

		b.logger.BindFailed(client.RemoteAddr().String(), address)

		return
	}

	b.logger.BindSuccessful(client.RemoteAddr().String(), address)

	b.bindSendFirstResponse(address, client)
}

func (b BaseV4aBindHandler) bindSendFirstResponse(address string, client net.Conn) {
	err := b.sender.SendSuccess(client)

	if err != nil {
		b.errorHandler.HandleV4aBindIOError(err, address, client)

		return
	}

	b.bindWait(address, client)

	b.bindManager.Remove(address)
}

func (b BaseV4aBindHandler) bindWait(address string, client net.Conn) {
	deadline := time.Second * time.Duration(b.config.GetBindDeadline())

	host, err := b.bindManager.ReceiveHost(address, deadline)

	if err != nil {
		b.sender.SendFailAndClose(client)

		b.logger.BindFailed(client.RemoteAddr().String(), address)

		return
	}

	b.bindCheckAddress(address, host, client)
}

func (b BaseV4aBindHandler) bindCheckAddress(address string, host net.Conn, client net.Conn) {
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

	if addrType != 1 {
		b.sender.SendFailAndClose(client)

		_ = host.Close()

		b.logger.BindFailed(client.RemoteAddr().String(), address)

		return
	}

	b.bindSendSecondResponse(address, hostAddr, uint16(hostPort), host, client)
}

func (b BaseV4aBindHandler) bindSendSecondResponse(address string, hostAddr string, hostPort uint16, host net.Conn, client net.Conn) {
	ip := net.ParseIP(hostAddr).To4()

	err := b.sender.SendSuccessWithParameters(ip, hostPort, client)

	if err != nil {
		b.errorHandler.HandleV4aBindIOErrorWithHost(err, address, client, host)

		return
	}

	err = b.bindManager.SendClient(address, client)

	if err != nil {
		_ = client.Close()
		_ = host.Close()

		b.logger.BindFailed(client.RemoteAddr().String(), address)

		return
	}

	b.logger.BindSuccessful(client.RemoteAddr().String(), address)

	b.streamHandler.ClientToHost(host, client)
}
