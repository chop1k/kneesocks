package v4

import (
	"github.com/stretchr/testify/require"
	"net"
	"testing"
)

func TestNewBaseBuilder(t *testing.T) {
}

func TestBaseBuilder_BuildResponse(t *testing.T) {
	tests := []struct {
		chunk  ResponseChunk
		result []byte
		err    error
	}{
		{
			ResponseChunk{
				SocksVersion:    4,
				CommandCode:     90,
				DestinationPort: 443,
				DestinationIp:   net.IP{255, 40, 3, 2},
			},
			[]byte{4, 90, 187, 1, 255, 40, 3, 2},
			nil,
		},
		{
			ResponseChunk{},
			[]byte{},
			DestinationIpIsNullError,
		},
		{
			ResponseChunk{
				SocksVersion:    255,
				CommandCode:     255,
				DestinationPort: 255,
				DestinationIp:   net.IP{0, 0, 0, 0},
			},
			[]byte{255, 255, 255, 0, 0, 0, 0, 0},
			nil,
		},
		{
			ResponseChunk{
				SocksVersion:    4,
				CommandCode:     91,
				DestinationPort: 60000,
				DestinationIp:   net.IP{127, 0, 0, 1},
			},
			[]byte{4, 91, 96, 234, 127, 0, 0, 1},
			nil,
		},
		{
			ResponseChunk{
				SocksVersion:    4,
				CommandCode:     91,
				DestinationPort: 0,
				DestinationIp:   net.IP{127, 0, 0, 1},
			},
			[]byte{4, 91, 0, 0, 127, 0, 0, 1},
			nil,
		},

		{
			ResponseChunk{
				SocksVersion:    4,
				CommandCode:     91,
				DestinationPort: 00000,
				DestinationIp:   net.IP{127, 0, 0, 1},
			},
			[]byte{4, 91, 0, 0, 127, 0, 0, 1},
			nil,
		},
	}

	builder := NewBuilder()

	for i, test := range tests {
		result, err := builder.BuildResponse(test.chunk)

		require.ErrorIsf(t, err, test.err, "Errors not equal (%d), expected `%s` to equal `%s`. ", i, err, test.err)

		if err == nil {
			require.Equalf(t, result, test.result, "Bytes not equal (%d), expected `%+v` to equal `%+v`. ", i, result, test.result)
		}
	}
}

func TestBaseBuilder_BuildRequest(t *testing.T) {
	t.Skip("Not implemented.")
}
