package server

import (
	"net"
	"socks/config"
	"socks/logger"
	"time"
)

type V4ConnectHandler interface {
	HandleV4Connect(address string, client net.Conn)
}

type BaseV4ConnectHandler struct {
	config        config.SocksV4Config
	streamHandler StreamHandler
	logger        logger.SocksV4Logger
	sender        V4Sender
	errorHandler  V4ErrorHandler
}

func NewBaseV4ConnectHandler(
	config config.SocksV4Config,
	streamHandler StreamHandler,
	logger logger.SocksV4Logger,
	sender V4Sender,
	errorHandler V4ErrorHandler,
) (BaseV4ConnectHandler, error) {
	return BaseV4ConnectHandler{
		config:        config,
		streamHandler: streamHandler,
		logger:        logger,
		sender:        sender,
		errorHandler:  errorHandler,
	}, nil
}

func (b BaseV4ConnectHandler) HandleV4Connect(address string, client net.Conn) {
	deadline := time.Second * time.Duration(b.config.GetConnectDeadline())

	host, err := net.DialTimeout("tcp4", address, deadline)

	if err != nil {
		b.errorHandler.HandleV4NetworkError(err, address, client)

		return
	}

	b.connectSendSuccess(address, host, client)
}

func (b BaseV4ConnectHandler) connectSendSuccess(address string, host net.Conn, client net.Conn) {
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
