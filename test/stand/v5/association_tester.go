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

type AssociationTester struct {
	t          *testing.T
	config     config.Config
	sender     Sender
	comparator Comparator
	server     server.Server
	scope      config.Scope
	picture    picture.Picture
}

func NewAssociationTester(
	t *testing.T,
	config config.Config,
	sender Sender,
	comparator Comparator,
	server server.Server,
	scope config.Scope,
	picture picture.Picture,
) (AssociationTester, error) {
	return AssociationTester{
		t:          t,
		config:     config,
		sender:     sender,
		comparator: comparator,
		server:     server,
		scope:      scope,
		picture:    picture,
	}, nil
}

func (t AssociationTester) Test(number int) {
	scope := t.scope.GetV5Associate(number)

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", t.config.Socks.IPv4, t.config.Socks.TcpPort))

	require.NoError(t.t, err)

	t.sender.SendMethods([]byte{0}, conn)
	t.comparator.CompareSelection(0, conn)
	t.sender.SendAssociateRequest(scope.AddressType, conn)
	t.comparator.CompareAssociateResponse(conn)

	udp := t.sender.SendPictureRequest(scope.Picture)

	t.picture.CompareUsingUdp(scope.Picture, udp)
}
