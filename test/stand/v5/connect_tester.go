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
}

func NewConnectTester(
	t *testing.T,
	config config.Config,
	sender Sender,
	comparator Comparator,
	server server.Server,
) (ConnectTester, error) {
	return ConnectTester{
		t:          t,
		config:     config,
		sender:     sender,
		comparator: comparator,
		server:     server,
	}, nil
}

func (t ConnectTester) Test(number int) {
	if number == 1 {
		t.connect(1, 1)
	} else if number == 2 {
		t.connect(1, 3)
	} else if number == 3 {
		t.connect(1, 4)
	} else if number == 4 {
		t.connect(2, 1)
	} else if number == 5 {
		t.connect(2, 3)
	} else if number == 6 {
		t.connect(2, 4)
	} else if number == 7 {
		t.connect(3, 1)
	} else if number == 8 {
		t.connect(3, 3)
	} else if number == 9 {
		t.connect(3, 4)
	} else {
		require.Fail(t.t, "Unsupported test number \"%d\".", number)
	}
}

func (t ConnectTester) connect(picture byte, addressType byte) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", t.config.Socks.IPv4, t.config.Socks.TcpPort))

	require.NoError(t.t, err)

	t.sender.SendMethods([]byte{0}, conn)
	t.comparator.CompareSelection(0, conn)
	t.sender.SendRequest(1, addressType, conn)
	t.comparator.CompareResponse(0, addressType, conn)

	t.server.SendPictureRequest(picture, conn)
}
