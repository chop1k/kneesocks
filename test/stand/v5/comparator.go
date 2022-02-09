package v5

import (
	"github.com/stretchr/testify/require"
	"net"
	"socks/protocol/auth/password"
	v5 "socks/protocol/v5"
	"socks/test/stand/config"
	"testing"
)

type Comparator struct {
	t               *testing.T
	config          config.Config
	builder         v5.Builder
	passwordBuilder password.Builder
}

func NewComparator(
	t *testing.T,
	config config.Config,
	builder v5.Builder,
	passwordBuilder password.Builder,
) (Comparator, error) {
	return Comparator{
		t:               t,
		config:          config,
		builder:         builder,
		passwordBuilder: passwordBuilder,
	}, nil
}

func (c Comparator) CompareSelection(method byte, conn net.Conn) {
	actual := make([]byte, 2)

	_, err := conn.Read(actual)

	require.NoError(c.t, err)

	expected, buildErr := c.builder.BuildMethodSelection(v5.MethodSelectionChunk{
		SocksVersion: 5,
		Method:       method,
	})

	require.NoError(c.t, buildErr)

	require.Equal(c.t, expected, actual)
}

func (c Comparator) ComparePassword(conn net.Conn) {
	actual := make([]byte, 600)

	i, err := conn.Read(actual)

	require.NoError(c.t, err)

	expected, buildErr := c.passwordBuilder.BuildResponse(password.ResponseChunk{
		Version: 1,
		Status:  0,
	})

	require.NoError(c.t, buildErr)

	require.Equal(c.t, expected, actual[:i])
}

func (c Comparator) CompareConnectResponse(addressType byte, conn net.Conn) {
	actual := make([]byte, 512)

	i, err := conn.Read(actual)

	require.NoError(c.t, err)

	var address string

	if addressType == 1 {
		address = c.config.Server.IPv4
	} else if addressType == 3 {
		address = c.config.Server.Domain
	} else if addressType == 4 {
		address = c.config.Server.IPv6
	} else {
		require.Fail(c.t, "Unsupported address type \"%d\". ", addressType)
	}

	expected, buildErr := c.builder.BuildResponse(v5.ResponseChunk{
		SocksVersion: 5,
		ReplyCode:    0,
		AddressType:  addressType,
		Address:      address,
		Port:         c.config.Server.ConnectPort,
	})

	require.NoError(c.t, buildErr)

	require.Equal(c.t, expected, actual[:i])
}

func (c Comparator) CompareFirstBindResponse(conn net.Conn) {
	actual := make([]byte, 512)

	i, err := conn.Read(actual)

	require.NoError(c.t, err)

	expected, buildErr := c.builder.BuildResponse(v5.ResponseChunk{
		SocksVersion: 5,
		ReplyCode:    0,
		AddressType:  1,
		Address:      "0.0.0.0",
		Port:         c.config.Socks.TcpPort,
	})

	require.NoError(c.t, buildErr)

	require.Equal(c.t, expected, actual[:i])
}

func (c Comparator) CompareSecondBindResponse(port uint16, addressType byte, conn net.Conn) {
	actual := make([]byte, 512)

	i, err := conn.Read(actual)

	require.NoError(c.t, err)

	var address string

	if addressType == 1 {
		address = c.config.Server.IPv4
	} else if addressType == 3 {
		address = c.config.Server.Domain
	} else if addressType == 4 {
		address = c.config.Server.IPv6
	} else {
		require.Fail(c.t, "Unsupported address type \"%d\". ", addressType)
	}

	expected, buildErr := c.builder.BuildResponse(v5.ResponseChunk{
		SocksVersion: 5,
		ReplyCode:    0,
		AddressType:  addressType,
		Address:      address,
		Port:         port,
	})

	require.NoError(c.t, buildErr)

	require.Equal(c.t, expected, actual[:i])
}

func (c Comparator) CompareAssociateResponse(conn net.Conn) {
	actual := make([]byte, 512)

	i, err := conn.Read(actual)

	require.NoError(c.t, err)

	expected, buildErr := c.builder.BuildResponse(v5.ResponseChunk{
		SocksVersion: 5,
		ReplyCode:    0,
		AddressType:  1,
		Address:      "0.0.0.0",
		Port:         c.config.Socks.UdpPort,
	})

	require.NoError(c.t, buildErr)

	require.Equal(c.t, expected, actual[:i])
}
