package server

import (
	"github.com/stretchr/testify/require"
	"net"
	"socks/cmd/e2e_test_server/protocol"
	"socks/test/stand/config"
	"socks/test/stand/picture"
	"testing"
)

type Server struct {
	t       *testing.T
	config  config.Config
	builder protocol.Builder
	picture picture.Picture
}

func NewServer(
	t *testing.T,
	config config.Config,
	builder protocol.Builder,
	picture picture.Picture,
) (Server, error) {
	return Server{
		t:       t,
		config:  config,
		builder: builder,
		picture: picture,
	}, nil
}

func (s Server) SendPictureRequest(picture byte, conn net.Conn) {
	request, err := s.builder.BuildRequest(protocol.RequestChunk{
		Command:     1,
		Picture:     picture,
		AddressType: 1,
		Address:     net.IP{0, 0, 0, 0},
		Port:        0,
	})

	require.NoError(s.t, err)

	_, err = conn.Write(request)

	require.NoError(s.t, err)

	s.picture.Compare(picture, conn)
}
