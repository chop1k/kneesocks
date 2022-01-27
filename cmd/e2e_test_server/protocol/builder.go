package protocol

import (
	"bytes"
	"encoding/binary"
	"errors"
)

var (
	NullAddressError = errors.New("Address is null. ")
)

type Builder struct {
}

func NewBuilder() (Builder, error) {
	return Builder{}, nil
}

func (b Builder) BuildResponse(chunk ResponseChunk) ([]byte, error) {
	return []byte{chunk.Status}, nil
}

func (b Builder) BuildRequest(chunk RequestChunk) ([]byte, error) {
	if chunk.Address == nil {
		return nil, NullAddressError
	}

	buffer := bytes.Buffer{}

	buffer.WriteByte(chunk.Command)
	buffer.WriteByte(chunk.Picture)
	buffer.WriteByte(chunk.AddressType)

	buffer.Write(chunk.Address)

	err := binary.Write(&buffer, binary.BigEndian, chunk.Port)

	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
