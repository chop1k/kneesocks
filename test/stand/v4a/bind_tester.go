package v4a

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
	server     server.Server
	sender     Sender
	comparator Comparator
	scope      config.Scope
}

func NewBindTester(
	config config.Config,
	t *testing.T,
	picture picture.Picture,
	server server.Server,
	sender Sender,
	comparator Comparator,
	scope config.Scope,
) (BindTester, error) {
	return BindTester{
		config:     config,
		t:          t,
		picture:    picture,
		server:     server,
		sender:     sender,
		comparator: comparator,
		scope:      scope,
	}, nil
}

func (t BindTester) Test(number int) {
	scope := t.scope.GetV4aBind(number)

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", t.config.Socks.IPv4, t.config.Socks.TcpPort))

	require.NoError(t.t, err)

	t.sender.SendBindRequest(scope.Port, conn)
	t.comparator.CompareConnectResponse(conn)
	t.server.SendBindRequest(scope.Picture, 1, scope.Port)
	t.comparator.CompareBindResponse(scope.Port, conn)
	t.picture.Compare(scope.Picture, conn)
}
