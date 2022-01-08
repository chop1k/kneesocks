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

func TestV5NoAuthentication(t *testing.T) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", socksTcpHost, socksTcpPort))

	require.NoError(t, err)

	writer := *bufio.NewWriter(conn)

	testV5ConnectByDomainAuthenticate(1, writer, conn, t)
}

func TestV5PasswordAuthentication(t *testing.T) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", socksTcpHost, socksTcpPort))

	require.NoError(t, err)

	writer := *bufio.NewWriter(conn)

	testV5PasswordAuthenticationAuthenticate(1, writer, conn, t)
}

func testV5PasswordAuthenticationAuthenticate(picture byte, writer bufio.Writer, reader io.Reader, t *testing.T) {
	require.NoError(t, writer.WriteByte(5))
	require.NoError(t, writer.WriteByte(1))
	require.NoError(t, writer.WriteByte(2))

	require.NoError(t, writer.Flush())

	selection := make([]byte, 10)

	_, err := reader.Read(selection)

	require.NoError(t, err)

	require.Equal(t, selection[1], byte(2))

	testV5PasswordAuthenticationSendAuthenticationRequest(picture, writer, reader, t)
}

func testV5PasswordAuthenticationSendAuthenticationRequest(picture byte, writer bufio.Writer, reader io.Reader, t *testing.T) {
	require.NoError(t, writer.WriteByte(1))
	require.NoError(t, writer.WriteByte(4))

	_, err := writer.Write([]byte("test"))

	require.NoError(t, err)

	require.NoError(t, writer.WriteByte(4))

	_, err = writer.Write([]byte("test"))

	require.NoError(t, err)

	require.NoError(t, writer.Flush())

	selection := make([]byte, 10)

	_, err = reader.Read(selection)

	require.NoError(t, err)

	require.Equal(t, selection[0], byte(1))
	require.Equal(t, selection[1], byte(0))

	testV5PasswordAuthenticationSendRequest(picture, writer, reader, t)
}

func testV5PasswordAuthenticationSendRequest(picture byte, writer bufio.Writer, reader io.Reader, t *testing.T) {
	require.NoError(t, writer.WriteByte(5))
	require.NoError(t, writer.WriteByte(1))
	require.NoError(t, writer.WriteByte(0))
	require.NoError(t, writer.WriteByte(3))

	host := []byte(tcpServerHost)

	require.NoError(t, writer.WriteByte(byte(len(host))))

	_, err := writer.Write(host)

	require.NoError(t, err)

	require.NoError(t, binary.Write(&writer, binary.BigEndian, tcpServerPort))

	require.NoError(t, writer.Flush())

	testV5PasswordAuthenticationReceiveReply(picture, writer, reader, t)
}

func testV5PasswordAuthenticationReceiveReply(picture byte, writer bufio.Writer, reader io.Reader, t *testing.T) {
	response := make([]byte, 22)

	_, err := reader.Read(response)

	require.NoError(t, err)

	require.Equal(t, response[0], byte(5))
	require.Equal(t, response[1], byte(0))
	require.Equal(t, response[2], byte(0))
	require.Equal(t, response[3], byte(4))

	ipv6 := net.IP{
		response[4], response[5], response[6], response[7],
		response[8], response[9], response[10], response[11],
		response[12], response[13], response[14], response[15],
		response[16], response[17], response[18], response[19],
	}.To16()

	ip := net.ParseIP(tcpServerIPv6)

	require.NotNil(t, ip)

	ip = ip.To16()

	require.NotNil(t, ip)

	require.Equal(t, ip, ipv6)

	port := make([]byte, 2)

	binary.LittleEndian.PutUint16(port, tcpServerPort)

	require.Equal(t, port, response[20:22])

	testV5PasswordAuthenticationSendPictureRequest(picture, writer, reader, t)
}

func testV5PasswordAuthenticationSendPictureRequest(picture byte, writer bufio.Writer, reader io.Reader, t *testing.T) {
	require.NoError(t, writer.WriteByte(picture))

	require.NoError(t, writer.Flush())

	testV5PasswordAuthenticationReceivePicture(picture, reader, t)
}

func testV5PasswordAuthenticationReceivePicture(picture byte, reader io.Reader, t *testing.T) {
	var file *os.File
	var err error

	if picture == 1 {
		file, err = ioutil.TempFile("", "v5-password-authentication-big-picture")
	} else if picture == 2 {
		file, err = ioutil.TempFile("", "v5-password-authentication-middle-picture")
	} else if picture == 3 {
		file, err = ioutil.TempFile("", "v5-password-authentication-small-picture")
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
