package e2e

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"github.com/stretchr/testify/require"
	"io"
	"io/ioutil"
	"net"
	"os"
	"strconv"
	"testing"
)

var (
	socksTcpHost      string
	socksTcpPort      uint16
	socksUdpHost      string
	socksUdpPort      uint16
	tcpServerHost     string
	tcpServerIPv4     string
	tcpServerIPv6     string
	tcpServerPort     uint16
	tcpServerBindHost string
	tcpServerBindIPv4 string
	tcpServerBindIPv6 string
	tcpServerBindPort uint16
	udpServerHost     string
	udpServerIPv4     string
	udpServerIPv6     string
	udpServerPort     uint16
	bigPictureHash    string
	middlePictureHash string
	smallPictureHash  string
)

func TestMain(m *testing.M) {
	testMainSetSocksTcpEnv()

	os.Exit(m.Run())
}

func testMainSetSocksTcpEnv() {
	host, ok := os.LookupEnv("socks_tcp_host")

	if !ok {
		panic("Socks tcp address is not specified. ")
	}

	port, ko := os.LookupEnv("socks_tcp_port")

	if !ko {
		panic("Socks tcp port is not specified. ")
	}

	portNumber, err := strconv.Atoi(port)

	if err != nil {
		panic(err)
	}

	socksTcpHost = host
	socksTcpPort = uint16(portNumber)

	testMainSetSocksUdpEnv()
}

func testMainSetSocksUdpEnv() {
	host, ok := os.LookupEnv("socks_udp_host")

	if !ok {
		panic("Socks udp address is not specified. ")
	}

	port, ko := os.LookupEnv("socks_udp_port")

	if !ko {
		panic("Socks udp port is not specified. ")
	}

	portNumber, err := strconv.Atoi(port)

	if err != nil {
		panic(err)
	}

	socksUdpHost = host
	socksUdpPort = uint16(portNumber)

	testMainSetTcpServerEnv()
}

func testMainSetTcpServerEnv() {
	host, ok := os.LookupEnv("tcp_server_host")

	if !ok {
		panic("Tcp server host address is not specified. ")
	}

	ipv4, oo := os.LookupEnv("tcp_server_ipv4")

	if !oo {
		panic("Tcp server ipv4 address is not specified. ")
	}

	ipv6, kk := os.LookupEnv("tcp_server_ipv6")

	if !kk {
		panic("Tcp server ipv6 address is not specified. ")
	}

	port, ko := os.LookupEnv("tcp_server_port")

	if !ko {
		panic("Tcp server port is not specified. ")
	}

	if len([]byte(host)) > 255 {
		panic("Host of the tcp server should not be > 255 bytes. ")
	}

	testPortNumber, err := strconv.Atoi(port)

	if err != nil {
		panic(err)
	}

	tcpServerHost = host
	tcpServerIPv4 = ipv4
	tcpServerIPv6 = ipv6
	tcpServerPort = uint16(testPortNumber)

	testMainSetTcpServerBindEnv()
}

func testMainSetTcpServerBindEnv() {
	host, ok := os.LookupEnv("tcp_server_bind_host")

	if !ok {
		panic("Tcp server bind address is not specified. ")
	}

	ipv4, oo := os.LookupEnv("tcp_server_bind_ipv4")

	if !oo {
		panic("Tcp server bind ipv4 address is not specified. ")
	}

	ipv6, kk := os.LookupEnv("tcp_server_bind_ipv6")

	if !kk {
		panic("Tcp server bind ipv6 address is not specified. ")
	}

	port, ko := os.LookupEnv("tcp_server_bind_port")

	if !ko {
		panic("Tcp server bind port is not specified. ")
	}

	if len([]byte(host)) > 255 {
		panic("Bind host of the tcp server should not be > 255 bytes. ")
	}

	testPortNumber, err := strconv.Atoi(port)

	if err != nil {
		panic(err)
	}

	tcpServerBindHost = host
	tcpServerBindIPv4 = ipv4
	tcpServerBindIPv6 = ipv6
	tcpServerBindPort = uint16(testPortNumber)

	testMainSetUdpServerEnv()
}

func testMainSetUdpServerEnv() {
	host, ok := os.LookupEnv("udp_server_host")

	if !ok {
		panic("Udp server address is not specified. ")
	}

	ipv4, oo := os.LookupEnv("udp_server_ipv4")

	if !oo {
		panic("Udp server ipv4 address is not specified. ")
	}

	ipv6, kk := os.LookupEnv("udp_server_ipv6")

	if !kk {
		panic("Udp server ipv6 address is not specified. ")
	}

	port, ko := os.LookupEnv("udp_server_port")

	if !ko {
		panic("Udp server port is not specified. ")
	}

	if len([]byte(host)) > 255 {
		panic("Host of the udp server should not be > 255 bytes. ")
	}

	testPortNumber, err := strconv.Atoi(port)

	if err != nil {
		panic(err)
	}

	udpServerHost = host
	udpServerIPv4 = ipv4
	udpServerIPv6 = ipv6
	udpServerPort = uint16(testPortNumber)

	testMainSetPicturesHashEnv()
}

