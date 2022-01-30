package v4

import (
	"github.com/stretchr/testify/require"
	"net"
	v4 "socks/protocol/v4"
	"socks/test/stand/config"
	"testing"
)

type Sender struct {
	t       *testing.T
	config  config.Config
	builder v4.Builder
}

func NewSender(t *testing.T, config config.Config, builder v4.Builder) (Sender, error) {
	return Sender{t: t, config: config, builder: builder}, nil
}

func (s Sender) send(command byte, port uint16, conn net.Conn) {
	ip := net.ParseIP(s.config.Server.IPv4)

	require.NotNil(s.t, ip)

	ip = ip.To4()

	require.NotNil(s.t, ip)

	request, buildErr := s.builder.BuildRequest(v4.RequestChunk{
		SocksVersion:    4,
		CommandCode:     command,
		DestinationPort: port,
		DestinationIp:   ip,
		UserId:          "",
	})

	require.NoError(s.t, buildErr)

	_, err := conn.Write(request)

	require.NoError(s.t, err)
}

func (s Sender) SendConnectRequest(conn net.Conn) {
	s.send(1, s.config.Server.ConnectPort, conn)
}

func (s Sender) SendBindRequest(port uint16, conn net.Conn) {
	s.send(2, port, conn)
}
