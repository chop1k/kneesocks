package e2e

import (
	"testing"
)

func TestV4BindWithBigPicture(t *testing.T) {
	conn := connectToServer(t)

	sendV4Request(conn, 2, tcpServerBindIPv4, 10010, t)
	compareV4Reply(conn, "0.0.0.0", socksTcpPort, t)

	host := connectToHost(tcpServerIPv4, tcpServerPort, t)

	sendPictureRequest(host, 2, 1, tcpServerBindIPv4, 10010, 1, t)
	compareV4Reply(conn, tcpServerBindIPv4, 10010, t)
	comparePictures(conn, "v4", "bind", 1, t)
}

func TestV4BindWithMiddlePicture(t *testing.T) {
	conn := connectToServer(t)

	sendV4Request(conn, 2, tcpServerBindIPv4, 10011, t)
	compareV4Reply(conn, "0.0.0.0", socksTcpPort, t)

	host := connectToHost(tcpServerIPv4, tcpServerPort, t)

	sendPictureRequest(host, 2, 1, tcpServerBindIPv4, 10011, 2, t)
	compareV4Reply(conn, tcpServerBindIPv4, 10011, t)
	comparePictures(conn, "v4", "bind", 2, t)
}

func TestV4BindWithSmallPicture(t *testing.T) {
	conn := connectToServer(t)

	sendV4Request(conn, 2, tcpServerBindIPv4, 10012, t)
	compareV4Reply(conn, "0.0.0.0", socksTcpPort, t)

	host := connectToHost(tcpServerIPv4, tcpServerPort, t)

	sendPictureRequest(host, 2, 1, tcpServerBindIPv4, 10012, 3, t)
	compareV4Reply(conn, tcpServerBindIPv4, 10012, t)
	comparePictures(conn, "v4", "bind", 3, t)
}
