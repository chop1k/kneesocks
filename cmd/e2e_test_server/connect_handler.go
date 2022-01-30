package main

import (
	"net"
)

type ConnectHandler struct {
	config Config
	logger Logger
	sender PictureSender
}

func NewConnectHandler(
	config Config,
	logger Logger,
	sender PictureSender,
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
