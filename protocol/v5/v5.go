package v5

import (
	"net"
)

type MethodsChunk struct {
	SocksVersion byte
	Methods []byte
}

type MethodSelectionChunk struct {
	SocksVersion byte
	Method byte
}

type RequestChunk struct {
	SocksVersion byte
	CommandCode byte
	AddressType byte
	Address string
	Port uint16
}

type ResponseChunk struct {
	SocksVersion byte
	ReplyCode byte
	AddressType byte
	Address string
	Port uint16
}

type UdpRequest struct {
	Fragment byte
	AddressType byte
	Address string
	Port uint16
	Data []byte
}

type Protocol struct {
	builder Builder
}

func NewProtocol(builder Builder) Protocol {
	return Protocol{
		builder: builder,
	}
}

func (p Protocol) SelectMethod(method byte, client net.Conn) error {
	selection := MethodSelectionChunk{
		SocksVersion: 5,
		Method:       method,
	}

	response, err := p.builder.BuildMethodSelection(selection)

	if err != nil {
		return err
	}

	_, err = client.Write(response)

	return err
}

func (p Protocol) ResponseWithCode(code byte, addrType byte, addr string, port uint16, client net.Conn) error {
	chunk := ResponseChunk{
		SocksVersion: 5,
		ReplyCode:    code,
		AddressType:  addrType,
		Address:      addr,
		Port:         port,
	}

	response, err := p.builder.BuildResponse(chunk)

	if err != nil {
		return err
	}

	_, err = client.Write(response)

	if err != nil {
		return err
	}

	return nil
}
