package e2e

import (
	"socks/test/stand"
	"testing"
)

func TestV4aConnectWithBigPicture(t *testing.T) {
	stand.New().Execute("v4a", "connect", 1, t)
	//conn := connectToServer(t)
	//
	//sendV4aRequest(conn, 1, tcpServerHost, tcpServerPort, t)
	//compareV4Reply(conn, "0.0.0.0", socksTcpPort, t)
	//sendPictureRequest(conn, 1, 1, "0.0.0.0", 0, 1, t)
	//comparePictures(conn, "v4a", "connect", 1, t)
}

func TestV4aConnectWithMiddlePicture(t *testing.T) {
	stand.New().Execute("v4a", "connect", 2, t)
	//conn := connectToServer(t)
	//
	//sendV4aRequest(conn, 1, tcpServerHost, tcpServerPort, t)
	//compareV4Reply(conn, "0.0.0.0", socksTcpPort, t)
	//sendPictureRequest(conn, 1, 1, "0.0.0.0", 0, 2, t)
	//comparePictures(conn, "v4a", "connect", 2, t)
}

func TestV4aConnectWithSmallPicture(t *testing.T) {
	stand.New().Execute("v4a", "connect", 3, t)
	//conn := connectToServer(t)
	//
	//sendV4aRequest(conn, 1, tcpServerHost, tcpServerPort, t)
	//compareV4Reply(conn, "0.0.0.0", socksTcpPort, t)
	//sendPictureRequest(conn, 1, 1, "0.0.0.0", 0, 3, t)
	//comparePictures(conn, "v4a", "connect", 3, t)
}
