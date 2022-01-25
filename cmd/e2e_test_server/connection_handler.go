package main

import (
	"fmt"
	"net"
	"socks/cmd/e2e_test_server/protocol"
)

type ConnectionHandler struct {
	config   Config
	logger   Logger
	sender   PictureSender
	protocol protocol.Protocol
}

func NewConnectionHandler(
	config Config,
	logger Logger,
	sender PictureSender,
	protocol protocol.Protocol,
) (ConnectionHandler, error) {
	return ConnectionHandler{
		config:   config,
		logger:   logger,
		sender:   sender,
		protocol: protocol,
	}, nil
}

func (h ConnectionHandler) HandleConnection(conn net.Conn) {
	request, err := h.protocol.ReceiveRequest(conn)

	if err != nil {

	}

	h.handleRequest(request, conn)
}

func (h ConnectionHandler) handleRequest(request protocol.RequestChunk, conn net.Conn) {
	h.logger.PictureRequest(conn.RemoteAddr().String(), request.Picture)

	if request.Picture > 3 || request.Picture < 1 {
		h.logger.InvalidPicture(conn.RemoteAddr().String(), request.Picture)

		err := h.protocol.SendResponse(conn, 254)

		if err != nil {
			h.logger.IOError(conn.RemoteAddr().String(), err)

			_ = conn.Close()

			return
		}

		_ = conn.Close()

		return
	}

	if request.Command == 1 {
		h.handleConnect(request.Picture, conn)
	} else if request.Command == 2 {
		h.handleBind(request.Picture, request.AddressType, request.Address, request.Port, conn)
	} else {

	}
}

func (h ConnectionHandler) handleConnect(picture byte, conn net.Conn) {
	err := h.protocol.SendResponse(conn, 0)

	if err != nil {
		h.logger.IOError(conn.RemoteAddr().String(), err)

		_ = conn.Close()

		return
	}

	h.sender.Send(conn.RemoteAddr().String(), picture, conn)
}

func (h ConnectionHandler) handleBind(picture byte, addressType byte, address net.IP, port uint16, conn net.Conn) {
	lAddr, lErr := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", address.String(), port))

	if lErr != nil {
		h.logger.ResolveError(conn.RemoteAddr().String(), lErr)

		err := h.protocol.SendResponse(conn, 255)

		if err != nil {
			h.logger.IOError(conn.RemoteAddr().String(), err)

			_ = conn.Close()

			return
		}

		_ = conn.Close()

		return
	}

	var socksAddress string

	if addressType == 2 {
		socksAddress = fmt.Sprintf("%s:%d", h.config.Socks.IPv6, h.config.Socks.Port)
	} else {
		socksAddress = fmt.Sprintf("%s:%d", h.config.Socks.IPv4, h.config.Socks.Port)
	}

	rAddr, rErr := net.ResolveTCPAddr("tcp", socksAddress)

	if rErr != nil {
		h.logger.ResolveError(conn.RemoteAddr().String(), rErr)

		err := h.protocol.SendResponse(conn, 255)

		if err != nil {
			h.logger.IOError(conn.RemoteAddr().String(), err)

			_ = conn.Close()

			return
		}

		_ = conn.Close()

		return
	}

	host, dialErr := net.DialTCP("tcp", lAddr, rAddr)

	if dialErr != nil {
		return
	}

	h.bindSendResponse(picture, host, conn)
}

func (h ConnectionHandler) bindSendResponse(picture byte, host net.Conn, conn net.Conn) {
	err := h.protocol.SendResponse(conn, 0)

	if err != nil {
		h.logger.IOError(conn.RemoteAddr().String(), err)

		_ = conn.Close()

		return
	}

	h.sender.Send(conn.RemoteAddr().String(), picture, host)
}
