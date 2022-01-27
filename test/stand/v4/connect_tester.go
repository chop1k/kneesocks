package v4

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"net"
	v4 "socks/protocol/v4"
	"socks/test/stand/config"
	"testing"
)

type ConnectTest struct {
	config  config.Config
	t       *testing.T
	builder v4.Builder
}

func (t ConnectTest) Test() {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", t.config.Socks.IPv4, t.config.Socks.TcpPort))

	require.NoError(t.t, err)

	t.connectSendRequest(conn)
}

func (t ConnectTest) connectSendRequest(conn net.Conn) {
	ip := net.ParseIP(t.config.Server.IPv4)

	require.NotNil(t.t, ip)

	ip = ip.To4()

	require.NotNil(t.t, ip)

	request, buildErr := t.builder.BuildRequest(v4.RequestChunk{
		SocksVersion:    4,
		CommandCode:     1,
		DestinationPort: t.config.Server.TcpPort,
		DestinationIp:   ip,
		UserId:          "",
	})

	require.NoError(t.t, buildErr)

	_, err := conn.Write(request)

	require.NoError(t.t, err)
}