func testMainSetPicturesHashEnv() {
	bigPicture, ok := os.LookupEnv("big_sized_picture_hash")

	if !ok {
		panic("Big picture hash is not specified. ")
	}

	middlePicture, ko := os.LookupEnv("middle_sized_picture_hash")

	if !ko {
		panic("Middle picture hash is not specified. ")
	}

	smallPicture, oo := os.LookupEnv("small_sized_picture_hash")

	if !oo {
		panic("Small picture hash is not specified. ")
	}

	bigPictureHash = bigPicture
	middlePictureHash = middlePicture
	smallPictureHash = smallPicture
}

func getLittleEndianPort(port uint16) []byte {
	portBinary := make([]byte, 2)

	binary.LittleEndian.PutUint16(portBinary, port)

	return portBinary
}

func getBigEndianPort(port uint16) []byte {
	portBinary := make([]byte, 2)

	binary.BigEndian.PutUint16(portBinary, port)

	return portBinary
}

func connectToServer(t *testing.T) net.Conn {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", socksTcpHost, socksTcpPort))

	require.NoError(t, err)

	return conn
}

func constructV4Request(command byte, ip string, port uint16, t *testing.T) []byte {
	ipv4 := net.ParseIP(ip)

	require.NotNil(t, ipv4)

	ipv4 = ipv4.To4()

	require.NotNil(t, ipv4)

	request := []byte{4, command}

	request = append(request, getBigEndianPort(port)...)
	request = append(request, ipv4...)
	request = append(request, 0)

	return request
}

func constructV4aRequest(command byte, address string, port uint16) []byte {
	request := []byte{4, command}

	request = append(request, getBigEndianPort(port)...)
	request = append(request, []byte{0, 0, 0, 255}...)
	request = append(request, []byte(address)...)
	request = append(request, 0)

	return request
}

func constructV5Request(command byte, addressType byte, address string, port uint16, t *testing.T) []byte {
	request := []byte{5, command, 0, addressType}

	if addressType == 1 {
		ipv4 := net.ParseIP(address)

		require.NotNil(t, ipv4)

		ipv4 = ipv4.To4()

		require.NotNil(t, ipv4)

		request = append(request, ipv4...)
	} else if addressType == 3 {
		request = append(request, []byte{byte(len(address))}...)
		request = append(request, []byte(address)...)
	} else if addressType == 4 {
		ipv6 := net.ParseIP(address)

		require.NotNil(t, ipv6)

		ipv6 = ipv6.To16()

		require.NotNil(t, ipv6)

		request = append(request, ipv6...)
	} else {
		t.Fatalf("Unknown address type %d.", addressType)
	}

	request = append(request, getBigEndianPort(port)...)

	return request
}

func sendV4Request(conn net.Conn, command byte, ip string, port uint16, t *testing.T) {
	_, err := conn.Write(constructV4Request(command, ip, port, t))

	require.NoError(t, err)
}

func sendV4aRequest(conn net.Conn, command byte, address string, port uint16, t *testing.T) {
	_, err := conn.Write(constructV4aRequest(command, address, port))

	require.NoError(t, err)
}

func sendV5Welcome(conn net.Conn, methods []byte, t *testing.T) {
	chunk := []byte{5, byte(len(methods))}

	chunk = append(chunk, methods...)

	_, err := conn.Write(chunk)

	require.NoError(t, err)
}

func compareV5Selection(conn net.Conn, method byte, t *testing.T) {
	response := make([]byte, 2)

	i, err := conn.Read(response)

	require.NoError(t, err)
	require.Equal(t, 2, i)

	require.Equal(t, []byte{5, method}, response)
}

func sendV5Password(conn net.Conn, name string, password string, t *testing.T) {
	sendV5Welcome(conn, []byte{2}, t)
	compareV5Selection(conn, 2, t)

	chunk := []byte{1, byte(len(name))}

	chunk = append(chunk, []byte(name)...)

	chunk = append(chunk, byte(len(password)))
	chunk = append(chunk, []byte(password)...)

	_, err := conn.Write(chunk)

	require.NoError(t, err)
}

