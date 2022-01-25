package main

import (
	"io"
	"net"
	"os"
)

type PictureSender struct {
	logger Logger
	config Config
}

func NewPictureSender(logger Logger, config Config) (PictureSender, error) {
	return PictureSender{logger: logger, config: config}, nil
}

func (s PictureSender) Send(address string, picture byte, conn net.Conn) {
	var file *os.File
	var err error

	if picture == 1 {
		file, err = os.Open(s.config.Picture.BigPicturePath)
	} else if picture == 2 {
		file, err = os.Open(s.config.Picture.MiddlePicturePath)
	} else if picture == 3 {
		file, err = os.Open(s.config.Picture.SmallPicturePath)
	} else {
		s.logger.InvalidPicture(address, picture)

		return
	}

	if err != nil {
		s.logger.FileError(err)

		return
	}

	_, _ = io.Copy(conn, file)

	_ = file.Close()
}
