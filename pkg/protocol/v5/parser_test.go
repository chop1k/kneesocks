package v5

import (
	"github.com/stretchr/testify/require"
	"socks/pkg/utils"
	"testing"
)

func TestNewBaseParser(t *testing.T) {
}

func TestBaseParser_ParseMethods(t *testing.T) {
	tests := []struct {
		bytes  []byte
		result MethodsChunk
		err    error
	}{
		{
			[]byte{5, 0},
			MethodsChunk{},
			InvalidChunkSizeError,
		},
		{
			[]byte{5, 2, 0},
			MethodsChunk{},
			InvalidChunkSizeError,
		},
		{
			[]byte{},
			MethodsChunk{},
			InvalidChunkSizeError,
		},
		{
			[]byte{228},
			MethodsChunk{},
			InvalidChunkSizeError,
		},
		{
			[]byte{0, 0, 0},
			MethodsChunk{},
			InvalidSocksVersionError,
		},
		{
			[]byte{5, 1, 0, 1, 2},
			MethodsChunk{
				SocksVersion: 5,
				Methods:      []byte{0},
			},
			nil,
		},
		{
			[]byte{5, 2, 0, 1},
			MethodsChunk{
				SocksVersion: 5,
				Methods:      []byte{0, 1},
			},
			nil,
		},
	}

	parser := NewParser(utils.AddressUtils{})

	for i, test := range tests {
		result, err := parser.ParseMethods(test.bytes)

		require.ErrorIsf(t, err, test.err, "Error not equals (%d). ", i)

		require.Equalf(t, result, test.result, "Bytes not equals (%d). ", i)
	}
}

