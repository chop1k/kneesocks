package server

import (
	"fmt"
	"net"
	"socks/config"
	"socks/logger"
	"socks/protocol/v4a"
)

type V4aHandler interface {
	HandleV4a(request []byte, client net.Conn)
}

type BaseV4aHandler struct {
	parser         v4a.Parser
	config         config.SocksV4aConfig
	logger         logger.SocksV4aLogger
	connectHandler V4aConnectHandler
	bindHandler    V4aBindHandler
	sender         V4aSender
	errorHandler   V4aErrorHandler
}

func NewBaseV4aHandler(
	parser v4a.Parser,
	config config.SocksV4aConfig,
	logger logger.SocksV4aLogger,
	connectHandler V4aConnectHandler,
	bindHandler V4aBindHandler,
	sender V4aSender,
	errorHandler V4aErrorHandler,
) (BaseV4aHandler, error) {
	return BaseV4aHandler{
		parser:         parser,
		config:         config,
		logger:         logger,
		connectHandler: connectHandler,
		bindHandler:    bindHandler,
		sender:         sender,
		errorHandler:   errorHandler,
	}, nil
}

func (b BaseV4aHandler) HandleV4a(request []byte, client net.Conn) {
	chunk, err := b.parser.ParseRequest(request)

	if err != nil {
		b.errorHandler.HandleV4aChunkParseError(err, client)

		return
	}

	address := fmt.Sprintf("%s:%d", chunk.Domain, chunk.DestinationPort)

	if chunk.CommandCode == 1 {
		b.handleConnect(address, client)
	} else if chunk.CommandCode == 2 {
		b.handleBind(address, client)
	} else {
		b.sender.SendFailAndClose(client)
	}
}

func (b BaseV4aHandler) handleConnect(address string, client net.Conn) {
	b.logger.ConnectRequest(client.RemoteAddr().String(), address)

	if !b.config.IsConnectAllowed() {
		b.sender.SendFailAndClose(client)

		b.logger.ConnectNotAllowed(client.RemoteAddr().String(), address)

		return
	}

	b.connectHandler.HandleV4aConnect(address, client)
}

func (b BaseV4aHandler) handleBind(address string, client net.Conn) {
	b.logger.BindRequest(client.RemoteAddr().String(), address)

	if !b.config.IsBindAllowed() {
		b.sender.SendFailAndClose(client)

		b.logger.BindNotAllowed(client.RemoteAddr().String(), address)

		return
	}

	b.bindHandler.HandleV4aBind(address, client)
}
