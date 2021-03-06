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
	DomainTooLongError                = errors.New("Domain too long. ")
)

type Builder struct {
}

func NewBuilder() (Builder, error) {
	return Builder{}, nil
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
		if len(chunk.Address) > 256 {
			return nil, DomainTooLongError
		}

		buffer.WriteByte(byte(len(chunk.Address)))
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

func (b Builder) BuildUdpRequest(chunk UdpRequest) ([]byte, error) {
	buffer := bytes.Buffer{}

	buffer.WriteByte(0)
	buffer.WriteByte(0)
	buffer.WriteByte(chunk.Fragment)
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
		if len(chunk.Address) > 256 {
			return nil, DomainTooLongError
		}

		buffer.WriteByte(byte(len(chunk.Address)))
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

	buffer.Write(chunk.Data)

	return buffer.Bytes(), nil
}
