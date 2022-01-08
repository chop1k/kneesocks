package e2e

import (
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
