package e2e

import (
	"bufio"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"github.com/stretchr/testify/require"
	"io"
	"io/ioutil"
	"net"
	"os"
	"testing"
)

func TestV4aConnectWithBigPicture(t *testing.T) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", socksTcpHost, socksTcpPort))

	require.NoError(t, err)

	writer := *bufio.NewWriter(conn)

	testV4aConnectSendRequest(1, writer, conn, t)
}

func TestV4aConnectWithMiddlePicture(t *testing.T) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", socksTcpHost, socksTcpPort))

	require.NoError(t, err)

	writer := *bufio.NewWriter(conn)

	testV4aConnectSendRequest(2, writer, conn, t)
}

func TestV4aConnectWithSmallPicture(t *testing.T) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", socksTcpHost, socksTcpPort))

	require.NoError(t, err)

	writer := *bufio.NewWriter(conn)

	testV4aConnectSendRequest(3, writer, conn, t)
}

func testV4aConnectSendRequest(picture byte, writer bufio.Writer, reader io.Reader, t *testing.T) {
	require.NoError(t, writer.WriteByte(4))
	require.NoError(t, writer.WriteByte(1))

	require.NoError(t, binary.Write(&writer, binary.BigEndian, tcpServerPort))

	require.NoError(t, writer.WriteByte(0))
	require.NoError(t, writer.WriteByte(0))
	require.NoError(t, writer.WriteByte(0))
	require.NoError(t, writer.WriteByte(255))

	host := []byte(tcpServerHost)

	_, err := writer.Write(host)

	require.NoError(t, err)

	require.NoError(t, writer.WriteByte(0))

	require.NoError(t, writer.Flush())

	testV4aConnectReceiveReply(picture, writer, reader, t)
}

func testV4aConnectReceiveReply(picture byte, writer bufio.Writer, reader io.Reader, t *testing.T) {
	response := make([]byte, 8)

	i, err := reader.Read(response)

	require.NoError(t, err)

	require.Equal(t, i, 8)

	require.Equal(t, response[0], byte(0))
	require.Equal(t, response[1], byte(90))

	port := make([]byte, 2)

	binary.LittleEndian.PutUint16(port, socksTcpPort)

	require.Equal(t, port, response[2:4])

	require.Equal(t, response[4], byte(0))
	require.Equal(t, response[5], byte(0))
	require.Equal(t, response[6], byte(0))
	require.Equal(t, response[7], byte(0))

	testV4aConnectSendPictureRequest(picture, writer, reader, t)
}

func testV4aConnectSendPictureRequest(picture byte, writer bufio.Writer, reader io.Reader, t *testing.T) {
	require.NoError(t, writer.WriteByte(picture))

	require.NoError(t, writer.Flush())

	testV4aConnectReceivePicture(picture, reader, t)
}

func testV4aConnectReceivePicture(picture byte, reader io.Reader, t *testing.T) {
	var file *os.File
	var err error

	if picture == 1 {
		file, err = ioutil.TempFile("", "v4a-connect-big-picture")
	} else if picture == 2 {
		file, err = ioutil.TempFile("", "v4a-connect-middle-picture")
	} else if picture == 3 {
		file, err = ioutil.TempFile("", "v4a-connect-small-picture")
	} else {
		t.Fatalf("Unknown picture %d. ", picture)
	}

	require.NoError(t, err)

	h := sha256.New()

	writers := io.MultiWriter(file, h)

	for {
		buffer := make([]byte, 512)

		i, err := reader.Read(buffer)

		if err != nil {
			break
		}

		_, err = writers.Write(buffer[:i])

		require.NoError(t, err)
	}

	if picture == 1 {
		require.Equal(t, bigPictureHash, fmt.Sprintf("%x", h.Sum(nil)))
	} else if picture == 2 {
		require.Equal(t, middlePictureHash, fmt.Sprintf("%x", h.Sum(nil)))
	} else if picture == 3 {
		require.Equal(t, smallPictureHash, fmt.Sprintf("%x", h.Sum(nil)))
	}

	h.Reset()

	require.NoError(t, file.Close())
}
