package v4a

import (
	"github.com/stretchr/testify/require"
	"net"
	"socks/protocol/v4a"
	"socks/test/stand/config"
	"testing"
)

type Sender struct {
	t       *testing.T
	config  config.Config
	builder v4a.Builder
}

func NewSender(t *testing.T, config config.Config, builder v4a.Builder) (Sender, error) {
	return Sender{t: t, config: config, builder: builder}, nil
}

func (s Sender) send(command byte, port uint16, conn net.Conn) {
	request, buildErr := s.builder.BuildRequest(v4a.RequestChunk{
		SocksVersion:    4,
		CommandCode:     command,
		DestinationPort: port,
		DestinationIp:   net.IP{0, 0, 0, 255},
		Domain:          s.config.Server.Domain,
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
