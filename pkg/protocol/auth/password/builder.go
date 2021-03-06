package password

import (
	"bytes"
	"errors"
)

var (
	InvalidVersionError  = errors.New("Invalid version, must be 1. ")
	NameTooLongError     = errors.New("Name too long, 256 maximum. ")
	PasswordTooLongError = errors.New("Password too long, 256 maximum. ")
)

type Builder struct {
}

func NewBuilder() (Builder, error) {
	return Builder{}, nil
}

func (b Builder) BuildResponse(chunk ResponseChunk) ([]byte, error) {
	if chunk.Version != 1 {
		return nil, InvalidVersionError
	}

	return []byte{chunk.Version, chunk.Status}, nil
}

func (b Builder) BuildRequest(chunk RequestChunk) ([]byte, error) {
	if chunk.Version != 1 {
		return nil, InvalidVersionError
	}

	if len(chunk.Name) > 256 {
		return nil, NameTooLongError
	}

	if len(chunk.Password) > 256 {
		return nil, PasswordTooLongError
	}

	buffer := bytes.Buffer{}

	buffer.WriteByte(chunk.Version)

	buffer.WriteByte(byte(len(chunk.Name)))
	buffer.Write([]byte(chunk.Name))

	buffer.WriteByte(byte(len(chunk.Password)))
	buffer.Write([]byte(chunk.Password))

	return buffer.Bytes(), nil
}
