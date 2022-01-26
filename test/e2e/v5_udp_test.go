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
	"testing"
)

func TestV5UdpAssociationByDomainWithSmallPicture(t *testing.T) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", socksTcpHost, socksTcpPort))

	require.NoError(t, err)

	writer := *bufio.NewWriter(conn)

	testV5UdpAssociationByDomainAuthenticate(255, writer, conn, t)
}

func testV5UdpAssociationByDomainAuthenticate(picture byte, writer bufio.Writer, reader io.Reader, t *testing.T) {
	require.NoError(t, writer.WriteByte(5))
	require.NoError(t, writer.WriteByte(1))
	require.NoError(t, writer.WriteByte(0))

	require.NoError(t, writer.Flush())

	selection := make([]byte, 10)

	_, err := reader.Read(selection)

	require.NoError(t, err)

	require.Equal(t, selection[1], byte(0))

	testV5UdpAssociationByDomainSendRequest(picture, writer, reader, t)
}

func testV5UdpAssociationByDomainSendRequest(picture byte, writer bufio.Writer, reader io.Reader, t *testing.T) {
	require.NoError(t, writer.WriteByte(5))
	require.NoError(t, writer.WriteByte(3))
	require.NoError(t, writer.WriteByte(0))
	require.NoError(t, writer.WriteByte(3))

	host := []byte(udpServerHost)

	require.NoError(t, writer.WriteByte(byte(len(host))))

	_, err := writer.Write(host)

	require.NoError(t, err)

	require.NoError(t, binary.Write(&writer, binary.BigEndian, udpServerPort))

	require.NoError(t, writer.Flush())

	testV5UdpAssociationByDomainReceiveReply(picture, reader, t)
}

func testV5UdpAssociationByDomainReceiveReply(picture byte, reader io.Reader, t *testing.T) {
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

	binary.LittleEndian.PutUint16(port, socksUdpPort)

	require.Equal(t, port, response[8:10])

	testV5UdpAssociationByDomainDialUdp(picture, t)
}

func testV5UdpAssociationByDomainDialUdp(picture byte, t *testing.T) {
	addr, lookupErr := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", socksUdpHost, socksUdpPort))

	require.NoError(t, lookupErr)

	conn, err := net.DialUDP("udp", nil, addr)

	require.NoError(t, err)

	testV5UdpAssociationByDomainSendPacket(picture, *bufio.NewWriter(conn), conn, t)
}

func testV5UdpAssociationByDomainSendPacket(picture byte, writer bufio.Writer, packet net.PacketConn, t *testing.T) {
	require.NoError(t, writer.WriteByte(0))
	require.NoError(t, writer.WriteByte(0))
	require.NoError(t, writer.WriteByte(0))
	require.NoError(t, writer.WriteByte(3))

	host := []byte(udpServerHost)

	require.NoError(t, writer.WriteByte(byte(len(host))))

	_, err := writer.Write(host)

	require.NoError(t, err)

	require.NoError(t, binary.Write(&writer, binary.BigEndian, udpServerPort))

	require.NoError(t, writer.WriteByte(picture))

	require.NoError(t, writer.Flush())

	testV5UdpAssociationByDomainReceivePicture(packet, t)
}

func testV5UdpAssociationByDomainReceivePicture(packet net.PacketConn, t *testing.T) {
	file, err := ioutil.TempFile("", "v5-udp-association-by-domain-small-picture")

	require.NoError(t, err)

	h := sha256.New()

	writers := io.MultiWriter(file, h)

	buffer := make([]byte, 65535)

	i, _, readErr := packet.ReadFrom(buffer)

	require.NoError(t, readErr)

	_, err = writers.Write(buffer[:i])

	require.NoError(t, err)

	require.Equal(t, smallPictureHash, fmt.Sprintf("%x", h.Sum(nil)))

	h.Reset()

	require.NoError(t, file.Close())
}

