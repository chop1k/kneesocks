package v4a

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"net"
	"socks/protocol/v4a"
	"socks/test/stand/config"
	"socks/test/stand/server"
	"testing"
)

type ConnectTester struct {
	config  config.Config
	t       *testing.T
	builder v4a.Builder
	server  server.Server
}

func NewConnectTester(
	config config.Config,
	t *testing.T,
	builder v4a.Builder,
	server server.Server,
) (ConnectTester, error) {
	return ConnectTester{
		config:  config,
		t:       t,
		builder: builder,
		server:  server,
	}, nil
}

func (t ConnectTester) Test(number int) {
	if number == 1 {
		t.connect(1)
	} else if number == 2 {
		t.connect(2)
	} else if number == 3 {
		t.connect(3)
	} else {
		require.Fail(t.t, "Unsupported test number \"%d\".", number)
	}
}

func (t ConnectTester) connect(picture byte) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", t.config.Socks.IPv4, t.config.Socks.TcpPort))

	require.NoError(t.t, err)

	t.sendRequest(picture, conn)
}

func (t ConnectTester) sendRequest(picture byte, conn net.Conn) {
	request, buildErr := t.builder.BuildRequest(v4a.RequestChunk{
		SocksVersion:    4,
		CommandCode:     1,
		DestinationPort: t.config.Server.TcpPort,
		DestinationIp:   net.IP{0, 0, 0, 255},
		Domain:          t.config.Server.Domain,
	})

	require.NoError(t.t, buildErr)

	_, err := conn.Write(request)

	require.NoError(t.t, err)

	t.compareResponse(picture, conn)
}

func (t ConnectTester) compareResponse(picture byte, conn net.Conn) {
	actual := make([]byte, 512)

	i, err := conn.Read(actual)

	require.NoError(t.t, err)

	expected, buildErr := t.builder.BuildResponse(v4a.ResponseChunk{
		SocksVersion:    0,
		CommandCode:     90,
		DestinationPort: t.config.Socks.TcpPort,
		DestinationIp:   net.IP{0, 0, 0, 0},
	})

	require.NoError(t.t, buildErr)

	require.Equal(t.t, expected, actual[:i])

	t.server.SendPictureRequest(picture, conn)
}
