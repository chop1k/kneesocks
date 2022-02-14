package v5

import (
	"errors"
	"socks/pkg/utils"
)

var (
	InvalidSocksVersionError    = errors.New("Invalid socks version error. ")
	InvalidChunkSizeError       = errors.New("Invalid chunk size error. ")
	InvalidNumberOfMethodsError = errors.New("Invalid number of methods. ")
	InvalidCommandCodeError     = errors.New("Invalid command code. ")
	InvalidReservedByteError    = errors.New("Invalid rsv byte. ")
	InvalidFragmentByteError    = errors.New("Invalid frag byte. ")
)

type Parser struct {
	utils utils.AddressUtils
}

func NewParser(utils utils.AddressUtils) Parser {
	return Parser{utils: utils}
}

func (b Parser) ParseMethods(bytes []byte) (MethodsChunk, error) {
	length := len(bytes)

	if length < 3 {
		return MethodsChunk{}, InvalidChunkSizeError
	}

	if bytes[0] != 5 {
		return MethodsChunk{}, InvalidSocksVersionError
	}

	numberOfMethods := bytes[1] + 2

	if numberOfMethods <= 0 {
		return MethodsChunk{}, InvalidNumberOfMethodsError
	}

	if int(numberOfMethods) > length {
		return MethodsChunk{}, InvalidChunkSizeError
	}

	methods := bytes[2:numberOfMethods]

	if len(methods) != int(numberOfMethods-2) {
		return MethodsChunk{}, InvalidNumberOfMethodsError
	}

	return MethodsChunk{
		SocksVersion: bytes[0],
		Methods:      methods,
	}, nil
}

func (b Parser) ParseRequest(bytes []byte) (RequestChunk, error) {
	length := len(bytes)

	if length < 10 {
		return RequestChunk{}, InvalidChunkSizeError
	}

	if bytes[0] != 5 {
		return RequestChunk{}, InvalidSocksVersionError
	}

	if bytes[1] != 1 && bytes[1] != 2 && bytes[1] != 3 {
		return RequestChunk{}, InvalidCommandCodeError
	}

	if bytes[2] != 0 {
		return RequestChunk{}, InvalidReservedByteError
	}

	var addrEnd int
	var addrStart int

	if bytes[3] == 1 {
		addrStart = 4
		addrEnd = 8
	} else if bytes[3] == 3 {
		addrStart = 5
		addrEnd = int(bytes[4] + 5)

		if length < addrEnd+2 {
			return RequestChunk{}, InvalidChunkSizeError
		}
	} else if bytes[3] == 4 {
		addrStart = 4
		addrEnd = 20

		if length < 22 {
			return RequestChunk{}, InvalidChunkSizeError
		}
	} else {
		return RequestChunk{}, UnknownAddressTypeError
	}

	addr, err := b.utils.ConvertAddress(bytes[3], bytes[addrStart:addrEnd])

	if err != nil {
		return RequestChunk{}, err
	}

	return RequestChunk{
		SocksVersion: bytes[0],
		CommandCode:  bytes[1],
		AddressType:  bytes[3],
		Address:      addr,
		Port:         uint16(bytes[addrEnd])<<8 | uint16(bytes[addrEnd+1]),
	}, nil
}

func (b Parser) ParseUdpRequest(bytes []byte) (UdpRequest, error) {
	length := len(bytes)

	if length < 10 {
		return UdpRequest{}, InvalidChunkSizeError
	}

	if bytes[0] != 0 || bytes[1] != 0 {
		return UdpRequest{}, InvalidReservedByteError
	}

	if bytes[2] < 0 || bytes[2] > 127 {
		return UdpRequest{}, InvalidFragmentByteError
	}

	var addrEnd int
	var addrStart int

	if bytes[3] == 1 {
		addrStart = 4
		addrEnd = 8
	} else if bytes[3] == 3 {
		addrStart = 5
		addrEnd = int(bytes[4] + 5)

		if length < addrEnd+2 {
			return UdpRequest{}, InvalidChunkSizeError
		}
	} else if bytes[3] == 4 {
		addrStart = 4
		addrEnd = 20

		if length < 22 {
			return UdpRequest{}, InvalidChunkSizeError
		}
	} else {
		return UdpRequest{}, UnknownAddressTypeError
	}

	addr, err := b.utils.ConvertAddress(bytes[3], bytes[addrStart:addrEnd])

	if err != nil {
		return UdpRequest{}, err
	}

	return UdpRequest{
		Fragment:    bytes[2],
		AddressType: bytes[3],
		Address:     addr,
		Port:        uint16(bytes[addrEnd])<<8 | uint16(bytes[addrEnd+1]),
		Data:        bytes[addrEnd+2:],
	}, nil
}
