package e2e

import (
	"testing"
)

func TestV4ConnectWithBigPicture(t *testing.T) {
	conn := connectToServer(t)

	sendV4Request(conn, 1, tcpServerIPv4, tcpServerPort, t)
	compareV4Reply(conn, socksTcpPort, t)
	sendPictureRequest(conn, 1, t)
	comparePictures(conn, "v4", "connect", 1, t)
}

func TestV4ConnectWithMiddlePicture(t *testing.T) {
	conn := connectToServer(t)

	sendV4Request(conn, 1, tcpServerIPv4, tcpServerPort, t)
	compareV4Reply(conn, socksTcpPort, t)
	sendPictureRequest(conn, 2, t)
	comparePictures(conn, "v4", "connect", 2, t)
}

func TestV4ConnectWithSmallPicture(t *testing.T) {
	conn := connectToServer(t)

	sendV4Request(conn, 1, tcpServerIPv4, tcpServerPort, t)
	compareV4Reply(conn, socksTcpPort, t)
	sendPictureRequest(conn, 3, t)
	comparePictures(conn, "v4", "connect", 3, t)
}
