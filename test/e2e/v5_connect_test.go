package e2e

import (
	"testing"
)

func TestV5ConnectByDomainWithBigPicture(t *testing.T) {
	conn := connectToServer(t)

	sendV5Request(conn, 1, 3, tcpServerHost, tcpServerPort, t)
	compareV5Reply(conn, 4, tcpServerIPv6, tcpServerPort, t)
	sendPictureRequest(conn, 1, 1, "0.0.0.0", 0, 1, t)
	comparePictures(conn, "v5", "connect-by-domain", 1, t)
}

func TestV5ConnectByDomainWithMiddlePicture(t *testing.T) {
	conn := connectToServer(t)

	sendV5Request(conn, 1, 3, tcpServerHost, tcpServerPort, t)
	compareV5Reply(conn, 4, tcpServerIPv6, tcpServerPort, t)
	sendPictureRequest(conn, 1, 1, "0.0.0.0", 0, 2, t)
	comparePictures(conn, "v5", "connect-by-domain", 2, t)
}

func TestV5ConnectByDomainWithSmallPicture(t *testing.T) {
	//conn := connectToServer(t)
	//
	//sendV5Request(conn, 1, 3, tcpServerHost, tcpServerPort, t)
	//compareV5Reply(conn, 4, tcpServerIPv6, tcpServerPort, t)
	//sendPictureRequest(conn, 1, 1, "0.0.0.0", 0, 3, t)
	//comparePictures(conn, "v5", "connect-by-domain", 3, t)
}

func TestV5ConnectByIPv4WithBigPicture(t *testing.T) {
	conn := connectToServer(t)

	sendV5Request(conn, 1, 1, tcpServerIPv4, tcpServerPort, t)
	compareV5Reply(conn, 1, tcpServerIPv4, tcpServerPort, t)
	sendPictureRequest(conn, 1, 1, "0.0.0.0", 0, 1, t)
	comparePictures(conn, "v5", "connect-by-ipv4", 1, t)
}

func TestV5ConnectByIPv4WithMiddlePicture(t *testing.T) {
	conn := connectToServer(t)

	sendV5Request(conn, 1, 1, tcpServerIPv4, tcpServerPort, t)
	compareV5Reply(conn, 1, tcpServerIPv4, tcpServerPort, t)
	sendPictureRequest(conn, 1, 1, "0.0.0.0", 0, 2, t)
	comparePictures(conn, "v5", "connect-by-ipv4", 2, t)
}

func TestV5ConnectByIPv4WithSmallPicture(t *testing.T) {
	conn := connectToServer(t)

	sendV5Request(conn, 1, 1, tcpServerIPv4, tcpServerPort, t)
	compareV5Reply(conn, 1, tcpServerIPv4, tcpServerPort, t)
	sendPictureRequest(conn, 1, 1, "0.0.0.0", 0, 3, t)
	comparePictures(conn, "v5", "connect-by-ipv4", 3, t)
}

func TestV5ConnectByIPv6WithBigPicture(t *testing.T) {
	conn := connectToServer(t)

	sendV5Request(conn, 1, 4, tcpServerIPv6, tcpServerPort, t)
	compareV5Reply(conn, 4, tcpServerIPv6, tcpServerPort, t)
	sendPictureRequest(conn, 1, 1, "0.0.0.0", 0, 1, t)
	comparePictures(conn, "v5", "connect-by-ipv6", 1, t)
}

func TestV5ConnectByIPv6WithMiddlePicture(t *testing.T) {
	conn := connectToServer(t)

	sendV5Request(conn, 1, 4, tcpServerIPv6, tcpServerPort, t)
	compareV5Reply(conn, 4, tcpServerIPv6, tcpServerPort, t)
	sendPictureRequest(conn, 1, 1, "0.0.0.0", 0, 1, t)
	comparePictures(conn, "v5", "connect-by-ipv6", 1, t)
}

func TestV5ConnectByIPv6WithSmallPicture(t *testing.T) {
	conn := connectToServer(t)

	sendV5Request(conn, 1, 4, tcpServerIPv6, tcpServerPort, t)
	compareV5Reply(conn, 4, tcpServerIPv6, tcpServerPort, t)
	sendPictureRequest(conn, 1, 1, "0.0.0.0", 0, 1, t)
	comparePictures(conn, "v5", "connect-by-ipv6", 1, t)
}
