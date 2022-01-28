package v5

import (
	"bytes"
	"encoding/binary"
	"errors"
	"net"
)

var (
	CannotParseIPError                = errors.New("Cannot parser ip address. ")
	CannotConvertIPToIPv4Error        = errors.New("Cannot convert IP address to ipv4. ")
	CannotConvertIPToIPv6Error        = errors.New("Cannot convert IP address to ipv6. ")
	UnknownAddressTypeError           = errors.New("Unknown address type. ")
	TooManyAuthenticationMethodsError = errors.New("Too many authentication methods, 256 maximum. ")
)

type Builder interface {
	BuildMethodSelection(chunk MethodSelectionChunk) ([]byte, error)
	BuildResponse(chunk ResponseChunk) ([]byte, error)
	BuildMethods(chunk MethodsChunk) ([]byte, error)
	BuildRequest(chunk RequestChunk) ([]byte, error)
}

type BaseBuilder struct {
}

func NewBaseBuilder() (BaseBuilder, error) {
	return BaseBuilder{}, nil
}

func (b BaseBuilder) BuildMethodSelection(chunk MethodSelectionChunk) ([]byte, error) {
	return []byte{chunk.SocksVersion, chunk.Method}, nil
}

func (b BaseBuilder) BuildResponse(chunk ResponseChunk) ([]byte, error) {
	buffer := bytes.Buffer{}

	buffer.WriteByte(chunk.SocksVersion)
	buffer.WriteByte(chunk.ReplyCode)
	buffer.WriteByte(0)
	buffer.WriteByte(chunk.AddressType)

	if chunk.AddressType == 1 {
		ip := net.ParseIP(chunk.Address)

		if ip == nil {
			return nil, CannotParseIPError
		}

		ipv4 := ip.To4()

		if ipv4 == nil {
			return nil, CannotConvertIPToIPv4Error
		}

		buffer.Write(ipv4)
	} else if chunk.AddressType == 4 {
		ip := net.ParseIP(chunk.Address)

		if ip == nil {
			return nil, CannotParseIPError
		}

		ipv6 := ip.To16()

		if ipv6 == nil {
			return nil, CannotConvertIPToIPv6Error
		}

		buffer.Write(ipv6)
	} else if chunk.AddressType == 3 {
		buffer.Write([]byte(chunk.Address))
	} else {
		return nil, UnknownAddressTypeError
	}

	err := binary.Write(&buffer, binary.LittleEndian, chunk.Port)

	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func (b BaseBuilder) BuildMethods(chunk MethodsChunk) ([]byte, error) {
	if len(chunk.Methods) > 256 {
		return nil, TooManyAuthenticationMethodsError
	}

	buffer := bytes.Buffer{}

	buffer.WriteByte(chunk.SocksVersion)
	buffer.WriteByte(byte(len(chunk.Methods)))

	buffer.Write(chunk.Methods)

	return buffer.Bytes(), nil
}

func (b BaseBuilder) BuildRequest(chunk RequestChunk) ([]byte, error) {
	buffer := bytes.Buffer{}

	buffer.WriteByte(chunk.SocksVersion)
	buffer.WriteByte(chunk.CommandCode)
	buffer.WriteByte(0)
	buffer.WriteByte(chunk.AddressType)

	if chunk.AddressType == 1 {
		ip := net.ParseIP(chunk.Address)

		if ip == nil {
			return nil, CannotParseIPError
		}

		ipv4 := ip.To4()

		if ipv4 == nil {
			return nil, CannotConvertIPToIPv4Error
		}

		buffer.Write(ipv4)
	} else if chunk.AddressType == 3 {
		buffer.Write([]byte(chunk.Address))
	} else if chunk.AddressType == 4 {
		ip := net.ParseIP(chunk.Address)

		if ip == nil {
			return nil, CannotParseIPError
		}

		ipv6 := ip.To16()

		if ipv6 == nil {
			return nil, CannotConvertIPToIPv6Error
		}

		buffer.Write(ipv6)
	} else {
		return nil, UnknownAddressTypeError
	}

	err := binary.Write(&buffer, binary.BigEndian, chunk.Port)

	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