func TestBaseParser_ParseRequest(t *testing.T) {
	tests := []struct {
		bytes  []byte
		result RequestChunk
		err    error
	}{
		{
			[]byte{5, 1, 0, 1, 1, 2, 3, 4, 1, 187},
			RequestChunk{
				SocksVersion: 5,
				CommandCode:  1,
				AddressType:  1,
				Address:      "1.2.3.4",
				Port:         443,
			},
			nil,
		},
		{
			[]byte{5, 1, 0, 3, 16, 114, 117, 46, 119, 105, 107, 105, 112, 101, 100, 105, 97, 46, 111, 114, 103, 1, 187},
			RequestChunk{
				SocksVersion: 5,
				CommandCode:  1,
				AddressType:  3,
				Address:      "ru.wikipedia.org",
				Port:         443,
			},
			nil,
		},
		{
			[]byte{5, 2, 0, 4, 255, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 1, 20, 0, 80},
			RequestChunk{
				SocksVersion: 5,
				CommandCode:  2,
				AddressType:  4,
				Address:      "ff02::2:114",
				Port:         80,
			},
			nil,
		},
		{
			[]byte{},
			RequestChunk{},
			InvalidChunkSizeError,
		},
		{
			[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			RequestChunk{},
			InvalidSocksVersionError,
		},
		{
			[]byte{5, 9, 0, 1, 1, 2, 3, 4, 1, 187},
			RequestChunk{},
			InvalidCommandCodeError,
		},
		{
			[]byte{5, 2, 0, 2, 1, 2, 3, 4, 1, 187},
			RequestChunk{},
			UnknownAddressTypeError,
		},
		{
			[]byte{5, 2, 255, 1, 1, 2, 3, 4, 1, 187},
			RequestChunk{},
			InvalidReservedByteError,
		},
		{
			[]byte{5, 1, 0, 3, 90, 0, 0, 0, 0, 0, 1, 187},
			RequestChunk{},
			InvalidChunkSizeError,
		},
		{
			[]byte{5, 1, 0, 3, 2, 114, 117, 0, 0, 0, 0, 0, 0, 1, 187},
			RequestChunk{
				SocksVersion: 5,
				CommandCode:  1,
				AddressType:  3,
				Address:      "ru",
				Port:         0,
			},
			nil,
		},
		{
			[]byte{5, 1, 0, 3, 16, 114, 117, 46, 119, 105, 107, 105, 112, 101, 100, 105, 97, 46, 111, 114, 103},
			RequestChunk{},
			InvalidChunkSizeError,
		},
		{
			[]byte{5, 1, 0, 3, 16, 114, 117, 46, 119, 105, 107, 105, 112, 101, 100, 105, 97, 46, 111, 114, 103, 0},
			RequestChunk{},
			InvalidChunkSizeError,
		},
		{
			[]byte{5, 2, 0, 4, 255, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			RequestChunk{},
			InvalidChunkSizeError,
		},
		{
			[]byte{5, 2, 0, 4, 255, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 1, 20},
			RequestChunk{},
			InvalidChunkSizeError,
		},
		{
			[]byte{5, 2, 0, 4, 255, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 1, 20, 0},
			RequestChunk{},
			InvalidChunkSizeError,
		},
	}

	parser := NewParser(utils.AddressUtils{})

	for i, test := range tests {
		result, err := parser.ParseRequest(test.bytes)

		require.ErrorIsf(t, err, test.err, "Error not equals (%d). ", i)

		require.Equalf(t, result, test.result, "Bytes not equals (%d). ", i)
	}
}

func TestBaseParser_ParseUdpRequest(t *testing.T) {
	tests := []struct {
		bytes  []byte
		result UdpRequest
		err    error
	}{
		{
			[]byte{0, 0, 0, 1, 1, 2, 3, 4, 1, 187, 1, 2, 3, 4},
			UdpRequest{
				Fragment:    0,
				AddressType: 1,
				Address:     "1.2.3.4",
				Port:        443,
				Data:        []byte{1, 2, 3, 4},
			},
			nil,
		},
		{
			[]byte{0, 0, 0, 3, 16, 114, 117, 46, 119, 105, 107, 105, 112, 101, 100, 105, 97, 46, 111, 114, 103, 1, 187, 1, 2, 3, 4},
			UdpRequest{
				Fragment:    0,
				AddressType: 3,
				Address:     "ru.wikipedia.org",
				Port:        443,
				Data:        []byte{1, 2, 3, 4},
			},
			nil,
		},
		{
			[]byte{0, 0, 0, 4, 255, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 1, 20, 0, 80, 1, 2, 3, 4},
			UdpRequest{
				Fragment:    0,
				AddressType: 4,
				Address:     "ff02::2:114",
				Port:        80,
				Data:        []byte{1, 2, 3, 4},
			},
			nil,
		},
		{
			[]byte{},
			UdpRequest{},
			InvalidChunkSizeError,
		},
		{
			[]byte{0, 0, 0, 2, 1, 2, 3, 4, 1, 187},
			UdpRequest{},
			UnknownAddressTypeError,
		},
		{
			[]byte{1, 2, 0, 1, 1, 2, 3, 4, 1, 187},
			UdpRequest{},
			InvalidReservedByteError,
		},
		{
			[]byte{0, 2, 0, 1, 1, 2, 3, 4, 1, 187},
			UdpRequest{},
			InvalidReservedByteError,
		},
		{
			[]byte{1, 0, 0, 1, 1, 2, 3, 4, 1, 187},
			UdpRequest{},
			InvalidReservedByteError,
		},
		{
			[]byte{0, 0, 0, 3, 90, 0, 0, 0, 0, 0, 1, 187},
			UdpRequest{},
			InvalidChunkSizeError,
		},
		{
			[]byte{0, 0, 0, 3, 2, 114, 117, 1, 187, 1, 2, 3, 4},
			UdpRequest{
				Fragment:    0,
				AddressType: 3,
				Address:     "ru",
				Port:        443,
				Data:        []byte{1, 2, 3, 4},
			},
			nil,
		},
		{
			[]byte{0, 0, 0, 3, 16, 114, 117, 46, 119, 105, 107, 105, 112, 101, 100, 105, 97, 46, 111, 114, 103},
			UdpRequest{},
			InvalidChunkSizeError,
		},
		{
			[]byte{0, 0, 0, 3, 16, 114, 117, 46, 119, 105, 107, 105, 112, 101, 100, 105, 97, 46, 111, 114, 103, 0},
			UdpRequest{},
			InvalidChunkSizeError,
		},
		{
			[]byte{0, 0, 0, 4, 255, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			UdpRequest{},
			InvalidChunkSizeError,
		},
		{
			[]byte{0, 0, 0, 4, 255, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 1, 20},
			UdpRequest{},
			InvalidChunkSizeError,
		},
		{
			[]byte{0, 0, 0, 4, 255, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 1, 20, 0},
			UdpRequest{},
			InvalidChunkSizeError,
		},
		{
			[]byte{0, 0, 0, 1, 1, 2, 3, 4, 1, 187},
			UdpRequest{
				Fragment:    0,
				AddressType: 1,
				Address:     "1.2.3.4",
				Port:        443,
				Data:        []byte{},
			},
			nil,
		},
		{
			[]byte{0, 0, 0, 3, 16, 114, 117, 46, 119, 105, 107, 105, 112, 101, 100, 105, 97, 46, 111, 114, 103, 1, 187},
			UdpRequest{
				Fragment:    0,
				AddressType: 3,
				Address:     "ru.wikipedia.org",
				Port:        443,
				Data:        []byte{},
			},
			nil,
		},
		{
			[]byte{0, 0, 0, 4, 255, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 1, 20, 0, 80},
			UdpRequest{
				Fragment:    0,
				AddressType: 4,
				Address:     "ff02::2:114",
				Port:        80,
				Data:        []byte{},
			},
			nil,
		},
		{
			[]byte{0, 0, 228, 1, 1, 2, 3, 4, 1, 187, 1, 2, 3, 4},
			UdpRequest{},
			InvalidFragmentByteError,
		},

		{
			[]byte{0, 0, 0, 1, 1, 2, 3, 4, 1, 187, 1},
			UdpRequest{
				Fragment:    0,
				AddressType: 1,
				Address:     "1.2.3.4",
				Port:        443,
				Data:        []byte{1},
			},
			nil,
		},
	}

	parser := NewParser(utils.AddressUtils{})

	for i, test := range tests {
		result, err := parser.ParseUdpRequest(test.bytes)

		require.ErrorIsf(t, err, test.err, "Error not equals (%d). ", i)

		require.Equalf(t, result, test.result, "Bytes not equals (%d). ", i)
	}
}
