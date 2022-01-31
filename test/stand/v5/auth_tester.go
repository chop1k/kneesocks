package v5

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"net"
	"socks/test/stand/config"
	"socks/test/stand/server"
	"testing"
)

type AuthTester struct {
	t          *testing.T
	config     config.Config
	server     server.Server
	sender     Sender
	comparator Comparator
	scope      config.Scope
}

func NewAuthTester(
	t *testing.T,
	config config.Config,
	server server.Server,
	sender Sender,
	comparator Comparator,
	scope config.Scope,
) (AuthTester, error) {
	return AuthTester{
		t:          t,
		config:     config,
		server:     server,
		sender:     sender,
		comparator: comparator,
		scope:      scope,
	}, nil
}

func (t AuthTester) Test(number int) {
	scope := t.scope.GetV5Auth(number)

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", t.config.Socks.IPv4, t.config.Socks.TcpPort))

	require.NoError(t.t, err)

	if scope.Method == 0 {
		t.handleNoAuth(scope.Picture, scope.AddressType, conn)
	} else if scope.Method == 2 {
		t.handlePasswordAuth(scope.Picture, scope.AddressType, conn)
	} else {
		require.Fail(t.t, "Unsupported method. ")
	}
}

func (t AuthTester) handleNoAuth(picture byte, addressType byte, conn net.Conn) {
	t.sender.SendMethods([]byte{0}, conn)
	t.comparator.CompareSelection(0, conn)
	t.sender.SendConnectRequest(addressType, conn)

	if addressType == 3 {
		t.comparator.CompareConnectResponse(4, conn)
	} else {
		t.comparator.CompareConnectResponse(addressType, conn)
	}

	t.server.SendPictureRequest(picture, conn)
}

func (t AuthTester) handlePasswordAuth(picture byte, addressType byte, conn net.Conn) {
	t.sender.SendMethods([]byte{2}, conn)
	t.comparator.CompareSelection(2, conn)
	t.sender.SendPassword(conn)
	t.comparator.ComparePassword(conn)
	t.sender.SendConnectRequest(addressType, conn)

	if addressType == 3 {
		t.comparator.CompareConnectResponse(4, conn)
	} else {
		t.comparator.CompareConnectResponse(addressType, conn)
	}

	t.server.SendPictureRequest(picture, conn)
}
