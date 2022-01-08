package main

import (
	"net"
	"os"
)

var (
	smallPicturePath string
)

func init() {
	setPictureEnv()
}

func setPictureEnv() {
	smallPicture, oo := os.LookupEnv("small_sized_picture_path")

	if !oo {
		panic("Small picture missed. ")
	}

	smallPicturePath = smallPicture
}

func main() {
	go listen("[::1]:9999")
	listen("127.0.0.1:9999")
}

func listen(addr string) {
	listener, err := net.ListenPacket("udp", addr)

	if err != nil {
		panic(err)
	}

	for {
		buf := make([]byte, 1024)

		_, addr, err := listener.ReadFrom(buf)

		if err != nil {
			continue
		}

		go handlePacket(buf[0], addr, listener)
	}
}

func handlePacket(picture byte, addr net.Addr, conn net.PacketConn) {
	if picture == 255 {
		sendPicture(addr, conn)
	}
}

func sendPicture(addr net.Addr, conn net.PacketConn) {
	file, err := os.Open(smallPicturePath)

	if err != nil {
		return
	}

	buffer := make([]byte, 60000)

	i, readErr := file.Read(buffer)

	if readErr != nil {
		_ = file.Close()

		return
	}

	_, _ = conn.WriteTo(buffer[:i], addr)

	_ = file.Close()
}
