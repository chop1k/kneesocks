package password

import (
	"github.com/stretchr/testify/require"
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
				Version: 0,
				Status:  0,
			},
			nil,
			InvalidVersionError,
		},
		{
			ResponseChunk{
				Version: 1,
				Status:  0,
			},
			[]byte{1, 0},
			nil,
		},
		{
			ResponseChunk{
				Version: 1,
				Status:  255,
			},
			[]byte{1, 255},
			nil,
		},
	}

	builder, err := NewBuilder()

	require.NoError(t, err)

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