func TestV5UdpAssociationByIPv4WithSmallPicture(t *testing.T) {
	//conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", socksTcpHost, socksTcpPort))
	//
	//require.NoError(t, err)
	//
	//writer := *bufio.NewWriter(conn)
	//
	//testV5UdpAssociationByIPv4Authenticate(255, writer, conn, t)
}

func testV5UdpAssociationByIPv4Authenticate(picture byte, writer bufio.Writer, reader io.Reader, t *testing.T) {
	require.NoError(t, writer.WriteByte(5))
	require.NoError(t, writer.WriteByte(1))
	require.NoError(t, writer.WriteByte(0))

	require.NoError(t, writer.Flush())

	selection := make([]byte, 10)

	_, err := reader.Read(selection)

	require.NoError(t, err)

	require.Equal(t, selection[1], byte(0))

	testV5UdpAssociationByIPv4SendRequest(picture, writer, reader, t)
}

func testV5UdpAssociationByIPv4SendRequest(picture byte, writer bufio.Writer, reader io.Reader, t *testing.T) {
	require.NoError(t, writer.WriteByte(5))
	require.NoError(t, writer.WriteByte(3))
	require.NoError(t, writer.WriteByte(0))
	require.NoError(t, writer.WriteByte(1))

	ip := net.ParseIP(udpServerIPv4)

	require.NotNil(t, ip)

	ip = ip.To4()

	require.NotNil(t, ip)

	_, err := writer.Write(ip)

	require.NoError(t, err)

	require.NoError(t, binary.Write(&writer, binary.BigEndian, udpServerPort))

	require.NoError(t, writer.Flush())

	testV5UdpAssociationByIPv4ReceiveReply(picture, reader, t)
}

func testV5UdpAssociationByIPv4ReceiveReply(picture byte, reader io.Reader, t *testing.T) {
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

	binary.LittleEndian.PutUint16(port, socksUdpPort)

	require.Equal(t, port, response[8:10])

	testV5UdpAssociationByIPv4DialUdp(picture, t)
}

func testV5UdpAssociationByIPv4DialUdp(picture byte, t *testing.T) {
	addr, lookupErr := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", socksUdpHost, socksUdpPort))

	require.NoError(t, lookupErr)

	conn, err := net.DialUDP("udp", nil, addr)

	require.NoError(t, err)

	testV5UdpAssociationByIPv4SendPacket(picture, *bufio.NewWriter(conn), conn, t)
}

func testV5UdpAssociationByIPv4SendPacket(picture byte, writer bufio.Writer, packet net.PacketConn, t *testing.T) {
	require.NoError(t, writer.WriteByte(0))
	require.NoError(t, writer.WriteByte(0))
	require.NoError(t, writer.WriteByte(0))
	require.NoError(t, writer.WriteByte(3))

	host := []byte(udpServerHost)

	require.NoError(t, writer.WriteByte(byte(len(host))))

	_, err := writer.Write(host)

	require.NoError(t, err)

	require.NoError(t, binary.Write(&writer, binary.BigEndian, udpServerPort))

	require.NoError(t, writer.WriteByte(picture))

	require.NoError(t, writer.Flush())

	testV5UdpAssociationByIPv4ReceivePicture(packet, t)
}

func testV5UdpAssociationByIPv4ReceivePicture(packet net.PacketConn, t *testing.T) {
	file, err := ioutil.TempFile("", "v5-udp-association-by-ipv4-small-picture")

	require.NoError(t, err)

	h := sha256.New()

	writers := io.MultiWriter(file, h)

	buffer := make([]byte, 65535)

	i, _, readErr := packet.ReadFrom(buffer)

	require.NoError(t, readErr)

	_, err = writers.Write(buffer[:i])

	require.NoError(t, err)

	require.Equal(t, smallPictureHash, fmt.Sprintf("%x", h.Sum(nil)))

	h.Reset()

	require.NoError(t, file.Close())
}

