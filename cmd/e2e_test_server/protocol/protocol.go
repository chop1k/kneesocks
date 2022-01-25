package protocol

import "net"

type RequestChunk struct {
	Command     byte
	Picture     byte
	AddressType byte
	Address     net.IP
	Port        uint16
}

type ResponseChunk struct {
	Status byte
}

type Protocol struct {
	parser  Parser
	builder Builder
}

func NewProtocol(parser Parser, builder Builder) (Protocol, error) {
	return Protocol{parser: parser, builder: builder}, nil
}

func (p Protocol) ReceiveRequest(conn net.Conn) (RequestChunk, error) {
	buffer := make([]byte, 21)

	i, err := conn.Read(buffer)

	if err != nil {
		return RequestChunk{}, err
	}

	return p.parser.ParseRequest(buffer[:i])
}

func (p Protocol) SendResponse(conn net.Conn, status byte) error {
	chunk := ResponseChunk{
		Status: status,
	}

	response, err := p.builder.BuildResponse(chunk)

	if err != nil {
		return err
	}

	_, err = conn.Write(response)

	return err
}
