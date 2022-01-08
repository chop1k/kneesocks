package v4

import (
	"github.com/stretchr/testify/require"
	"net"
	"testing"
)

func TestNewBaseParser(t *testing.T) {
}

func TestBaseParser_ParseRequest(t *testing.T) {
	tests := []struct {
		bytes  []byte
		result RequestChunk
		err    error
	}{
		{
			[]byte{4, 2, 0, 80, 255, 40, 3, 2, 0},
			RequestChunk{
				SocksVersion:    4,
				CommandCode:     2,
				DestinationPort: 80,
				DestinationIp:   net.IP{255, 40, 3, 2},
			},
			nil,
		},
		{
			[]byte{4, 2, 1, 187, 255, 40, 3, 2, 0},
			RequestChunk{
				SocksVersion:    4,
				CommandCode:     2,
				DestinationPort: 443,
				DestinationIp:   net.IP{255, 40, 3, 2},
			},
			nil,
		},
		{
			[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0},
			RequestChunk{},
			InvalidSocksVersionError,
		},
		{
			[]byte{4, 255, 0, 255, 0, 0, 0, 0, 0},
			RequestChunk{
				SocksVersion:    4,
				CommandCode:     255,
				DestinationPort: 255,
				DestinationIp:   net.IP{0, 0, 0, 0},
			},
			nil,
		},
		{
			[]byte{4, 91, 234, 96, 127, 0, 0, 1, 0},
			RequestChunk{
				SocksVersion:    4,
				CommandCode:     91,
				DestinationPort: 60000,
				DestinationIp:   net.IP{127, 0, 0, 1},
			},
			nil,
		},
		{
			[]byte{4, 91, 0, 0, 127, 0, 0, 1, 0},
			RequestChunk{
				SocksVersion:    4,
				CommandCode:     91,
				DestinationPort: 0,
				DestinationIp:   net.IP{127, 0, 0, 1},
			},
			nil,
		},
		{
			[]byte{4, 91, 0, 0, 127, 0, 0, 1, 0},
			RequestChunk{
				SocksVersion:    4,
				CommandCode:     91,
				DestinationPort: 00000,
				DestinationIp:   net.IP{127, 0, 0, 1},
			},
			nil,
		},
		{
			[]byte{},
			RequestChunk{},
			InvalidChunkSizeError,
		},
	}

	parser := NewBaseParser()

	for i, test := range tests {
		result, err := parser.ParseRequest(test.bytes)

		require.ErrorIsf(t, err, test.err, "Errors not equal (%d), expected `%s` to equal `%s`. ", i, err, test.err)

		if err == nil {
			require.EqualValuesf(t, result, test.result, "Chunks not equals (%d). ", i)
		}
	}
}
