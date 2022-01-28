package v5

import (
	"github.com/stretchr/testify/require"
	"net"
	"socks/protocol/auth/password"
	v5 "socks/protocol/v5"
	"socks/test/stand/config"
	"testing"
)

type Sender struct {
	t               *testing.T
	config          config.Config
	builder         v5.Builder
	passwordBuilder password.Builder
}

func NewSender(
	t *testing.T,
	config config.Config,
	builder v5.Builder,
	passwordBuilder password.Builder,
) (Sender, error) {
	return Sender{
		t:               t,
		config:          config,
		builder:         builder,
		passwordBuilder: passwordBuilder,
	}, nil
}

func (s Sender) SendMethods(methods []byte, conn net.Conn) {
	chunk, err := s.builder.BuildMethods(v5.MethodsChunk{
		SocksVersion: 5,
		Methods:      methods,
	})

	require.NoError(s.t, err)

	_, err = conn.Write(chunk)

	require.NoError(s.t, err)
}

func (s Sender) SendPassword(conn net.Conn) {
	request, err := s.passwordBuilder.BuildRequest(password.RequestChunk{
		Version:  1,
		Name:     s.config.User.Name,
		Password: s.config.User.Password,
	})

	require.NoError(s.t, err)

	_, err = conn.Write(request)

	require.NoError(s.t, err)
}

func (s Sender) SendRequest(command byte, addressType byte, conn net.Conn) {
	var address string

	if addressType == 1 {
		address = s.config.Server.IPv4
	} else if addressType == 3 {
		address = s.config.Server.Domain
	} else if addressType == 4 {
		address = s.config.Server.IPv6
	} else {
		require.Fail(s.t, "Unsupported address type \"%d\". ", addressType)
	}

	request, err := s.builder.BuildRequest(v5.RequestChunk{
		SocksVersion: 5,
		CommandCode:  command,
		AddressType:  addressType,
		Address:      address,
		Port:         s.config.Server.TcpPort,
	})

	require.NoError(s.t, err)

	_, err = conn.Write(request)

	require.NoError(s.t, err)
}
