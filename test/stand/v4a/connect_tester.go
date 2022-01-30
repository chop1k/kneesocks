package v4a

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"net"
	"socks/test/stand/config"
	"socks/test/stand/server"
	"testing"
)

type ConnectTester struct {
	config     config.Config
	t          *testing.T
	server     server.Server
	sender     Sender
	comparator Comparator
	scope      config.Scope
}

func NewConnectTester(
	config config.Config,
	t *testing.T,
	server server.Server,
	sender Sender,
	comparator Comparator,
	scope config.Scope,
) (ConnectTester, error) {
	return ConnectTester{
		config:     config,
		t:          t,
		server:     server,
		sender:     sender,
		comparator: comparator,
		scope:      scope,
	}, nil
}

func (t ConnectTester) Test(number int) {
	scope := t.scope.GetV4aConnect(number)

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", t.config.Socks.IPv4, t.config.Socks.TcpPort))

	require.NoError(t.t, err)

	t.sender.SendConnectRequest(conn)
	t.comparator.CompareConnectResponse(conn)

	t.server.SendPictureRequest(scope.Picture, conn)
}