func compareV5Password(conn net.Conn, t *testing.T) {
	expected := []byte{1, 0}

	response := make([]byte, 2)

	i, err := conn.Read(response)

	require.NoError(t, err)
	require.Equal(t, 2, i)

	require.Equal(t, expected, response)
}

func sendV5Request(conn net.Conn, command byte, addressType byte, address string, port uint16, t *testing.T) {
	sendV5Welcome(conn, []byte{0}, t)
	compareV5Selection(conn, 0, t)

	_, err := conn.Write(constructV5Request(command, addressType, address, port, t))

	require.NoError(t, err)
}

func compareV4Reply(conn net.Conn, ip string, port uint16, t *testing.T) {
	response := make([]byte, 8)

	i, err := conn.Read(response)

	require.NoError(t, err)
	require.Equal(t, 8, i)

	expected := []byte{0, 90}

	ipv4 := net.ParseIP(ip)

	require.NotNil(t, ipv4)

	ipv4 = ipv4.To4()

	require.NotNil(t, ipv4)

	expected = append(expected, getLittleEndianPort(port)...)
	expected = append(expected, ipv4...)

	require.Equal(t, expected, response)
}

func compareV5Reply(conn net.Conn, addressType byte, address string, port uint16, t *testing.T) {
	expected := []byte{5, 0, 0, addressType}

	if addressType == 1 {
		ipv4 := net.ParseIP(address)

		require.NotNil(t, ipv4)

		ipv4 = ipv4.To4()

		require.NotNil(t, ipv4)

		expected = append(expected, ipv4...)
	} else if addressType == 3 {
		expected = append(expected, []byte(address)...)
	} else if addressType == 4 {
		ipv6 := net.ParseIP(address)

		require.NotNil(t, ipv6)

		ipv6 = ipv6.To16()

		require.NotNil(t, ipv6)

		expected = append(expected, ipv6...)
	} else {
		t.Fatalf("Unknown address type %d.", addressType)
	}

	expected = append(expected, getLittleEndianPort(port)...)

	response := make([]byte, 512)

	i, err := conn.Read(response)

	require.NoError(t, err)

	require.Equal(t, expected, response[:i])
}

func sendPictureRequest(conn net.Conn, command byte, addressType byte, address string, port uint16, picture byte, t *testing.T) {
	request := []byte{command, picture, addressType}

	ip := net.ParseIP(address)

	require.NotNil(t, ip)

	if addressType == 1 {
		ip = ip.To4()

		require.NotNil(t, ip)
	} else if addressType == 2 {
		ip = ip.To16()

		require.NotNil(t, ip)
	}

	request = append(request, ip...)
	request = append(request, getBigEndianPort(port)...)

	_, err := conn.Write(request)

	require.NoError(t, err)
}

func comparePictureResponse(conn net.Conn, t *testing.T) {
	buffer := make([]byte, 6)

	i, err := conn.Read(buffer)

	require.NoError(t, err)

	require.Equal(t, 1, i)
	require.Equal(t, byte(0), buffer[0])
}

func connectToHost(address string, port uint16, t *testing.T) net.Conn {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", address, port))

	require.NoError(t, err)

	return conn
}

func comparePictures(conn net.Conn, prefix string, command string, picture byte, t *testing.T) {
	var file *os.File
	var err error

	if picture == 1 || picture == 4 {
		file, err = ioutil.TempFile("", fmt.Sprintf("%s-%s-%s.png", prefix, command, "big-picture"))
	} else if picture == 2 || picture == 5 {
		file, err = ioutil.TempFile("", fmt.Sprintf("%s-%s-%s.png", prefix, command, "middle-picture"))
	} else if picture == 3 || picture == 6 {
		file, err = ioutil.TempFile("", fmt.Sprintf("%s-%s-%s.png", prefix, command, "small-picture"))
	} else {
		t.Fatalf("Unknown picture %d. ", picture)
	}

	require.NoError(t, err)

	h := sha256.New()

	writers := io.MultiWriter(file, h)

	for {
		buffer := make([]byte, 512)

		i, err := conn.Read(buffer)

		if err != nil {
			break
		}

		_, err = writers.Write(buffer[:i])

		require.NoError(t, err)
	}

	if picture == 1 || picture == 4 {
		require.Equal(t, bigPictureHash, fmt.Sprintf("%x", h.Sum(nil)))
	} else if picture == 2 || picture == 5 {
		require.Equal(t, middlePictureHash, fmt.Sprintf("%x", h.Sum(nil)))
	} else if picture == 3 || picture == 6 {
		require.Equal(t, smallPictureHash, fmt.Sprintf("%x", h.Sum(nil)))
	}

	h.Reset()

	require.NoError(t, file.Close())
}
