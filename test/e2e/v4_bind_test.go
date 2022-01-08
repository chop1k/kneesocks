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

func TestV4BindWithBigPicture(t *testing.T) {
	t.Skipf("It works too unstable... idk what to do with this random behavior.")

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", socksTcpHost, socksTcpPort))

	require.NoError(t, err)

	writer := *bufio.NewWriter(conn)

	testV4BindSendRequest(4, writer, conn, t)
}

func TestV4BindWithMiddlePicture(t *testing.T) {
	t.Skipf("It works too unstable... idk what to do with this random behavior.")

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", socksTcpHost, socksTcpPort))

	require.NoError(t, err)

	writer := *bufio.NewWriter(conn)

	testV4BindSendRequest(5, writer, conn, t)
}

func TestV4BindWithSmallPicture(t *testing.T) {
	t.Skipf("It works too unstable... idk what to do with this random behavior.")

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", socksTcpHost, socksTcpPort))

	require.NoError(t, err)

	writer := *bufio.NewWriter(conn)

	testV4BindSendRequest(6, writer, conn, t)
}

func testV4BindSendRequest(picture byte, writer bufio.Writer, reader io.Reader, t *testing.T) {
	require.NoError(t, writer.WriteByte(4))
	require.NoError(t, writer.WriteByte(2))

	require.NoError(t, binary.Write(&writer, binary.BigEndian, tcpServerBindPort))

	ip := net.ParseIP(tcpServerBindIPv4)

	require.NotNil(t, ip)

	ip = ip.To4()

	require.NotNil(t, ip)

	_, err := writer.Write(ip)

	require.NoError(t, err)

	require.NoError(t, writer.WriteByte(0))

	require.NoError(t, writer.Flush())

	testV4BindReceiveFirstReply(picture, reader, t)
}

func testV4BindReceiveFirstReply(picture byte, reader io.Reader, t *testing.T) {
	response := make([]byte, 8)

	print('a')

	i, err := reader.Read(response)

	print('b')

	require.NoError(t, err)

	require.Equal(t, i, 8)

	require.Equal(t, response[0], byte(0))
	require.Equal(t, response[1], byte(90))
	require.Equal(t, response[2], byte(0))
	require.Equal(t, response[3], byte(0))
	require.Equal(t, response[4], byte(0))
	require.Equal(t, response[5], byte(0))
	require.Equal(t, response[6], byte(0))
	require.Equal(t, response[7], byte(0))

	testV4BindConnectToServer(picture, reader, t)
}

func testV4BindConnectToServer(picture byte, reader io.Reader, t *testing.T) {
	print('h')
	host, err := net.Dial("tcp", fmt.Sprintf("%s:%d", tcpServerHost, tcpServerPort))

	require.NoError(t, err)

	print('t')

	testV4BindSendPictureRequest(picture, *bufio.NewWriter(host), reader, t)
}

func testV4BindSendPictureRequest(picture byte, writer bufio.Writer, reader io.Reader, t *testing.T) {
	require.NoError(t, writer.WriteByte(picture))
	require.NoError(t, writer.Flush())

	testV4BindReceiveSecondReply(picture, reader, t)
}

func testV4BindReceiveSecondReply(picture byte, reader io.Reader, t *testing.T) {
	response := make([]byte, 8)

	i, err := reader.Read(response)

	require.NoError(t, err)

	require.Equal(t, 8, i)

	require.Equal(t, response[0], byte(0))
	require.Equal(t, response[1], byte(90))

	port := make([]byte, 2)

	binary.LittleEndian.PutUint16(port, tcpServerBindPort)

	require.Equal(t, port, response[2:4])

	ip := net.ParseIP(tcpServerBindIPv4)

	require.NotNil(t, ip)

	ip = ip.To4()

	require.NotNil(t, ip)

	require.Equal(t, ip, net.IP{response[4], response[5], response[6], response[7]})

	testV4BindReceivePicture(picture, reader, t)
}

func testV4BindReceivePicture(picture byte, reader io.Reader, t *testing.T) {
	var file *os.File
	var err error

	if picture == 4 {
		file, err = ioutil.TempFile("", "v4-bind-big-picture")
	} else if picture == 5 {
		file, err = ioutil.TempFile("", "v4-bind-middle-picture")
	} else if picture == 6 {
		file, err = ioutil.TempFile("", "v4-bind-small-picture")
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

	if picture == 4 {
		require.Equal(t, bigPictureHash, fmt.Sprintf("%x", h.Sum(nil)))
	} else if picture == 5 {
		require.Equal(t, middlePictureHash, fmt.Sprintf("%x", h.Sum(nil)))
	} else if picture == 6 {
		require.Equal(t, smallPictureHash, fmt.Sprintf("%x", h.Sum(nil)))
	}

	h.Reset()

	require.NoError(t, file.Close())
}
