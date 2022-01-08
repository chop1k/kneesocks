package password

import "errors"

var (
	InvalidVersionError = errors.New("Invalid version, must be 1. ")
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
	if chunk.Version != 1 {
		return nil, InvalidVersionError
	}

	return []byte{chunk.Version, chunk.Status}, nil
}
