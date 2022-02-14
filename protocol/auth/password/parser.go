package password

import "errors"

var (
	InvalidChunkSizeError    = errors.New("Invalid chunk size. ")
	InvalidNameSizeError     = errors.New("Invalid name size. ")
	InvalidPasswordSizeError = errors.New("Invalid password size. ")
)

type Parser struct {
}

func NewParser() Parser {
	return Parser{}
}

func (b Parser) ParseRequest(bytes []byte) (RequestChunk, error) {
	length := len(bytes)

	if length < 5 {
		return RequestChunk{}, InvalidChunkSizeError
	}

	if bytes[0] != 1 {
		return RequestChunk{}, InvalidVersionError
	}

	if bytes[1] <= 0 {
		return RequestChunk{}, InvalidNameSizeError
	}

	nameStart := 2
	nameEnd := int(bytes[1]) + 2

	if nameEnd <= 0 {
		return RequestChunk{}, InvalidNameSizeError
	}

	if length < nameEnd+1 {
		return RequestChunk{}, InvalidChunkSizeError
	}

	if bytes[nameEnd] <= 0 {
		return RequestChunk{}, InvalidPasswordSizeError
	}

	passwordStart := nameEnd + 1
	passwordEnd := nameEnd + 1 + int(bytes[nameEnd])

	if length < passwordEnd || length < passwordStart {
		return RequestChunk{}, InvalidChunkSizeError
	}

	return RequestChunk{
		Version:  1,
		Name:     string(bytes[nameStart:nameEnd]),
		Password: string(bytes[passwordStart:passwordEnd]),
	}, nil
}
