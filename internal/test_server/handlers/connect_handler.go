package handlers

import (
	"net"
	"socks/internal/test_server/config"
	"socks/internal/test_server/logger"
	"socks/internal/test_server/picture"
)

type ConnectHandler struct {
	config config.Config
	logger logger.Logger
	sender picture.Sender
}

func NewConnectHandler(
	config config.Config,
	logger logger.Logger,
	sender picture.Sender,
) (ConnectHandler, error) {
	return ConnectHandler{
		config: config,
		logger: logger,
		sender: sender,
	}, nil
}

func (h ConnectHandler) HandleConnect(conn net.Conn) {
	picture := make([]byte, 1)

	_, err := conn.Read(picture)

	if err != nil {
		h.logger.IOError(conn.RemoteAddr().String(), err)

		_ = conn.Close()

		return
	}

	h.handleConnect(picture[0], conn)
}

func (h ConnectHandler) handleConnect(picture byte, conn net.Conn) {
	err := h.sender.Send(conn.RemoteAddr().String(), picture, conn)

	if err != nil {
		_ = conn.Close()

		return
	}
}
