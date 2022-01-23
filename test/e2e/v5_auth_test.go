package e2e

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestV5NoAuthentication(t *testing.T) {
	conn := connectToServer(t)

	sendV5Request(conn, 1, 3, tcpServerHost, tcpServerPort, t)
	compareV5Reply(conn, 4, tcpServerIPv6, tcpServerPort, t)
	sendPictureRequest(conn, 1, t)
	comparePictures(conn, "v5", "auth", 1, t)
}

func TestV5PasswordAuthentication(t *testing.T) {
	conn := connectToServer(t)

	sendV5Password(conn, "test", "test", t)
	compareV5Password(conn, t)

	chunk := constructV5Request(1, 3, tcpServerHost, tcpServerPort, t)

	_, err := conn.Write(chunk)

	require.NoError(t, err)

	compareV5Reply(conn, 4, tcpServerIPv6, tcpServerPort, t)
	sendPictureRequest(conn, 1, t)
	comparePictures(conn, "v5", "auth", 1, t)
}
