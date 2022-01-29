package server

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
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

func (s Server) SendBindRequest(picture byte, port uint16) {
	ip := net.ParseIP(s.config.Socks.IPv4)

	require.NotNil(s.t, ip)

	ip = ip.To4()

	require.NotNil(s.t, ip)

	chunk := struct {
		Picture     byte
		AddressType byte
		Address     net.IP
		Port        uint16
	}{
		picture, 1, ip, port,
	}

	data, err := json.Marshal(chunk)

	require.NoError(s.t, err)

	s.sendBindRequest(data)
}

func (s Server) sendBindRequest(json []byte) {
	address := fmt.Sprintf("http://%s:%d/%s", s.config.Server.IPv4, s.config.Server.HttpPort, s.config.Server.HttpUri)

	resp, err := resty.New().R().SetHeader("content-type", "application/json").SetBody(json).Post(address)

	require.NoError(s.t, err)

	require.Equal(s.t, "200", resp.Status())
}
