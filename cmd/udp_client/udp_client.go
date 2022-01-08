package main

import (
	"fmt"
	"net"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:6666")

	if err != nil {
		panic(err)
	}

	conn, dialErr := net.DialUDP("udp", nil, addr)

	if dialErr != nil {
		panic(dialErr)
	}

	fmt.Println(conn.RemoteAddr())
	fmt.Println(conn.LocalAddr())

	sendFirst(conn)

	readFirst(conn)

	sendSecond(conn)

	readSecond(conn)
}

func sendFirst(conn net.Conn) {
	_, err := conn.Write([]byte("ok"))

	if err != nil {
		panic(err)
	}
}

func sendSecond(conn net.Conn) {
	_, err := conn.Write([]byte("ko"))

	if err != nil {
		panic(err)
	}
}

func readFirst(conn net.PacketConn) {

	buf := make([]byte, 1024)

	n, address, readErr := conn.ReadFrom(buf)

	if readErr != nil {
		panic(readErr)
	}

	fmt.Printf("Packet from %s, bytes: %s\n", address, buf[:n])
}

func readSecond(conn net.PacketConn) {
	buf := make([]byte, 1024)

	n, address, readErr := conn.ReadFrom(buf)

	if readErr != nil {
		panic(readErr)
	}

	fmt.Printf("Packet from %s, bytes: %s\n", address, buf[:n])
}
