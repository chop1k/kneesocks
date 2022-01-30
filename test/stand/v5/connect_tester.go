package v5

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"net"
	"socks/test/stand/config"
	"socks/test/stand/server"
	"testing"
)

type ConnectTester struct {
	t          *testing.T
	config     config.Config
	sender     Sender
	comparator Comparator
	server     server.Server
	scope      config.Scope
}

func NewConnectTester(
	t *testing.T,
	config config.Config,
	sender Sender,
	comparator Comparator,
	server server.Server,
	scope config.Scope,
) (ConnectTester, error) {
	return ConnectTester{
		t:          t,
		config:     config,
		sender:     sender,
		comparator: comparator,
		server:     server,
		scope:      scope,
	}, nil
}

func (t ConnectTester) Test(number int) {
	scope := t.scope.GetV5Connect(number)

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", t.config.Socks.IPv4, t.config.Socks.TcpPort))

	require.NoError(t.t, err)

	t.sender.SendMethods([]byte{0}, conn)
	t.comparator.CompareSelection(0, conn)
	t.sender.SendConnectRequest(scope.AddressType, conn)
	t.comparator.CompareConnectResponse(scope.AddressType, conn)

	t.server.SendPictureRequest(scope.Picture, conn)
}
