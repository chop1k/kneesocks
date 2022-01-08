package v4

import (
	"bytes"
	"encoding/binary"
	"errors"
)

var (
	DestinationIpIsNullError = errors.New("Destination ip is null. ")
)

type Builder interface {
	BuildResponse(chunk ResponseChunk) ([]byte, error)
}

type BaseBuilder struct {
}

func NewBaseBuilder() BaseBuilder {
	return BaseBuilder{}
}

func (b BaseBuilder) BuildResponse(chunk ResponseChunk) ([]byte, error) {
	if chunk.DestinationIp == nil {
		return nil, DestinationIpIsNullError
	}

	buffer := bytes.Buffer{}

	buffer.WriteByte(chunk.SocksVersion)
	buffer.WriteByte(chunk.CommandCode)

	err := binary.Write(&buffer, binary.LittleEndian, chunk.DestinationPort)

	if err != nil {
		return nil, err
	}

	buffer.Write(chunk.DestinationIp)

	return buffer.Bytes(), nil
}
