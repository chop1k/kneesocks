package server

import (
	"net"
	"socks/config"
	"socks/logger"
	"socks/utils"
	"time"
)

type V4BindHandler interface {
	HandleV4Bind(address string, client net.Conn)
}

type BaseV4BindHandler struct {
	config        config.SocksV4Config
	logger        logger.SocksV4Logger
	streamHandler StreamHandler
	bindManager   BindManager
	utils         utils.AddressUtils
	sender        V4Sender
	whitelist     WhitelistManager
	blacklist     BlacklistManager
}

func NewBaseV4BindHandler(
	config config.SocksV4Config,
	logger logger.SocksV4Logger,
	streamHandler StreamHandler,
	bindManager BindManager,
	utils utils.AddressUtils,
	sender V4Sender,
	whitelist WhitelistManager,
	blacklist BlacklistManager,
) (BaseV4BindHandler, error) {
	return BaseV4BindHandler{
		config:        config,
		logger:        logger,
		streamHandler: streamHandler,
		bindManager:   bindManager,
		utils:         utils,
		sender:        sender,
		whitelist:     whitelist,
		blacklist:     blacklist,
	}, nil
}

func (b BaseV4BindHandler) HandleV4Bind(address string, client net.Conn) {
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

func (b BaseV4BindHandler) bind(address string, client net.Conn) {
	err := b.bindManager.Bind(address)

	if err != nil {
		b.sender.SendFailAndClose(client)

		b.logger.BindFailed(client.RemoteAddr().String(), address)

		return
	}

	b.logger.BindSuccessful(client.RemoteAddr().String(), address)

	b.bindSendFirstResponse(address, client)
}

func (b BaseV4BindHandler) bindSendFirstResponse(address string, client net.Conn) {
	err := b.sender.SendSuccess(client)

	if err != nil {
		b.sender.SendFailAndClose(client)

		b.logger.BindFailed(client.RemoteAddr().String(), address)

		return
	}

	b.bindWait(address, client)

	b.bindManager.Remove(address)
}

func (b BaseV4BindHandler) bindWait(address string, client net.Conn) {
	deadline := time.Second * time.Duration(b.config.GetBindDeadline())

	host, err := b.bindManager.ReceiveHost(address, deadline)

	if err != nil {
		b.sender.SendFailAndClose(client)

		if err == TimeoutError {
			b.logger.BindTimeout(client.RemoteAddr().String(), address)
		} else {
			b.logger.BindFailed(client.RemoteAddr().String(), address)
		}

		return
	}

	b.bindCheckAddress(address, host, client)
}

func (b BaseV4BindHandler) bindCheckAddress(address string, host net.Conn, client net.Conn) {
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

func (b BaseV4BindHandler) bindSendSecondResponse(address string, hostAddr string, hostPort uint16, host net.Conn, client net.Conn) {
	ip := net.ParseIP(hostAddr).To4()

	err := b.sender.SendSuccessWithParameters(ip, hostPort, client)

	if err != nil {
		b.sender.SendFailAndClose(client)

		_ = host.Close()

		b.logger.BindFailed(client.RemoteAddr().String(), address)

		return
	}

	err = b.bindManager.SendClient(address, client)

	if err != nil {
		_ = client.Close()
		_ = host.Close()

		b.logger.BindFailed(client.RemoteAddr().String(), address)

		return
	}

	b.logger.Bound(client.RemoteAddr().String(), host.RemoteAddr().String())

	b.streamHandler.ClientToHost(client, host)
}
