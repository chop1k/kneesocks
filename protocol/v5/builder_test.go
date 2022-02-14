package v5

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewBaseBuilder(t *testing.T) {
}

func TestBaseBuilder_BuildMethodSelection(t *testing.T) {
	tests := []struct {
		chunk  MethodSelectionChunk
		result []byte
		err    error
	}{
		{
			MethodSelectionChunk{
				SocksVersion: 0,
				Method:       0,
			},
			[]byte{0, 0},
			nil,
		},
		{
			MethodSelectionChunk{
				SocksVersion: 4,
				Method:       0,
			},
			[]byte{4, 0},
			nil,
		},
		{
			MethodSelectionChunk{
				SocksVersion: 4,
				Method:       1,
			},
			[]byte{4, 1},
			nil,
		},
	}

	builder, err := NewBuilder()

	require.NoError(t, err)

	for i, test := range tests {
		result, err := builder.BuildMethodSelection(test.chunk)

		require.ErrorIsf(t, err, test.err, "Error not equals (%d). ", i)

		require.Equalf(t, result, test.result, "Bytes not equals (%d). ", i)
	}
}

func TestBaseBuilder_BuildResponse(t *testing.T) {
	tests := []struct {
		chunk  ResponseChunk
		result []byte
		err    error
	}{
		{
			ResponseChunk{
				SocksVersion: 5,
				ReplyCode:    3,
				Port:         443,
				AddressType:  1,
				Address:      "255.40.3.2",
			},
			[]byte{5, 3, 0, 1, 255, 40, 3, 2, 187, 1},
			nil,
		},
		{
			ResponseChunk{
				SocksVersion: 0,
				ReplyCode:    0,
				Port:         0,
				AddressType:  1,
				Address:      "0.0.0.0",
			},
			[]byte{0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
			nil,
		},
		{
			ResponseChunk{
				SocksVersion: 5,
				ReplyCode:    2,
				Port:         443,
				AddressType:  3,
				Address:      "en.wikipedia.org",
			},
			[]byte{5, 2, 0, 3, 101, 110, 46, 119, 105, 107, 105, 112, 101, 100, 105, 97, 46, 111, 114, 103, 187, 1},
			nil,
		},
		{
			ResponseChunk{
				SocksVersion: 5,
				ReplyCode:    6,
				Port:         80,
				AddressType:  4,
				Address:      "ff02::2:114",
			},
			[]byte{5, 6, 0, 4, 255, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 1, 20, 80, 0},
			nil,
		},
		{
			ResponseChunk{
				SocksVersion: 4,
				ReplyCode:    91,
				Port:         0,
				AddressType:  255,
				Address:      "a",
			},
			nil,
			UnknownAddressTypeError,
		},
	}

	builder, err := NewBuilder()

	require.NoError(t, err)

	for i, test := range tests {
		result, err := builder.BuildResponse(test.chunk)

		require.ErrorIsf(t, err, test.err, "Error not equals (%d). ", i)

		require.Equalf(t, result, test.result, "Bytes not equals (%d). ", i)
	}
}

func TestBaseBuilder_BuildMethods(t *testing.T) {
	t.Skip("Not implemented.")
}

func TestBaseBuilder_BuildRequest(t *testing.T) {
	t.Skip("Not implemented.")
}

func TestBaseBuilder_BuildUdpRequest(t *testing.T) {
	t.Skip("Not implemented.")
}
