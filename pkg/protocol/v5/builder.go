package v5

import (
	"bytes"
	"encoding/binary"
	"errors"
	"net"
)

var (
	CannotParseIPError                = errors.New("cannot parser ip address")
	CannotConvertIPToIPv4Error        = errors.New("cannot convert IP address to ipv4")
	CannotConvertIPToIPv6Error        = errors.New("cannot convert IP address to ipv6")
	UnknownAddressTypeError           = errors.New("unknown address type")
	TooManyAuthenticationMethodsError = errors.New("too many authentication methods, 256 maximum")
	DomainTooLongError                = errors.New("domain too long")
)

type Builder struct {
}

func NewBuilder() Builder {
	return Builder{}
}

func (b Builder) BuildMethodSelection(chunk MethodSelectionChunk) ([]byte, error) {
	return []byte{chunk.SocksVersion, chunk.Method}, nil
}

func (b Builder) BuildResponse(chunk ResponseChunk) ([]byte, error) {
	buffer := bytes.Buffer{}

	buffer.WriteByte(chunk.SocksVersion)
	buffer.WriteByte(chunk.ReplyCode)
	buffer.WriteByte(0)
	buffer.WriteByte(chunk.AddressType)

	switch chunk.AddressType {
	case 1:
		ip := net.ParseIP(chunk.Address)

		if ip == nil {
			return nil, CannotParseIPError
		}

		ipv4 := ip.To4()

		if ipv4 == nil {
			return nil, CannotConvertIPToIPv4Error
		}

		buffer.Write(ipv4)
	case 4:
		ip := net.ParseIP(chunk.Address)

		if ip == nil {
			return nil, CannotParseIPError
		}

		ipv6 := ip.To16()

		if ipv6 == nil {
			return nil, CannotConvertIPToIPv6Error
		}

		buffer.Write(ipv6)
	case 3:
		buffer.Write([]byte(chunk.Address))
	default:
		return nil, UnknownAddressTypeError
	}

	err := binary.Write(&buffer, binary.LittleEndian, chunk.Port)

	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func (b Builder) BuildMethods(chunk MethodsChunk) ([]byte, error) {
	if len(chunk.Methods) > 256 {
		return nil, TooManyAuthenticationMethodsError
	}

	buffer := bytes.Buffer{}

	buffer.WriteByte(chunk.SocksVersion)
	buffer.WriteByte(byte(len(chunk.Methods)))

	buffer.Write(chunk.Methods)

	return buffer.Bytes(), nil
}

func (b Builder) BuildRequest(chunk RequestChunk) ([]byte, error) {
	buffer := bytes.Buffer{}

	buffer.WriteByte(chunk.SocksVersion)
	buffer.WriteByte(chunk.CommandCode)
	buffer.WriteByte(0)
	buffer.WriteByte(chunk.AddressType)

	switch chunk.AddressType {
	case 1:
		ip := net.ParseIP(chunk.Address)

		if ip == nil {
			return nil, CannotParseIPError
		}

		ipv4 := ip.To4()

		if ipv4 == nil {
			return nil, CannotConvertIPToIPv4Error
		}

		buffer.Write(ipv4)
	case 3:
		if len(chunk.Address) > 256 {
			return nil, DomainTooLongError
		}

		buffer.WriteByte(byte(len(chunk.Address)))
		buffer.Write([]byte(chunk.Address))
	case 4:
		ip := net.ParseIP(chunk.Address)

		if ip == nil {
			return nil, CannotParseIPError
		}

		ipv6 := ip.To16()

		if ipv6 == nil {
			return nil, CannotConvertIPToIPv6Error
		}

		buffer.Write(ipv6)
	default:
		return nil, UnknownAddressTypeError
	}

	err := binary.Write(&buffer, binary.BigEndian, chunk.Port)

	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func (b Builder) BuildUdpRequest(chunk UdpRequest) ([]byte, error) {
	buffer := bytes.Buffer{}

	buffer.WriteByte(0)
	buffer.WriteByte(0)
	buffer.WriteByte(chunk.Fragment)
	buffer.WriteByte(chunk.AddressType)

	switch chunk.AddressType {
	case 1:
		ip := net.ParseIP(chunk.Address)

		if ip == nil {
			return nil, CannotParseIPError
		}

		ipv4 := ip.To4()

		if ipv4 == nil {
			return nil, CannotConvertIPToIPv4Error
		}

		buffer.Write(ipv4)
	case 3:
		if len(chunk.Address) > 256 {
			return nil, DomainTooLongError
		}

		buffer.WriteByte(byte(len(chunk.Address)))
		buffer.Write([]byte(chunk.Address))
	case 4:
		ip := net.ParseIP(chunk.Address)

		if ip == nil {
			return nil, CannotParseIPError
		}

		ipv6 := ip.To16()

		if ipv6 == nil {
			return nil, CannotConvertIPToIPv6Error
		}

		buffer.Write(ipv6)
	default:
		return nil, UnknownAddressTypeError
	}

	err := binary.Write(&buffer, binary.BigEndian, chunk.Port)

	if err != nil {
		return nil, err
	}

	buffer.Write(chunk.Data)

	return buffer.Bytes(), nil
}
