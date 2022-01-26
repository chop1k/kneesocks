package e2e

import (
	"testing"
)

func TestV5BindByDomainWithBigPicture(t *testing.T) {
	t.Skip("Random behavior.¯\\_( ͡° ͜ʖ ͡°)_/¯")

	conn := connectToServer(t)

	sendV5Request(conn, 2, 3, tcpServerBindHost, 10035, t)
	compareV5Reply(conn, 1, "0.0.0.0", socksTcpPort, t)

	host := connectToHost(tcpServerIPv4, tcpServerPort, t)

	sendPictureRequest(host, 2, 1, tcpServerBindIPv4, 10035, 1, t)
	compareV5Reply(conn, 1, tcpServerBindIPv4, 10035, t)
	comparePictures(conn, "v5", "bind-with-by-domain", 1, t)
}

func TestV5BindByDomainWithMiddlePicture(t *testing.T) {
	t.Skip("Random behavior.¯\\_( ͡° ͜ʖ ͡°)_/¯")

	conn := connectToServer(t)

	sendV5Request(conn, 2, 3, tcpServerBindHost, 10036, t)
	compareV5Reply(conn, 1, "0.0.0.0", socksTcpPort, t)

	host := connectToHost(tcpServerIPv4, tcpServerPort, t)

	sendPictureRequest(host, 2, 1, tcpServerBindIPv4, 10036, 2, t)
	compareV5Reply(conn, 1, tcpServerBindIPv4, 10036, t)
	comparePictures(conn, "v5", "bind-with-by-domain", 2, t)
}

func TestV5BindByDomainWithSmallPicture(t *testing.T) {
	t.Skip("Random behavior.¯\\_( ͡° ͜ʖ ͡°)_/¯")

	conn := connectToServer(t)

	sendV5Request(conn, 2, 3, tcpServerBindHost, 10037, t)
	compareV5Reply(conn, 1, "0.0.0.0", socksTcpPort, t)

	host := connectToHost(tcpServerIPv4, tcpServerPort, t)

	sendPictureRequest(host, 2, 1, tcpServerBindIPv4, 10037, 3, t)
	compareV5Reply(conn, 1, tcpServerBindIPv4, 10037, t)
	comparePictures(conn, "v5", "bind-with-by-domain", 3, t)
}

func TestV5BindByIPv4WithBigPicture(t *testing.T) {
	t.Skip("Random behavior.¯\\_( ͡° ͜ʖ ͡°)_/¯")

	conn := connectToServer(t)

	sendV5Request(conn, 2, 1, tcpServerBindIPv4, 10040, t)
	compareV5Reply(conn, 1, "0.0.0.0", socksTcpPort, t)

	host := connectToHost(tcpServerIPv4, tcpServerPort, t)

	sendPictureRequest(host, 2, 1, tcpServerBindIPv4, 10040, 1, t)
	compareV5Reply(conn, 1, tcpServerBindIPv4, 10040, t)
	comparePictures(conn, "v5", "bind-with-by-ipv4", 1, t)
}

func TestV5BindByIPv4WithMiddlePicture(t *testing.T) {
	t.Skip("Random behavior.¯\\_( ͡° ͜ʖ ͡°)_/¯")

	conn := connectToServer(t)

	sendV5Request(conn, 2, 1, tcpServerBindIPv4, 10041, t)
	compareV5Reply(conn, 1, "0.0.0.0", socksTcpPort, t)

	host := connectToHost(tcpServerIPv4, tcpServerPort, t)

	sendPictureRequest(host, 2, 1, tcpServerBindIPv4, 10041, 2, t)
	compareV5Reply(conn, 1, tcpServerBindIPv4, 10041, t)
	comparePictures(conn, "v5", "bind-with-by-ipv4", 2, t)
}

func TestV5BindByIPv4WithSmallPicture(t *testing.T) {
	t.Skip("Random behavior.¯\\_( ͡° ͜ʖ ͡°)_/¯")

	conn := connectToServer(t)

	sendV5Request(conn, 2, 1, tcpServerBindIPv4, 10042, t)
	compareV5Reply(conn, 1, "0.0.0.0", socksTcpPort, t)

	host := connectToHost(tcpServerIPv4, tcpServerPort, t)

	sendPictureRequest(host, 2, 1, tcpServerBindIPv4, 10042, 3, t)
	compareV5Reply(conn, 1, tcpServerBindIPv4, 10042, t)
	comparePictures(conn, "v5", "bind-with-by-ipv4", 3, t)
}

func TestV5BindByIPv6WithBigPicture(t *testing.T) {
	t.Skip("Random behavior.¯\\_( ͡° ͜ʖ ͡°)_/¯")

	conn := connectToServer(t)

	sendV5Request(conn, 2, 4, tcpServerBindIPv6, 10050, t)
	compareV5Reply(conn, 1, "0.0.0.0", socksTcpPort, t)

	host := connectToHost(tcpServerIPv4, tcpServerPort, t)

	sendPictureRequest(host, 2, 1, tcpServerBindIPv4, 10050, 1, t)
	compareV5Reply(conn, 1, tcpServerBindIPv4, 10050, t)
	comparePictures(conn, "v5", "bind-with-by-ipv6", 1, t)
}

func TestV5BindByIPv6WithMiddlePicture(t *testing.T) {
	t.Skip("Random behavior.¯\\_( ͡° ͜ʖ ͡°)_/¯")

	conn := connectToServer(t)

	sendV5Request(conn, 2, 4, tcpServerBindIPv6, 10051, t)
	compareV5Reply(conn, 1, "0.0.0.0", socksTcpPort, t)

	host := connectToHost(tcpServerIPv4, tcpServerPort, t)

	sendPictureRequest(host, 2, 1, tcpServerBindIPv4, 10051, 2, t)
	compareV5Reply(conn, 1, tcpServerBindIPv4, 10051, t)
	comparePictures(conn, "v5", "bind-with-by-ipv6", 2, t)
}

func TestV5BindByIPv6WithSmallPicture(t *testing.T) {
	t.Skip("Random behavior.¯\\_( ͡° ͜ʖ ͡°)_/¯")

	conn := connectToServer(t)

	sendV5Request(conn, 2, 1, tcpServerBindIPv4, 10052, t)
	compareV5Reply(conn, 1, "0.0.0.0", socksTcpPort, t)

	host := connectToHost(tcpServerIPv4, tcpServerPort, t)

	sendPictureRequest(host, 2, 1, tcpServerBindIPv4, 10052, 3, t)
	compareV5Reply(conn, 1, tcpServerBindIPv4, 10052, t)
	comparePictures(conn, "v5", "bind-with-by-ipv6", 3, t)
}
