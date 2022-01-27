package v4

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"net"
	"socks/cmd/e2e_test_server/protocol"
	v4 "socks/protocol/v4"
	"socks/test/stand/config"
	"socks/test/stand/picture"
	"testing"
)

type ConnectTester struct {
	config        config.Config
	t             *testing.T
	builder       v4.Builder
	serverBuilder protocol.Builder
	picture       picture.Picture
}

func NewConnectTester(
	config config.Config,
	t *testing.T,
	builder v4.Builder,
	serverBuilder protocol.Builder,
	picture picture.Picture,
) (ConnectTester, error) {
	return ConnectTester{
		config:        config,
		t:             t,
		builder:       builder,
		serverBuilder: serverBuilder,
		picture:       picture,
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

	t.compareResponse(picture, conn)
}

func (t ConnectTester) compareResponse(picture byte, conn net.Conn) {
	actual := make([]byte, 512)

	i, err := conn.Read(actual)

	require.NoError(t.t, err)

	expected, buildErr := t.builder.BuildResponse(v4.ResponseChunk{
		SocksVersion:    0,
		CommandCode:     90,
		DestinationPort: t.config.Socks.TcpPort,
		DestinationIp:   net.IP{0, 0, 0, 0},
	})

	require.NoError(t.t, buildErr)

	require.Equal(t.t, expected, actual[:i])

	t.sendPictureRequest(picture, conn)
}

func (t ConnectTester) sendPictureRequest(picture byte, conn net.Conn) {
	request, err := t.serverBuilder.BuildRequest(protocol.RequestChunk{
		Command:     1,
		Picture:     picture,
		AddressType: 1,
		Address:     net.IP{127, 0, 0, 1},
		Port:        0,
	})

	require.NoError(t.t, err)

	_, err = conn.Write(request)

	require.NoError(t.t, err)

	t.picture.Compare(picture, conn)
}
