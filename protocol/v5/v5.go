package v5

import (
	"net"
)

type MethodsChunk struct {
	SocksVersion byte
	Methods      []byte
}

type MethodSelectionChunk struct {
	SocksVersion byte
	Method       byte
}

type RequestChunk struct {
	SocksVersion byte
	CommandCode  byte
	AddressType  byte
	Address      string
	Port         uint16
}

type ResponseChunk struct {
	SocksVersion byte
	ReplyCode    byte
	AddressType  byte
	Address      string
	Port         uint16
}

type UdpRequest struct {
	Fragment    byte
	AddressType byte
	Address     string
	Port        uint16
	Data        []byte
}

type Protocol struct {
	builder Builder
	parser  Parser
}

func NewProtocol(builder Builder, parser Parser) Protocol {
	return Protocol{
		builder: builder,
		parser:  parser,
	}
}

func (p Protocol) ReceiveRequest(client net.Conn) (RequestChunk, error) {
	request := make([]byte, 1024)

	i, err := client.Read(request)

	if err != nil {
		_ = client.Close()

		return RequestChunk{}, err
	}

	return p.parser.ParseRequest(request[:i])
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

func (p Protocol) responseWithCode(code byte, addrType byte, addr string, port uint16, client net.Conn) error {
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

func (p Protocol) ResponseWithSuccess(addrType byte, addr string, port uint16, client net.Conn) error {
	return p.responseWithCode(0, addrType, addr, port, client)
}

func (p Protocol) ResponseWithFail(addrType byte, addr string, port uint16, client net.Conn) error {
	return p.responseWithCode(1, addrType, addr, port, client)
}

func (p Protocol) ResponseWithNotAllowed(addrType byte, addr string, port uint16, client net.Conn) error {
	return p.responseWithCode(2, addrType, addr, port, client)
}

func (p Protocol) ResponseWithNetworkUnreachable(addrType byte, addr string, port uint16, client net.Conn) error {
	return p.responseWithCode(3, addrType, addr, port, client)
}

func (p Protocol) ResponseWithHostUnreachable(addrType byte, addr string, port uint16, client net.Conn) error {
	return p.responseWithCode(4, addrType, addr, port, client)
}

func (p Protocol) ResponseWithConnectionRefused(addrType byte, addr string, port uint16, client net.Conn) error {
	return p.responseWithCode(5, addrType, addr, port, client)
}

func (p Protocol) ResponseWithCommandNotSupported(addrType byte, addr string, port uint16, client net.Conn) error {
	return p.responseWithCode(7, addrType, addr, port, client)
}

func (p Protocol) ResponseWithAddressNotSupported(addrType byte, addr string, port uint16, client net.Conn) error {
	return p.responseWithCode(8, addrType, addr, port, client)
}
