package e2e

import (
	"testing"
)

func TestV4aBindWithBigPicture(t *testing.T) {
	t.Skip("Random behavior.¯\\_( ͡° ͜ʖ ͡°)_/¯")

	conn := connectToServer(t)

	sendV4aRequest(conn, 2, tcpServerBindHost, 10020, t)
	compareV4Reply(conn, "0.0.0.0", socksTcpPort, t)

	host := connectToHost(tcpServerIPv4, tcpServerPort, t)

	sendPictureRequest(host, 2, 1, tcpServerBindIPv4, 10020, 1, t)
	compareV4Reply(conn, tcpServerBindIPv4, 10020, t)
	comparePictures(conn, "v4a", "bind", 1, t)
}

func TestV4aBindWithMiddlePicture(t *testing.T) {
	t.Skip("Random behavior.¯\\_( ͡° ͜ʖ ͡°)_/¯")

	conn := connectToServer(t)

	sendV4aRequest(conn, 2, tcpServerBindHost, 10021, t)
	compareV4Reply(conn, "0.0.0.0", socksTcpPort, t)

	host := connectToHost(tcpServerIPv4, tcpServerPort, t)

	sendPictureRequest(host, 2, 1, tcpServerBindIPv4, 10021, 2, t)
	compareV4Reply(conn, tcpServerBindIPv4, 10021, t)
	comparePictures(conn, "v4a", "bind", 2, t)
}

func TestV4aBindWithSmallPicture(t *testing.T) {
	t.Skip("Random behavior.¯\\_( ͡° ͜ʖ ͡°)_/¯")

	conn := connectToServer(t)

	sendV4aRequest(conn, 2, tcpServerBindHost, 10022, t)
	compareV4Reply(conn, "0.0.0.0", socksTcpPort, t)

	host := connectToHost(tcpServerIPv4, tcpServerPort, t)

	sendPictureRequest(host, 2, 1, tcpServerBindIPv4, 10022, 3, t)
	compareV4Reply(conn, tcpServerBindIPv4, 10022, t)
	comparePictures(conn, "v4a", "bind", 3, t)
}
