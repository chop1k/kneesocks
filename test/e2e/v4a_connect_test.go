package e2e

import (
	"testing"
)

func TestV4aConnectWithBigPicture(t *testing.T) {
	conn := connectToServer(t)

	sendV4aRequest(conn, 1, tcpServerHost, tcpServerPort, t)
	compareV4Reply(conn, socksTcpPort, t)
	sendPictureRequest(conn, 1, t)
	comparePictures(conn, "v4a", "connect", 1, t)
}

func TestV4aConnectWithMiddlePicture(t *testing.T) {
	conn := connectToServer(t)

	sendV4aRequest(conn, 1, tcpServerHost, tcpServerPort, t)
	compareV4Reply(conn, socksTcpPort, t)
	sendPictureRequest(conn, 2, t)
	comparePictures(conn, "v4a", "connect", 2, t)
}

func TestV4aConnectWithSmallPicture(t *testing.T) {
	conn := connectToServer(t)

	sendV4aRequest(conn, 1, tcpServerHost, tcpServerPort, t)
	compareV4Reply(conn, socksTcpPort, t)
	sendPictureRequest(conn, 3, t)
	comparePictures(conn, "v4a", "connect", 3, t)
}
