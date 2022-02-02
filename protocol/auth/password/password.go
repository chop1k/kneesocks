package password

import "net"

type RequestChunk struct {
	Version  byte
	Name     string
	Password string
}

type ResponseChunk struct {
	Version byte
	Status  byte
}

type Password struct {
	parser  Parser
	builder Builder
}

func NewPassword(parser Parser, builder Builder) Password {
	return Password{parser: parser, builder: builder}
}

func (p Password) ResponseWith(code byte, client net.Conn) error {
	bytes, err := p.builder.BuildResponse(ResponseChunk{
		Version: 1,
		Status:  code,
	})

	if err != nil {
		return err
	}

	_, err = client.Write(bytes)

	return err
}
