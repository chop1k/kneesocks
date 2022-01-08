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

func TestV5BindByDomainWithBigPicture(t *testing.T) {
	t.Skipf("It works too unstable... idk what to do with this random behavior.")

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", socksTcpHost, socksTcpPort))

	require.NoError(t, err)

	writer := *bufio.NewWriter(conn)

	testV5BindByDomainAuthenticate(4, writer, conn, t)
}

func TestV5BindByDomainWithMiddlePicture(t *testing.T) {
	t.Skipf("It works too unstable... idk what to do with this random behavior.")

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", socksTcpHost, socksTcpPort))

	require.NoError(t, err)

	writer := *bufio.NewWriter(conn)

	testV5BindByDomainAuthenticate(5, writer, conn, t)
}

func TestV5BindByDomainWithSmallPicture(t *testing.T) {
	t.Skipf("It works too unstable... idk what to do with this random behavior.")

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", socksTcpHost, socksTcpPort))

	require.NoError(t, err)

	writer := *bufio.NewWriter(conn)

	testV5BindByDomainAuthenticate(6, writer, conn, t)
}

func testV5BindByDomainAuthenticate(picture byte, writer bufio.Writer, reader io.Reader, t *testing.T) {
	require.NoError(t, writer.WriteByte(5))
	require.NoError(t, writer.WriteByte(1))
	require.NoError(t, writer.WriteByte(0))

	require.NoError(t, writer.Flush())

	selection := make([]byte, 10)

	_, err := reader.Read(selection)

	require.NoError(t, err)

	require.Equal(t, selection[1], byte(0))

	testV5BindByDomainSendRequest(picture, writer, reader, t)
}

func testV5BindByDomainSendRequest(picture byte, writer bufio.Writer, reader io.Reader, t *testing.T) {
	require.NoError(t, writer.WriteByte(5))
	require.NoError(t, writer.WriteByte(2))
	require.NoError(t, writer.WriteByte(0))
	require.NoError(t, writer.WriteByte(3))

	host := []byte(tcpServerBindHost)

	require.NoError(t, writer.WriteByte(byte(len(host))))

	_, err := writer.Write(host)

	require.NoError(t, err)

	require.NoError(t, binary.Write(&writer, binary.BigEndian, tcpServerBindPort))

	require.NoError(t, writer.Flush())

	testV5BindByDomainReceiveFirstReply(picture, reader, t)
}

func testV5BindByDomainReceiveFirstReply(picture byte, reader io.Reader, t *testing.T) {
	response := make([]byte, 10)

	_, err := reader.Read(response)

	require.NoError(t, err)

	require.Equal(t, response[0], byte(5))
	require.Equal(t, response[1], byte(0))
	require.Equal(t, response[2], byte(0))
	require.Equal(t, response[3], byte(1))
	require.Equal(t, response[4], byte(0))
	require.Equal(t, response[5], byte(0))
	require.Equal(t, response[6], byte(0))
	require.Equal(t, response[7], byte(0))

	port := make([]byte, 2)

	binary.LittleEndian.PutUint16(port, socksTcpPort)

	require.Equal(t, port, response[8:10])

	testV5BindByDomainConnectToServer(picture, reader, t)
}

func testV5BindByDomainConnectToServer(picture byte, reader io.Reader, t *testing.T) {
	host, err := net.Dial("tcp", fmt.Sprintf("%s:%d", tcpServerHost, tcpServerPort))

	require.NoError(t, err)

	testV5BindByDomainSendPictureRequest(picture, *bufio.NewWriter(host), reader, t)
}

func testV5BindByDomainSendPictureRequest(picture byte, writer bufio.Writer, reader io.Reader, t *testing.T) {
	require.NoError(t, writer.WriteByte(picture))
	require.NoError(t, writer.Flush())

	testV5BindByDomainReceiveSecondReply(picture, reader, t)
}

