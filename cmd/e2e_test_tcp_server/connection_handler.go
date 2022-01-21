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
		h.connectAndSendPicture(4)
	} else if picture == 5 {
		h.connectAndSendPicture(5)
	} else if picture == 6 {
		h.connectAndSendPicture(6)
	} else if picture == 7 {
		h.connectAndSendPicture(7)
	} else if picture == 8 {
		h.connectAndSendPicture(8)
	} else if picture == 9 {
		h.connectAndSendPicture(9)
	}

	_ = conn.Close()
}

func (h ConnectionHandler) sendPicture(picture byte, writer io.Writer) {
	var file *os.File
	var err error

	if picture == 1 {
		file, err = os.Open(h.config.BigPicturePath)
	} else if picture == 2 {
		file, err = os.Open(h.config.MiddlePicturePath)
	} else if picture == 3 {
		file, err = os.Open(h.config.SmallPicturePath)
	} else {
		return
	}

	if err != nil {
		h.logger.FileError(err)

		return
	}

	_, _ = io.Copy(writer, file)

	_ = file.Close()
}

func (h ConnectionHandler) connectAndSendPicture(picture byte) {
	var file *os.File
	var err error

	if picture == 4 || picture == 7 {
		file, err = os.Open(h.config.BigPicturePath)
	} else if picture == 5 || picture == 8 {
		file, err = os.Open(h.config.MiddlePicturePath)
	} else if picture == 6 || picture == 9 {
		file, err = os.Open(h.config.SmallPicturePath)
	} else {
		return
	}

	if err != nil {
		h.logger.FileError(err)

		return
	}

	var address string

	if picture == 7 || picture == 8 || picture == 9 {
		address = fmt.Sprintf("[%s]:%d", "::1", bindPort)
	} else {
		address = fmt.Sprintf("%s:%d", "127.0.0.1", bindPort)
	}

	lAddr, lErr := net.ResolveTCPAddr("tcp", address)

	if lErr != nil {
		return
	}

	rAddr, rErr := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", socksHost, socksPort))

	if rErr != nil {
		return
	}

	host, dialErr := net.DialTCP("tcp", lAddr, rAddr)

	if dialErr != nil {
		_ = file.Close()

		return
	}

	for {
		buf := make([]byte, 512)

		i, err := file.Read(buf)

		if err != nil {
			break
		}

		_, writeErr := host.Write(buf[:i])

		if writeErr != nil {
			break
		}
	}

	_ = host.Close()
	_ = file.Close()
}
