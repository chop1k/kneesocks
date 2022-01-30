package v4a

import (
	"net"
	"socks/config"
	"socks/logger"
	"socks/managers"
	"socks/transfer"
	"socks/utils"
	"time"
)

type BindHandler interface {
	HandleBind(address string, client net.Conn)
}

type BaseBindHandler struct {
	config        config.SocksV4aConfig
	logger        logger.SocksV4aLogger
	streamHandler transfer.StreamHandler
	bindManager   managers.BindManager
	utils         utils.AddressUtils
	sender        Sender
	whitelist     Whitelist
	blacklist     Blacklist
	errorHandler  ErrorHandler
}

func NewBaseBindHandler(
	config config.SocksV4Config,
	logger logger.SocksV4aLogger,
	streamHandler transfer.StreamHandler,
	bindManager managers.BindManager,
	utils utils.AddressUtils,
	sender Sender,
	whitelist Whitelist,
	blacklist Blacklist,
	errorHandler ErrorHandler,
) (BaseBindHandler, error) {
	return BaseBindHandler{
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

func (b BaseBindHandler) HandleBind(address string, client net.Conn) {
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

func (b BaseBindHandler) bind(address string, client net.Conn) {
	err := b.bindManager.Bind(address)

	if err != nil {
		b.errorHandler.HandleBindManagerBindError(err, address, client)

		return
	}

	b.logger.BindSuccessful(client.RemoteAddr().String(), address)

	b.bindSendFirstResponse(address, client)
}

func (b BaseBindHandler) bindSendFirstResponse(address string, client net.Conn) {
	err := b.sender.SendSuccess(client)

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

func (b BaseBindHandler) bindCheckAddress(address string, host net.Conn, client net.Conn) {
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

	if addrType != 1 {
		b.errorHandler.HandleInvalidAddressTypeError(address, client, host)

		return
	}

	b.bindSendSecondResponse(address, hostAddr, uint16(hostPort), host, client)
}

func (b BaseBindHandler) bindSendSecondResponse(address string, hostAddr string, hostPort uint16, host net.Conn, client net.Conn) {
	ip := net.ParseIP(hostAddr).To4()

	err := b.sender.SendSuccessWithParameters(ip, hostPort, client)

	if err != nil {
		b.errorHandler.HandleBindIOErrorWithHost(err, address, client, host)

		return
	}

	err = b.bindManager.SendClient(address, client)

	if err != nil {
		b.errorHandler.HandleBindManagerSendClientError(err, address, client, host)

		return
	}

	b.logger.BindSuccessful(client.RemoteAddr().String(), address)

	b.streamHandler.ClientToHost(host, client)
}
