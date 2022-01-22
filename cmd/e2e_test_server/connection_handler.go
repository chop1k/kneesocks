package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

type ConnectionHandler struct {
	config Config
	logger Logger
}

func NewConnectionHandler(config Config, logger Logger) (ConnectionHandler, error) {
	return ConnectionHandler{
		config: config,
		logger: logger,
	}, nil
}

func (h ConnectionHandler) HandleConnection(conn net.Conn) {
	buffer := make([]byte, 1)

	i, err := conn.Read(buffer)

	if err != nil {
		_ = conn.Close()

		h.logger.IOError(conn.RemoteAddr().String(), err)

		return
	}

	if i != 1 {
		_ = conn.Close()

		h.logger.IOError(conn.RemoteAddr().String(), err)

		return
	}

	h.handleRequest(buffer[0], conn)
}

func (h ConnectionHandler) handleRequest(picture byte, conn net.Conn) {
	h.logger.PictureRequest(conn.RemoteAddr().String(), picture)

	if picture == 1 {
		h.sendPicture(1, conn)
	} else if picture == 2 {
		h.sendPicture(2, conn)
	} else if picture == 3 {
		h.sendPicture(3, conn)
	} else if picture == 4 {
		h.connectAndSendPicture(conn.RemoteAddr().String(), 4)
	} else if picture == 5 {
		h.connectAndSendPicture(conn.RemoteAddr().String(), 5)
	} else if picture == 6 {
		h.connectAndSendPicture(conn.RemoteAddr().String(), 6)
	} else if picture == 7 {
		h.connectAndSendPicture(conn.RemoteAddr().String(), 7)
	} else if picture == 8 {
		h.connectAndSendPicture(conn.RemoteAddr().String(), 8)
	} else if picture == 9 {
		h.connectAndSendPicture(conn.RemoteAddr().String(), 9)
	}

	_ = conn.Close()
}

func (h ConnectionHandler) sendPicture(picture byte, conn net.Conn) {
	var file *os.File
	var err error

	if picture == 1 {
		file, err = os.Open(h.config.Picture.BigPicturePath)
	} else if picture == 2 {
		file, err = os.Open(h.config.Picture.MiddlePicturePath)
	} else if picture == 3 {
		file, err = os.Open(h.config.Picture.SmallPicturePath)
	} else {
		h.logger.InvalidPicture(conn.RemoteAddr().String(), picture)

		return
	}

	if err != nil {
		h.logger.FileError(err)

		return
	}

	_, _ = io.Copy(conn, file)

	_ = file.Close()
}

func (h ConnectionHandler) connectAndSendPicture(address string, picture byte) {
	var file *os.File
	var err error

	if picture == 4 || picture == 7 {
		file, err = os.Open(h.config.Picture.BigPicturePath)
	} else if picture == 5 || picture == 8 {
		file, err = os.Open(h.config.Picture.MiddlePicturePath)
	} else if picture == 6 || picture == 9 {
		file, err = os.Open(h.config.Picture.SmallPicturePath)
	} else {
		h.logger.InvalidPicture(address, picture)

		return
	}

	if err != nil {
		h.logger.FileError(err)

		return
	}

	var addr string

	if picture == 7 || picture == 8 || picture == 9 {
		addr = fmt.Sprintf("[%s]:%d", h.config.Socks.IPv6, h.config.Socks.Port)
	} else {
		addr = fmt.Sprintf("%s:%d", h.config.Socks.IPv4, h.config.Socks.Port)
	}

	host, dialErr := net.Dial("tcp", addr)

	if dialErr != nil {
		_ = file.Close()

		h.logger.DialError(addr, dialErr)

		return
	}

	_, copyError := io.Copy(host, file)

	h.logger.IOError(addr, copyError)

	_ = host.Close()
	_ = file.Close()
}