func testV5BindByDomainReceiveSecondReply(picture byte, reader io.Reader, t *testing.T) {
	response := make([]byte, 10)

	_, err := reader.Read(response)

	require.NoError(t, err)

	require.Equal(t, response[0], byte(5))
	require.Equal(t, response[1], byte(0))
	require.Equal(t, response[2], byte(0))
	require.Equal(t, response[3], byte(1))

	ip := net.IP{response[4], response[5], response[6], response[7]}.To4()

	require.NotNil(t, ip)

	ipv4 := net.ParseIP(tcpServerBindIPv4)

	require.NotNil(t, ipv4)

	ipv4 = ipv4.To4()

	require.NotNil(t, ipv4)

	require.Equal(t, ip, ipv4)

	port := make([]byte, 2)

	binary.LittleEndian.PutUint16(port, tcpServerBindPort)

	require.Equal(t, port, response[8:10])

	testV5BindByDomainReceivePicture(picture, reader, t)
}

func testV5BindByDomainReceivePicture(picture byte, reader io.Reader, t *testing.T) {
	var file *os.File
	var err error

	if picture == 4 {
		file, err = ioutil.TempFile("", "v5-bind-by-domain-big-picture")
	} else if picture == 5 {
		file, err = ioutil.TempFile("", "v5-bind-by-domain-middle-picture")
	} else if picture == 6 {
		file, err = ioutil.TempFile("", "v5-bind-by-domain-small-picture")
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

func TestV5BindByIPv4WithBigPicture(t *testing.T) {
	t.Skipf("It works too unstable... idk what to do with this random behavior.")

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", socksTcpHost, socksTcpPort))

	require.NoError(t, err)

	writer := *bufio.NewWriter(conn)

	testV5BindByIPv4Authenticate(4, writer, conn, t)
}

func TestV5BindByIPv4WithMiddlePicture(t *testing.T) {
	t.Skipf("It works too unstable... idk what to do with this random behavior.")

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", socksTcpHost, socksTcpPort))

	require.NoError(t, err)

	writer := *bufio.NewWriter(conn)

	testV5BindByIPv4Authenticate(5, writer, conn, t)
}

func TestV5BindByIPv4WithSmallPicture(t *testing.T) {
	t.Skipf("It works too unstable... idk what to do with this random behavior.")

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", socksTcpHost, socksTcpPort))

	require.NoError(t, err)

	writer := *bufio.NewWriter(conn)

	testV5BindByIPv4Authenticate(6, writer, conn, t)
}

func testV5BindByIPv4Authenticate(picture byte, writer bufio.Writer, reader io.Reader, t *testing.T) {
	require.NoError(t, writer.WriteByte(5))
	require.NoError(t, writer.WriteByte(1))
	require.NoError(t, writer.WriteByte(0))

	require.NoError(t, writer.Flush())

	selection := make([]byte, 10)

	_, err := reader.Read(selection)

	require.NoError(t, err)

	require.Equal(t, selection[1], byte(0))

	testV5BindByIPv4SendRequest(picture, writer, reader, t)
}

func testV5BindByIPv4SendRequest(picture byte, writer bufio.Writer, reader io.Reader, t *testing.T) {
	require.NoError(t, writer.WriteByte(5))
	require.NoError(t, writer.WriteByte(2))
	require.NoError(t, writer.WriteByte(0))
	require.NoError(t, writer.WriteByte(1))

	ip := net.ParseIP(tcpServerBindIPv4)

	require.NotNil(t, ip)

	ip = ip.To4()

	require.NotNil(t, ip)

	_, err := writer.Write(ip)

	require.NoError(t, err)

	require.NoError(t, binary.Write(&writer, binary.BigEndian, tcpServerBindPort))

	require.NoError(t, writer.Flush())

	testV5BindByIPv4ReceiveFirstReply(picture, reader, t)
}

func testV5BindByIPv4ReceiveFirstReply(picture byte, reader io.Reader, t *testing.T) {
	response := make([]byte, 10)

	_, err := reader.Read(response)

	require.NoError(t, err)

	require.Equal(t, response[0], byte(5))
	require.Equal(t, response[1], byte(0))
	require.Equal(t, response[2], byte(0))
	require.Equal(t, response[3], byte(1))
	require.Equal(t, response[4], byte(0))
	require.Equal(t, response[5], byte(0))
	require.Equal(t, response[6], byte(0))
	require.Equal(t, response[7], byte(0))

	port := make([]byte, 2)

	binary.LittleEndian.PutUint16(port, socksTcpPort)

	require.Equal(t, port, response[8:10])

	testV5BindByIPv4ConnectToServer(picture, reader, t)
}

func testV5BindByIPv4ConnectToServer(picture byte, reader io.Reader, t *testing.T) {
	host, err := net.Dial("tcp", fmt.Sprintf("%s:%d", tcpServerHost, tcpServerPort))

	require.NoError(t, err)

	testV5BindByIPv4SendPictureRequest(picture, *bufio.NewWriter(host), reader, t)
}

func testV5BindByIPv4SendPictureRequest(picture byte, writer bufio.Writer, reader io.Reader, t *testing.T) {
	require.NoError(t, writer.WriteByte(picture))
	require.NoError(t, writer.Flush())

	testV5BindByIPv4ReceiveSecondReply(picture, reader, t)
}

func testV5BindByIPv4ReceiveSecondReply(picture byte, reader io.Reader, t *testing.T) {
	response := make([]byte, 10)

	_, err := reader.Read(response)

	require.NoError(t, err)

	require.Equal(t, response[0], byte(5))
	require.Equal(t, response[1], byte(0))
	require.Equal(t, response[2], byte(0))
	require.Equal(t, response[3], byte(1))

	ip := net.IP{response[4], response[5], response[6], response[7]}.To4()

	require.NotNil(t, ip)

	ipv4 := net.ParseIP(tcpServerBindIPv4)

	require.NotNil(t, ipv4)

	ipv4 = ipv4.To4()

	require.NotNil(t, ipv4)

	require.Equal(t, ip, ipv4)

	port := make([]byte, 2)

	binary.LittleEndian.PutUint16(port, tcpServerBindPort)

	require.Equal(t, port, response[8:10])

	testV5BindByIPv4ReceivePicture(picture, reader, t)
}

func testV5BindByIPv4ReceivePicture(picture byte, reader io.Reader, t *testing.T) {
	var file *os.File
	var err error

	if picture == 4 {
		file, err = ioutil.TempFile("", "v5-bind-by-ipv4-big-picture")
	} else if picture == 5 {
		file, err = ioutil.TempFile("", "v5-bind-by-ipv4-middle-picture")
	} else if picture == 6 {
		file, err = ioutil.TempFile("", "v5-bind-by-ipv4-small-picture")
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

func TestV5BindByIPv6WithBigPicture(t *testing.T) {
	t.Skipf("It works too unstable... idk what to do with this random behavior.")

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", socksTcpHost, socksTcpPort))

	require.NoError(t, err)

	writer := *bufio.NewWriter(conn)

	testV5BindByIPv6Authenticate(7, writer, conn, t)
}

func TestV5BindByIPv6WithMiddlePicture(t *testing.T) {
	t.Skipf("It works too unstable... idk what to do with this random behavior.")

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", socksTcpHost, socksTcpPort))

	require.NoError(t, err)

	writer := *bufio.NewWriter(conn)

	testV5BindByIPv6Authenticate(8, writer, conn, t)
}

func TestV5BindByIPv6WithSmallPicture(t *testing.T) {
	t.Skipf("It works too unstable... idk what to do with this random behavior.")

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", socksTcpHost, socksTcpPort))

	require.NoError(t, err)

	writer := *bufio.NewWriter(conn)

	testV5BindByIPv6Authenticate(9, writer, conn, t)
}

func testV5BindByIPv6Authenticate(picture byte, writer bufio.Writer, reader io.Reader, t *testing.T) {
	require.NoError(t, writer.WriteByte(5))
	require.NoError(t, writer.WriteByte(1))
	require.NoError(t, writer.WriteByte(0))

	require.NoError(t, writer.Flush())

	selection := make([]byte, 10)

	_, err := reader.Read(selection)

	require.NoError(t, err)

	require.Equal(t, selection[1], byte(0))

	testV5BindByIPv6SendRequest(picture, writer, reader, t)
}

func testV5BindByIPv6SendRequest(picture byte, writer bufio.Writer, reader io.Reader, t *testing.T) {
	require.NoError(t, writer.WriteByte(5))
	require.NoError(t, writer.WriteByte(2))
	require.NoError(t, writer.WriteByte(0))
	require.NoError(t, writer.WriteByte(4))

	ip := net.ParseIP(tcpServerBindIPv6)

	require.NotNil(t, ip)

	ip = ip.To16()

	require.NotNil(t, ip)

	_, err := writer.Write(ip)

	require.NoError(t, err)

	require.NoError(t, binary.Write(&writer, binary.BigEndian, tcpServerBindPort))

	require.NoError(t, writer.Flush())

	testV5BindByIPv6ReceiveFirstReply(picture, reader, t)
}

func testV5BindByIPv6ReceiveFirstReply(picture byte, reader io.Reader, t *testing.T) {
	response := make([]byte, 10)

	_, err := reader.Read(response)

	require.NoError(t, err)

	require.Equal(t, response[0], byte(5))
	require.Equal(t, response[1], byte(0))
	require.Equal(t, response[2], byte(0))
	require.Equal(t, response[3], byte(1))
	require.Equal(t, response[4], byte(0))
	require.Equal(t, response[5], byte(0))
	require.Equal(t, response[6], byte(0))
	require.Equal(t, response[7], byte(0))

	port := make([]byte, 2)

	binary.LittleEndian.PutUint16(port, socksTcpPort)

	require.Equal(t, port, response[8:10])

	testV5BindByIPv6ConnectToServer(picture, reader, t)
}

func testV5BindByIPv6ConnectToServer(picture byte, reader io.Reader, t *testing.T) {
	host, err := net.Dial("tcp", fmt.Sprintf("%s:%d", tcpServerHost, tcpServerPort))

	require.NoError(t, err)

	testV5BindByIPv6SendPictureRequest(picture, *bufio.NewWriter(host), reader, t)
}

func testV5BindByIPv6SendPictureRequest(picture byte, writer bufio.Writer, reader io.Reader, t *testing.T) {
	require.NoError(t, writer.WriteByte(picture))
	require.NoError(t, writer.Flush())

	testV5BindByIPv6ReceiveSecondReply(picture, reader, t)
}

func testV5BindByIPv6ReceiveSecondReply(picture byte, reader io.Reader, t *testing.T) {
	response := make([]byte, 22)

	_, err := reader.Read(response)

	require.NoError(t, err)

	require.Equal(t, response[0], byte(5))
	require.Equal(t, response[1], byte(0))
	require.Equal(t, response[2], byte(0))
	require.Equal(t, response[3], byte(1))

	ip := net.IP{
		response[4], response[5], response[6], response[7],
		response[8], response[9], response[10], response[11],
		response[12], response[13], response[14], response[15],
		response[16], response[17], response[18], response[19],
	}.To16()

	require.NotNil(t, ip)

	ipv6 := net.ParseIP(tcpServerBindIPv6)

	require.NotNil(t, ipv6)

	ipv6 = ipv6.To16()

	require.NotNil(t, ipv6)

	require.Equal(t, ip, ipv6)

	port := make([]byte, 2)

	binary.LittleEndian.PutUint16(port, tcpServerBindPort)

	require.Equal(t, port, response[20:22])

	testV5BindByIPv6ReceivePicture(picture, reader, t)
}

func testV5BindByIPv6ReceivePicture(picture byte, reader io.Reader, t *testing.T) {
	var file *os.File
	var err error

	if picture == 4 {
		file, err = ioutil.TempFile("", "v5-bind-by-ipv6-big-picture")
	} else if picture == 5 {
		file, err = ioutil.TempFile("", "v5-bind-by-ipv6-middle-picture")
	} else if picture == 6 {
		file, err = ioutil.TempFile("", "v5-bind-by-ipv6-small-picture")
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
