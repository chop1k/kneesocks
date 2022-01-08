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

func TestV5ConnectByDomainWithBigPicture(t *testing.T) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", socksTcpHost, socksTcpPort))

	require.NoError(t, err)

	writer := *bufio.NewWriter(conn)

	testV5ConnectByDomainAuthenticate(1, writer, conn, t)
}

func TestV5ConnectByDomainWithMiddlePicture(t *testing.T) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", socksTcpHost, socksTcpPort))

	require.NoError(t, err)

	writer := *bufio.NewWriter(conn)

	testV5ConnectByDomainAuthenticate(2, writer, conn, t)
}

func TestV5ConnectByDomainWithSmallPicture(t *testing.T) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", socksTcpHost, socksTcpPort))

	require.NoError(t, err)

	writer := *bufio.NewWriter(conn)

	testV5ConnectByDomainAuthenticate(3, writer, conn, t)
}

func testV5ConnectByDomainAuthenticate(picture byte, writer bufio.Writer, reader io.Reader, t *testing.T) {
	require.NoError(t, writer.WriteByte(5))
	require.NoError(t, writer.WriteByte(1))
	require.NoError(t, writer.WriteByte(0))

	require.NoError(t, writer.Flush())

	selection := make([]byte, 10)

	_, err := reader.Read(selection)

	require.NoError(t, err)

	require.Equal(t, selection[1], byte(0))

	testV5ConnectByDomainSendRequest(picture, writer, reader, t)
}

func testV5ConnectByDomainSendRequest(picture byte, writer bufio.Writer, reader io.Reader, t *testing.T) {
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

	testV5ConnectByDomainReceiveReply(picture, writer, reader, t)
}

func testV5ConnectByDomainReceiveReply(picture byte, writer bufio.Writer, reader io.Reader, t *testing.T) {
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

	testV5ConnectByDomainSendPictureRequest(picture, writer, reader, t)
}

func testV5ConnectByDomainSendPictureRequest(picture byte, writer bufio.Writer, reader io.Reader, t *testing.T) {
	require.NoError(t, writer.WriteByte(picture))

	require.NoError(t, writer.Flush())

	testV5ConnectByDomainReceivePicture(picture, reader, t)
}

func testV5ConnectByDomainReceivePicture(picture byte, reader io.Reader, t *testing.T) {
	var file *os.File
	var err error

	if picture == 1 {
		file, err = ioutil.TempFile("", "v5-connect-by-domain-big-picture")
	} else if picture == 2 {
		file, err = ioutil.TempFile("", "v5-connect-by-domain-middle-picture")
	} else if picture == 3 {
		file, err = ioutil.TempFile("", "v5-connect-by-domain-small-picture")
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

func TestV5ConnectByIPv4WithBigPicture(t *testing.T) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", socksTcpHost, socksTcpPort))

	require.NoError(t, err)

	writer := *bufio.NewWriter(conn)

	testV5ConnectByIPv4Authenticate(1, writer, conn, t)
}

func TestV5ConnectByIPv4WithMiddlePicture(t *testing.T) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", socksTcpHost, socksTcpPort))

	require.NoError(t, err)

	writer := *bufio.NewWriter(conn)

	testV5ConnectByIPv4Authenticate(2, writer, conn, t)
}

func TestV5ConnectByIPv4WithSmallPicture(t *testing.T) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", socksTcpHost, socksTcpPort))

	require.NoError(t, err)

	writer := *bufio.NewWriter(conn)

	testV5ConnectByIPv4Authenticate(3, writer, conn, t)
}

func testV5ConnectByIPv4Authenticate(picture byte, writer bufio.Writer, reader io.Reader, t *testing.T) {
	require.NoError(t, writer.WriteByte(5))
	require.NoError(t, writer.WriteByte(1))
	require.NoError(t, writer.WriteByte(0))

	require.NoError(t, writer.Flush())

	selection := make([]byte, 10)

	_, err := reader.Read(selection)

	require.NoError(t, err)

	require.Equal(t, selection[1], byte(0))

	testV5ConnectByIPv4SendRequest(picture, writer, reader, t)
}

func testV5ConnectByIPv4SendRequest(picture byte, writer bufio.Writer, reader io.Reader, t *testing.T) {
	require.NoError(t, writer.WriteByte(5))
	require.NoError(t, writer.WriteByte(1))
	require.NoError(t, writer.WriteByte(0))
	require.NoError(t, writer.WriteByte(1))

	ip := net.ParseIP(tcpServerIPv4)

	require.NotNil(t, ip)

	ip = ip.To4()

	require.NotNil(t, ip)

	require.NoError(t, writer.WriteByte(ip[0]))
	require.NoError(t, writer.WriteByte(ip[1]))
	require.NoError(t, writer.WriteByte(ip[2]))
	require.NoError(t, writer.WriteByte(ip[3]))

	require.NoError(t, binary.Write(&writer, binary.BigEndian, tcpServerPort))

	require.NoError(t, writer.Flush())

	testV5ConnectByIPv4ReceiveReply(ip, picture, writer, reader, t)
}

