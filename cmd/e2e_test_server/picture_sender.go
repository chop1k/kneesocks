package main

import (
	"errors"
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

func (s PictureSender) Send(address string, picture byte, conn net.Conn) error {
	path, err := s.getAddress(picture, address)

	if err != nil {
		_ = conn.Close()

		return err
	}

	return s.send(address, path, picture, conn)
}

func (s PictureSender) getAddress(picture byte, address string) (string, error) {
	if picture == 1 {
		return s.config.Picture.BigPicturePath, nil
	} else if picture == 2 {
		return s.config.Picture.MiddlePicturePath, nil
	} else if picture == 3 {
		return s.config.Picture.SmallPicturePath, nil
	} else {
		s.logger.InvalidPicture(address, picture)

		return "", errors.New("Invalid picture. ")
	}
}

func (s PictureSender) send(address string, path string, picture byte, conn net.Conn) error {
	file, err := os.Open(path)

	if err != nil {
		s.logger.FileError(err)

		_ = conn.Close()

		return err
	}

	s.logger.PictureOpened(address, picture)

	for {
		buffer := make([]byte, 512)

		i, err := file.Read(buffer)

		if err != nil {
			break
		}

		_, err = conn.Write(buffer[:i])

		if err != nil {
			s.logger.IOError(address, err)

			break
		}
	}

	s.logger.PictureSent(address, picture)

	return s.clean(file, conn)
}

func (s PictureSender) clean(file *os.File, conn net.Conn) error {
	_ = conn.Close()

	_ = file.Close()

	return nil
}
