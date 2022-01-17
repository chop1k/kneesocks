package v4

import (
	"net"
)

type RequestChunk struct {
	SocksVersion    uint8
	CommandCode     uint8
	DestinationPort uint16
	DestinationIp   net.IP
	UserId          string
}

type ResponseChunk struct {
	SocksVersion    uint8
	CommandCode     uint8
	DestinationPort uint16
	DestinationIp   net.IP
}

type Protocol struct {
	builder Builder
}

func NewProtocol(
	builder Builder,
) Protocol {
	return Protocol{
		builder: builder,
	}
}

func (p Protocol) responseWithCode(code byte, port uint16, ip net.IP, client net.Conn) error {
	chunk := ResponseChunk{
		SocksVersion:    0, // НЕ МЕНЯЙ ВСЁ НАЕБНЕТСЯ СПЕЦИФИКАЦИЯ ГОВНО
		CommandCode:     code,
		DestinationPort: port,
		DestinationIp:   ip,
	}

	response, buildErr := p.builder.BuildResponse(chunk)

	if buildErr != nil {
		return buildErr
	}

	_, err := client.Write(response)

	return err
}

func (p Protocol) ResponseWithSuccess(port uint16, ip net.IP, client net.Conn) error {
	return p.responseWithCode(90, port, ip, client)
}

func (p Protocol) ResponseWithFail(port uint16, ip net.IP, client net.Conn) error {
	return p.responseWithCode(91, port, ip, client)
}
