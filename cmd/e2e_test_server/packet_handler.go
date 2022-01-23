package main

import (
	"net"
	"os"
)

type PacketHandler struct {
	config Config
	logger Logger
}

func NewPacketHandler(config Config, logger Logger) (PacketHandler, error) {
	return PacketHandler{
		config: config,
		logger: logger,
	}, nil
}

func (p PacketHandler) HandlePacket(picture byte, addr net.Addr, conn net.PacketConn) {
	if picture != 255 {
		p.logger.InvalidPicture(addr.String(), picture)

		return
	}

	file, err := os.Open(p.config.Picture.SmallPicturePath)

	if err != nil {
		p.logger.FileError(err)

		return
	}

	buffer := make([]byte, 60000)

	i, readErr := file.Read(buffer)

	if readErr != nil {
		_ = file.Close()

		p.logger.FileError(readErr)

		return
	}

	_, _ = conn.WriteTo(buffer[:i], addr)

	_ = file.Close()
}
