package handlers

import (
	"fmt"
	"net"
	"socks/internal/test_server/config"
	"socks/internal/test_server/logger"
	"socks/internal/test_server/picture"
)

type BindHandler struct {
	config config.Config
	logger logger.Logger
	sender picture.Sender
}

func NewBindHandler(config config.Config, logger logger.Logger, sender picture.Sender) (BindHandler, error) {
	return BindHandler{config: config, logger: logger, sender: sender}, nil
}

func (h BindHandler) HandleBind(conn net.Conn) {
	buffer := make([]byte, 4)

	_, err := conn.Read(buffer)

	if err != nil {
		h.logger.IOError(conn.RemoteAddr().String(), err)

		_ = conn.Close()

		return
	}

	port := uint16(buffer[2])<<8 | uint16(buffer[3])

	h.resolveLAddr(buffer[1], port, buffer[0], conn)
}

func (h BindHandler) resolveLAddr(addressType byte, port uint16, picture byte, conn net.Conn) {
	var selfAddress string

	if addressType == 4 {
		selfAddress = fmt.Sprintf("[%s]:%d", h.config.Socks.IPv6, port)
	} else {
		selfAddress = fmt.Sprintf("%s:%d", h.config.Socks.IPv4, port)
	}

	lAddr, err := net.ResolveTCPAddr("tcp", selfAddress)

	if err != nil {
		h.logger.ResolveError(conn.RemoteAddr().String(), err)

		_ = conn.Close()

		return
	}

	h.resolveRAddr(addressType, lAddr, picture, conn)
}

func (h BindHandler) resolveRAddr(addressType byte, lAddr *net.TCPAddr, picture byte, conn net.Conn) {
	var socksAddress string

	if addressType == 4 {
		socksAddress = fmt.Sprintf("[%s]:%d", h.config.Socks.IPv6, h.config.Socks.Port)
	} else {
		socksAddress = fmt.Sprintf("%s:%d", h.config.Socks.IPv4, h.config.Socks.Port)
	}

	rAddr, err := net.ResolveTCPAddr("tcp", socksAddress)

	if err != nil {
		h.logger.ResolveError(conn.RemoteAddr().String(), err)

		_ = conn.Close()

		return
	}

	h.bind(lAddr, rAddr, picture, conn)
}

func (h BindHandler) bind(lAddr *net.TCPAddr, rAddr *net.TCPAddr, picture byte, conn net.Conn) {
	host, dialErr := net.DialTCP("tcp", lAddr, rAddr)

	if dialErr != nil {
		h.logger.DialError(conn.RemoteAddr().String(), dialErr)

		_ = conn.Close()

		return
	}

	err := h.sender.Send(conn.RemoteAddr().String(), picture, host)

	if err != nil {
		_ = conn.Close()
		_ = host.Close()

		return
	}

	_ = conn.Close()
	_ = host.Close()
}
