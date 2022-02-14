package v4a

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
			[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0},
			RequestChunk{},
			InvalidSocksVersionError,
		},
		{
			[]byte{4, 2, 1, 187, 0, 0, 0, 0, 0},
			RequestChunk{
				SocksVersion:    4,
				CommandCode:     2,
				DestinationPort: 443,
				DestinationIp:   net.IP{0, 0, 0, 0},
				Domain:          "",
			},
			nil,
		},
		{
			[]byte{4, 2, 0, 80, 0, 0, 0, 0, 101, 110, 46, 119, 105, 107, 105, 112, 101, 100, 105, 97, 46, 111, 114, 103, 0},
			RequestChunk{
				SocksVersion:    4,
				CommandCode:     2,
				DestinationPort: 80,
				DestinationIp:   net.IP{0, 0, 0, 0},
				Domain:          "en.wikipedia.org",
			},
			nil,
		},
		{
			[]byte{4, 2, 1, 187, 0, 0, 0, 0, 101, 110, 46, 119, 105, 107, 105, 112, 101, 100, 105, 97, 46, 111, 114, 103, 0},
			RequestChunk{
				SocksVersion:    4,
				CommandCode:     2,
				DestinationPort: 443,
				DestinationIp:   net.IP{0, 0, 0, 0},
				Domain:          "en.wikipedia.org",
			},
			nil,
		},
		{
			[]byte{},
			RequestChunk{},
			InvalidChunkSizeError,
		},
	}

	parser := NewParser()

	for i, test := range tests {
		result, err := parser.ParseRequest(test.bytes)

		require.ErrorIsf(t, err, test.err, "Errors not equals (%d). ", i)

		if err == nil {
			require.Equalf(t, result, test.result, "Chunk not equals (%d). ", i)
		}
	}
}
