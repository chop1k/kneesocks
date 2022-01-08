package v4a

import (
	"errors"
	"net"
)

var (
	InvalidSocksVersionError = errors.New("Invalid socks version. ")
	InvalidChunkSizeError    = errors.New("Invalid chunk size. ")
)

type Parser interface {
	ParseRequest(bytes []byte) (RequestChunk, error)
}

type BaseParser struct {
}

func NewBaseParser() BaseParser {
	return BaseParser{}
}

func (b BaseParser) ParseRequest(bytes []byte) (RequestChunk, error) {
	if len(bytes) < 9 {
		return RequestChunk{}, InvalidChunkSizeError
	}

	if bytes[0] != 4 {
		return RequestChunk{}, InvalidSocksVersionError
	}

	return RequestChunk{
		SocksVersion: bytes[0],
		CommandCode: bytes[1],
		DestinationPort: uint16(bytes[2]) << 8 | uint16(bytes[3]),
		DestinationIp: net.IP{bytes[4], bytes[5], bytes[6], bytes[7]},
		Domain: string(bytes[8:len(bytes) - 1]),
	}, nil
}
