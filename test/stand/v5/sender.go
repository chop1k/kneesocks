package v5

import (
	"fmt"
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

func (s Sender) sendRequest(command byte, addressType byte, port uint16, conn net.Conn) {
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
		Port:         port,
	})

	require.NoError(s.t, err)

	_, err = conn.Write(request)

	require.NoError(s.t, err)
}

func (s Sender) SendConnectRequest(addressType byte, conn net.Conn) {
	s.sendRequest(1, addressType, s.config.Server.ConnectPort, conn)
}

func (s Sender) SendBindRequest(port uint16, addressType byte, conn net.Conn) {
	s.sendRequest(2, addressType, port, conn)
}

func (s Sender) SendAssociateRequest(addressType byte, conn net.Conn) {
	s.sendRequest(3, addressType, s.config.Server.ConnectPort, conn)
}

func (s Sender) SendPictureRequest(picture byte) net.PacketConn {
	addr, lookupErr := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", s.config.Socks.IPv4, s.config.Socks.UdpPort))

	require.NoError(s.t, lookupErr)

	conn, err := net.DialUDP("udp", nil, addr)

	require.NoError(s.t, err)

	request, buildErr := s.builder.BuildUdpRequest(v5.UdpRequest{
		Fragment:    0,
		AddressType: 1,
		Address:     s.config.Server.IPv4,
		Port:        s.config.Server.UdpPort,
		Data:        []byte{picture},
	})

	require.NoError(s.t, buildErr)

	_, err = conn.Write(request)

	require.NoError(s.t, err)

	return conn
}
