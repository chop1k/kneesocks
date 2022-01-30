package server

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/stretchr/testify/require"
	"net"
	"socks/test/stand/config"
	"socks/test/stand/picture"
	"testing"
)

type Server struct {
	t       *testing.T
	config  config.Config
	picture picture.Picture
}

func NewServer(
	t *testing.T,
	config config.Config,
	picture picture.Picture,
) (Server, error) {
	return Server{
		t:       t,
		config:  config,
		picture: picture,
	}, nil
}

func (s Server) SendPictureRequest(picture byte, conn net.Conn) {
	_, err := conn.Write([]byte{picture})

	require.NoError(s.t, err)

	s.picture.Compare(picture, conn)
}

func (s Server) SendBindRequest(picture byte, addressType byte, port uint16) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", s.config.Server.IPv4, s.config.Server.BindPort))

	require.NoError(s.t, err)

	buffer := bytes.Buffer{}

	buffer.WriteByte(picture)
	buffer.WriteByte(addressType)

	err = binary.Write(&buffer, binary.BigEndian, port)

	require.NoError(s.t, err)

	_, err = conn.Write(buffer.Bytes())

	require.NoError(s.t, err)
}
