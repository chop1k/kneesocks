package utils

import "io"

type BufferReader struct {
}

func NewBufferReader() (BufferReader, error) {
	return BufferReader{}, nil
}

func (b BufferReader) Read(reader io.Reader, length int) ([]byte, error) {
	buffer := make([]byte, length)

	i, err := reader.Read(buffer)

	if err != nil {
		return nil, err
	}

	return buffer[:i], nil
}
