package v4a

import (
	"github.com/stretchr/testify/require"
	"net"
	"socks/protocol/v4a"
	"socks/test/stand/config"
	"testing"
)

type Comparator struct {
	t       *testing.T
	config  config.Config
	builder v4a.Builder
}

func NewComparator(t *testing.T, config config.Config, builder v4a.Builder) (Comparator, error) {
	return Comparator{t: t, config: config, builder: builder}, nil
}

func (c Comparator) compare(port uint16, ip net.IP, conn net.Conn) {
	actual := make([]byte, 512)

	i, err := conn.Read(actual)

	require.NoError(c.t, err)

	expected, buildErr := c.builder.BuildResponse(v4a.ResponseChunk{
		SocksVersion:    0,
		CommandCode:     90,
		DestinationPort: port,
		DestinationIp:   ip,
	})

	require.NoError(c.t, buildErr)

	require.Equal(c.t, expected, actual[:i])
}

func (c Comparator) CompareConnectResponse(conn net.Conn) {
	c.compare(c.config.Socks.TcpPort, net.IP{0, 0, 0, 0}, conn)
}

func (c Comparator) CompareBindResponse(port uint16, conn net.Conn) {
	ip := net.ParseIP(c.config.Server.IPv4)

	require.NotNil(c.t, ip)

	ip = ip.To4()

	require.NotNil(c.t, ip)

	c.compare(port, ip, conn)
}
