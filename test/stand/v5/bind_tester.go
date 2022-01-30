package v5

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
	t          *testing.T
	config     config.Config
	sender     Sender
	comparator Comparator
	server     server.Server
	scope      config.Scope
	picture    picture.Picture
}

func NewBindTester(
	t *testing.T,
	config config.Config,
	sender Sender,
	comparator Comparator,
	server server.Server,
	scope config.Scope,
	picture picture.Picture,
) (BindTester, error) {
	return BindTester{
		t:          t,
		config:     config,
		sender:     sender,
		comparator: comparator,
		server:     server,
		scope:      scope,
		picture:    picture,
	}, nil
}

func (t BindTester) Test(number int) {
	scope := t.scope.GetV5Bind(number)

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", t.config.Socks.IPv4, t.config.Socks.TcpPort))

	require.NoError(t.t, err)

	t.sender.SendMethods([]byte{0}, conn)
	t.comparator.CompareSelection(0, conn)
	t.sender.SendBindRequest(scope.Port, scope.AddressType, conn)
	t.comparator.CompareFirstBindResponse(conn)
	t.server.SendBindRequest(scope.Picture, scope.AddressType, scope.Port)
	t.comparator.CompareSecondBindResponse(scope.Port, scope.AddressType, conn)
	t.picture.Compare(scope.Picture, conn)
}
