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
}

func NewAuthTester(
	t *testing.T,
	config config.Config,
	server server.Server,
	sender Sender,
	comparator Comparator,
) (AuthTester, error) {
	return AuthTester{
		t:          t,
		config:     config,
		server:     server,
		sender:     sender,
		comparator: comparator,
	}, nil
}

func (t AuthTester) Test(number int) {
	if number == 1 {
		t.handleNoAuth(1, 1)
	} else if number == 2 {
		t.handleNoAuth(1, 3)
	} else if number == 3 {
		t.handleNoAuth(1, 4)
	} else if number == 4 {
		t.handleNoAuth(2, 1)
	} else if number == 5 {
		t.handleNoAuth(2, 3)
	} else if number == 6 {
		t.handleNoAuth(2, 4)
	} else if number == 7 {
		t.handleNoAuth(3, 1)
	} else if number == 8 {
		t.handleNoAuth(3, 3)
	} else if number == 9 {
		t.handleNoAuth(3, 4)
	} else if number == 10 {
		t.handlePasswordAuth(1, 1)
	} else if number == 11 {
		t.handlePasswordAuth(1, 3)
	} else if number == 12 {
		t.handlePasswordAuth(1, 4)
	} else if number == 13 {
		t.handlePasswordAuth(2, 1)
	} else if number == 14 {
		t.handlePasswordAuth(2, 3)
	} else if number == 15 {
		t.handlePasswordAuth(2, 4)
	} else if number == 16 {
		t.handlePasswordAuth(3, 1)
	} else if number == 17 {
		t.handlePasswordAuth(3, 3)
	} else if number == 18 {
		t.handlePasswordAuth(3, 4)
	} else {
		require.Fail(t.t, "Unsupported test number \"%d\".", number)
	}
}

func (t AuthTester) handleNoAuth(picture byte, addressType byte) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", t.config.Socks.IPv4, t.config.Socks.TcpPort))

	require.NoError(t.t, err)

	t.sender.SendMethods([]byte{0}, conn)
	t.comparator.CompareSelection(0, conn)
	t.sender.SendRequest(1, addressType, conn)
	t.comparator.CompareResponse(0, addressType, conn)

	t.server.SendPictureRequest(picture, conn)
}

func (t AuthTester) handlePasswordAuth(picture byte, addressType byte) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", t.config.Socks.IPv4, t.config.Socks.TcpPort))

	require.NoError(t.t, err)

	t.sender.SendMethods([]byte{2}, conn)
	t.comparator.CompareSelection(2, conn)
	t.sender.SendPassword(conn)
	t.comparator.ComparePassword(conn)
	t.sender.SendRequest(1, addressType, conn)
	t.comparator.CompareResponse(0, addressType, conn)

	t.server.SendPictureRequest(picture, conn)
}
