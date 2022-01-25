package protocol

import (
	"errors"
	"net"
)

var (
	InvalidChunkSizeError   = errors.New("Invalid chunk size. ")
	InvalidAddressTypeError = errors.New("Invalid address type. ")
)

type Parser struct {
}

func NewParser() (Parser, error) {
	return Parser{}, nil
}

func (p Parser) ParseRequest(data []byte) (RequestChunk, error) {
	if len(data) < 3 {
		return RequestChunk{}, InvalidChunkSizeError
	}

	if data[2] == 1 {
		if len(data) < 9 {
			return RequestChunk{}, InvalidChunkSizeError
		}
	} else if data[2] == 4 {
		if len(data) < 21 {
			return RequestChunk{}, InvalidChunkSizeError
		}
	} else {
		return RequestChunk{}, InvalidAddressTypeError
	}

	var address net.IP

	if data[2] == 1 {
		address = net.IP{
			data[3], data[4], data[5], data[6],
		}
	} else if data[2] == 4 {
		address = net.IP{
			data[3], data[4], data[5], data[6],
			data[7], data[8], data[9], data[10],
			data[11], data[12], data[13], data[14],
			data[15], data[16], data[17], data[18],
		}.To16()
	}

	var port uint16

	if data[2] == 1 {
		port = uint16(data[7])<<8 | uint16(data[8])
	} else if data[2] == 4 {
		port = uint16(data[19])<<8 | uint16(data[20])
	}

	chunk := RequestChunk{
		Command:     data[0],
		Picture:     data[1],
		AddressType: data[2],
		Address:     address,
		Port:        port,
	}

	return chunk, nil
}
