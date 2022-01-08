package utils

import (
	"errors"
	"net"
	"strconv"
	"strings"
)

var (
	CannotParseIpError        = errors.New("Cannot parse IP address. ")
	CannotParseAddressError   = errors.New("Cannot parse address. ")
	InvalidIPv6AddressError   = errors.New("Invalid IPv6 address. ")
	InvalidAddressLengthError = errors.New("Invalid address length. ")
	InvalidAddressTypeError   = errors.New("Invalid address type. ")
)

type AddressUtils struct {
}

func NewUtils() AddressUtils {
	return AddressUtils{}
}

func (u AddressUtils) DetermineAddressType(addr string) (byte, error) {
	ip := net.ParseIP(addr)

	if ip == nil {
		return 0, CannotParseIpError
	}

	if ip.To4() != nil {
		return 1, nil
	} else if ip.To16() != nil {
		return 4, nil
	} else {
		return 0, CannotParseIpError
	}
}

func (u AddressUtils) ParseAddress(addr string) (string, int, error) {
	chunks := strings.Split(addr, ":")

	if len(chunks) < 2 {
		return "", 0, CannotParseAddressError
	}

	port, err := strconv.Atoi(chunks[len(chunks)-1])

	if err != nil {
		return "", 0, err
	}

	if len(chunks) > 2 {
		ipv6 := strings.Trim(strings.Join(chunks[:len(chunks)-1], ":"), "[]")

		return ipv6, port, nil
	} else {
		return chunks[0], port, nil
	}
}

func (u AddressUtils) ConvertAddress(addrType byte, bytes []byte) (string, error) {
	var addr string

	if addrType == 1 {
		if len(bytes) < 4 {
			return "", InvalidAddressLengthError
		}

		addr = net.IPv4(bytes[0], bytes[1], bytes[2], bytes[3]).String()
	} else if addrType == 3 {
		addr = string(bytes)
	} else if addrType == 4 {
		if len(bytes) < 16 {
			return "", InvalidAddressLengthError
		}

		ipv6 := net.IP{
			bytes[0], bytes[1], bytes[2], bytes[3],
			bytes[4], bytes[5], bytes[6], bytes[7],
			bytes[8], bytes[9], bytes[10], bytes[11],
			bytes[12], bytes[13], bytes[14], bytes[15],
		}.To16()

		if ipv6 == nil {
			return "", InvalidIPv6AddressError
		}

		addr = ipv6.String()
	} else {
		return "", InvalidAddressTypeError
	}

	return addr, nil
}
