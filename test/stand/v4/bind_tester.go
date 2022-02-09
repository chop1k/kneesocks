package v4

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"net"
	"socks/test/stand/config"
	"socks/test/stand/picture"
	"socks/test/stand/server"
	"testing"
)

type BindTester struct {
	config     config.Config
	t          *testing.T
	picture    picture.Picture
	sender     Sender
	comparator Comparator
	scope      config.Scope
	server     server.Server
}

func NewBindTester(
	config config.Config,
	t *testing.T,
	picture picture.Picture,
	sender Sender,
	comparator Comparator,
	scope config.Scope,
	server server.Server,
) (BindTester, error) {
	return BindTester{
		config:     config,
		t:          t,
		picture:    picture,
		sender:     sender,
		comparator: comparator,
		scope:      scope,
		server:     server,
	}, nil
}

func (t BindTester) Test(number int) {
	scope := t.scope.GetV4Bind(number)

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", t.config.Socks.IPv4, t.config.Socks.TcpPort))

	require.NoError(t.t, err)

	t.sender.SendBindRequest(scope.Port, conn)
	t.comparator.CompareConnectResponse(conn)
	t.server.SendBindRequest(scope.Picture, 1, scope.Port)
	t.comparator.CompareBindResponse(scope.Port, conn)
	t.picture.CompareUsingTcp(scope.Picture, conn)
}