func TestV5UdpAssociationByIPv6WithSmallPicture(t *testing.T) {
	//conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", socksTcpHost, socksTcpPort))
	//
	//require.NoError(t, err)
	//
	//writer := *bufio.NewWriter(conn)
	//
	//testV5UdpAssociationByIPv6Authenticate(255, writer, conn, t)
}

func testV5UdpAssociationByIPv6Authenticate(picture byte, writer bufio.Writer, reader io.Reader, t *testing.T) {
	require.NoError(t, writer.WriteByte(5))
	require.NoError(t, writer.WriteByte(1))
	require.NoError(t, writer.WriteByte(0))

	require.NoError(t, writer.Flush())

	selection := make([]byte, 10)

	_, err := reader.Read(selection)

	require.NoError(t, err)

	require.Equal(t, selection[1], byte(0))

	testV5UdpAssociationByIPv6SendRequest(picture, writer, reader, t)
}

func testV5UdpAssociationByIPv6SendRequest(picture byte, writer bufio.Writer, reader io.Reader, t *testing.T) {
	require.NoError(t, writer.WriteByte(5))
	require.NoError(t, writer.WriteByte(3))
	require.NoError(t, writer.WriteByte(0))
	require.NoError(t, writer.WriteByte(4))

	ip := net.ParseIP(udpServerIPv6)

	require.NotNil(t, ip)

	ip = ip.To16()

	require.NotNil(t, ip)

	_, err := writer.Write(ip)

	require.NoError(t, err)

	require.NoError(t, binary.Write(&writer, binary.BigEndian, udpServerPort))

	require.NoError(t, writer.Flush())

	testV5UdpAssociationByIPv6ReceiveReply(picture, reader, t)
}

func testV5UdpAssociationByIPv6ReceiveReply(picture byte, reader io.Reader, t *testing.T) {
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

	binary.LittleEndian.PutUint16(port, socksUdpPort)

	require.Equal(t, port, response[8:10])

	testV5UdpAssociationByIPv6DialUdp(picture, t)
}

func testV5UdpAssociationByIPv6DialUdp(picture byte, t *testing.T) {
	addr, lookupErr := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", socksUdpHost, socksUdpPort))

	require.NoError(t, lookupErr)

	conn, err := net.DialUDP("udp", nil, addr)

	require.NoError(t, err)

	testV5UdpAssociationByIPv6SendPacket(picture, *bufio.NewWriter(conn), conn, t)
}

func testV5UdpAssociationByIPv6SendPacket(picture byte, writer bufio.Writer, packet net.PacketConn, t *testing.T) {
	require.NoError(t, writer.WriteByte(0))
	require.NoError(t, writer.WriteByte(0))
	require.NoError(t, writer.WriteByte(0))
	require.NoError(t, writer.WriteByte(3))

	host := []byte(udpServerHost)

	require.NoError(t, writer.WriteByte(byte(len(host))))

	_, err := writer.Write(host)

	require.NoError(t, err)

	require.NoError(t, binary.Write(&writer, binary.BigEndian, udpServerPort))

	require.NoError(t, writer.WriteByte(picture))

	require.NoError(t, writer.Flush())

	testV5UdpAssociationByIPv6ReceivePicture(packet, t)
}

func testV5UdpAssociationByIPv6ReceivePicture(packet net.PacketConn, t *testing.T) {
	file, err := ioutil.TempFile("", "v5-udp-association-by-ipv6-small-picture")

	require.NoError(t, err)

	h := sha256.New()

	writers := io.MultiWriter(file, h)

	buffer := make([]byte, 65535)

	i, _, readErr := packet.ReadFrom(buffer)

	require.NoError(t, readErr)

	_, err = writers.Write(buffer[:i])

	require.NoError(t, err)

	require.Equal(t, smallPictureHash, fmt.Sprintf("%x", h.Sum(nil)))

	h.Reset()

	require.NoError(t, file.Close())
}