func testV5ConnectByIPv4ReceiveReply(ip net.IP, picture byte, writer bufio.Writer, reader io.Reader, t *testing.T) {
	response := make([]byte, 10)

	_, err := reader.Read(response)

	require.NoError(t, err)

	require.Equal(t, response[0], byte(5))
	require.Equal(t, response[1], byte(0))
	require.Equal(t, response[2], byte(0))
	require.Equal(t, response[3], byte(1))

	require.Equal(t, ip, net.IP{response[4], response[5], response[6], response[7]})

	port := make([]byte, 2)

	binary.LittleEndian.PutUint16(port, tcpServerPort)

	require.Equal(t, port, response[8:10])

	testV5ConnectByIPv4SendPictureRequest(picture, writer, reader, t)
}

func testV5ConnectByIPv4SendPictureRequest(picture byte, writer bufio.Writer, reader io.Reader, t *testing.T) {
	require.NoError(t, writer.WriteByte(picture))

	require.NoError(t, writer.Flush())

	testV5ConnectByIPv4ReceivePicture(picture, reader, t)
}

func testV5ConnectByIPv4ReceivePicture(picture byte, reader io.Reader, t *testing.T) {
	var file *os.File
	var err error

	if picture == 1 {
		file, err = ioutil.TempFile("", "v5-connect-by-ipv4-big-picture")
	} else if picture == 2 {
		file, err = ioutil.TempFile("", "v5-connect-by-ipv4-middle-picture")
	} else if picture == 3 {
		file, err = ioutil.TempFile("", "v5-connect-by-ipv4-small-picture")
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

func TestV5ConnectByIPv6WithBigPicture(t *testing.T) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", socksTcpHost, socksTcpPort))

	require.NoError(t, err)

	writer := *bufio.NewWriter(conn)

	testV5ConnectByIPv6Authenticate(1, writer, conn, t)
}

func TestV5ConnectByIPv6WithMiddlePicture(t *testing.T) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", socksTcpHost, socksTcpPort))

	require.NoError(t, err)

	writer := *bufio.NewWriter(conn)

	testV5ConnectByIPv6Authenticate(2, writer, conn, t)
}

func TestV5ConnectByIPv6WithSmallPicture(t *testing.T) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", socksTcpHost, socksTcpPort))

	require.NoError(t, err)

	writer := *bufio.NewWriter(conn)

	testV5ConnectByIPv6Authenticate(3, writer, conn, t)
}

func testV5ConnectByIPv6Authenticate(picture byte, writer bufio.Writer, reader io.Reader, t *testing.T) {
	require.NoError(t, writer.WriteByte(5))
	require.NoError(t, writer.WriteByte(1))
	require.NoError(t, writer.WriteByte(0))

	require.NoError(t, writer.Flush())

	selection := make([]byte, 10)

	_, err := reader.Read(selection)

	require.NoError(t, err)

	require.Equal(t, selection[1], byte(0))

	testV5ConnectByIPv6SendRequest(picture, writer, reader, t)
}

func testV5ConnectByIPv6SendRequest(picture byte, writer bufio.Writer, reader io.Reader, t *testing.T) {
	require.NoError(t, writer.WriteByte(5))
	require.NoError(t, writer.WriteByte(1))
	require.NoError(t, writer.WriteByte(0))
	require.NoError(t, writer.WriteByte(4))

	ip := net.ParseIP(tcpServerIPv6)

	require.NotNil(t, ip)

	ip = ip.To16()

	require.NotNil(t, ip)

	_, err := writer.Write(ip)

	require.NoError(t, err)

	require.NoError(t, binary.Write(&writer, binary.BigEndian, tcpServerPort))

	require.NoError(t, writer.Flush())

	testV5ConnectByIPv6ReceiveReply(ip, picture, writer, reader, t)
}

func testV5ConnectByIPv6ReceiveReply(ip net.IP, picture byte, writer bufio.Writer, reader io.Reader, t *testing.T) {
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

	require.Equal(t, ip, ipv6)

	port := make([]byte, 2)

	binary.LittleEndian.PutUint16(port, tcpServerPort)

	require.Equal(t, port, response[20:22])

	testV5ConnectByIPv6SendPictureRequest(picture, writer, reader, t)
}

func testV5ConnectByIPv6SendPictureRequest(picture byte, writer bufio.Writer, reader io.Reader, t *testing.T) {
	require.NoError(t, writer.WriteByte(picture))

	require.NoError(t, writer.Flush())

	testV5ConnectByIPv6ReceivePicture(picture, reader, t)
}

func testV5ConnectByIPv6ReceivePicture(picture byte, reader io.Reader, t *testing.T) {
	var file *os.File
	var err error

	if picture == 1 {
		file, err = ioutil.TempFile("", "v5-connect-by-ipv6-big-picture")
	} else if picture == 2 {
		file, err = ioutil.TempFile("", "v5-connect-by-ipv6-middle-picture")
	} else if picture == 3 {
		file, err = ioutil.TempFile("", "v5-connect-by-ipv6-small-picture")
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
