package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
)

var (
	bigPicturePath    string
	middlePicturePath string
	smallPicturePath  string
	socksHost         string
	socksPort         uint16
	bindPort          uint16
)

func init() {
	setSocksEnv()
	setPictureEnv()
	setBindEnv()
}

func setSocksEnv() {
	host, ok := os.LookupEnv("socks_host")

	if !ok {
		panic("Socks address is not specified. ")
	}

	port, ko := os.LookupEnv("socks_port")

	if !ko {
		panic("Socks port is not specified. ")
	}

	portNumber, err := strconv.Atoi(port)

	if err != nil {
		panic(err)
	}

	socksHost = host
	socksPort = uint16(portNumber)
}

func setPictureEnv() {
	bigPicture, ok := os.LookupEnv("big_sized_picture_path")

	if !ok {
		panic("Big picture missed. ")
	}

	middlePicture, ko := os.LookupEnv("middle_sized_picture_path")

	if !ko {
		panic("Middle picture missed. ")
	}

	smallPicture, oo := os.LookupEnv("small_sized_picture_path")

	if !oo {
		panic("Small picture missed. ")
	}

	bigPicturePath = bigPicture
	middlePicturePath = middlePicture
	smallPicturePath = smallPicture
}

func setBindEnv() {
	port, ko := os.LookupEnv("bind_port")

	if !ko {
		panic("Bind port is not specified. ")
	}

	portNumber, err := strconv.Atoi(port)

	if err != nil {
		panic(err)
	}

	bindPort = uint16(portNumber)
}

func main() {
	go listen("[::1]:8888")
	listen("127.0.0.1:8888")
}

func listen(addr string) {
	listener, err := net.Listen("tcp", addr)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Test server listening on %s.\n", addr)

	for {
		conn, err := listener.Accept()

		if err != nil {
			panic(err)
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	buffer := make([]byte, 1)

	i, err := conn.Read(buffer)

	if err != nil {
		_ = conn.Close()

		return
	}

	if i != 1 {
		_ = conn.Close()

		return
	}

	fmt.Printf("Picture request #%d.\n ", buffer[0])

	if buffer[0] == 1 {
		sendPicture(1, conn)
	} else if buffer[0] == 2 {
		sendPicture(2, conn)
	} else if buffer[0] == 3 {
		sendPicture(3, conn)
	} else if buffer[0] == 4 {
		connectAndSendPicture(4)
	} else if buffer[0] == 5 {
		connectAndSendPicture(5)
	} else if buffer[0] == 6 {
		connectAndSendPicture(6)
	} else if buffer[0] == 7 {
		connectAndSendPicture(7)
	} else if buffer[0] == 8 {
		connectAndSendPicture(8)
	} else if buffer[0] == 9 {
		connectAndSendPicture(9)
	}

	_ = conn.Close()
}

func sendPicture(picture byte, writer io.Writer) {
	var file *os.File
	var err error

	if picture == 1 {
		file, err = os.Open(bigPicturePath)
	} else if picture == 2 {
		file, err = os.Open(middlePicturePath)
	} else if picture == 3 {
		file, err = os.Open(smallPicturePath)
	} else {
		return
	}

	if err != nil {
		return
	}

	_, _ = io.Copy(writer, file)

	_ = file.Close()
}

func connectAndSendPicture(picture byte) {
	var file *os.File
	var err error

	if picture == 4 || picture == 7 {
		file, err = os.Open(bigPicturePath)
	} else if picture == 5 || picture == 8 {
		file, err = os.Open(middlePicturePath)
	} else if picture == 6 || picture == 9 {
		file, err = os.Open(smallPicturePath)
	} else {
		return
	}

	if err != nil {
		return
	}

	var address string

	if picture == 7 || picture == 8 || picture == 9 {
		address = fmt.Sprintf("[%s]:%d", "::1", bindPort)
	} else {
		address = fmt.Sprintf("%s:%d", "127.0.0.1", bindPort)
	}

	lAddr, lErr := net.ResolveTCPAddr("tcp", address)

	if lErr != nil {
		return
	}

	rAddr, rErr := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", socksHost, socksPort))

	if rErr != nil {
		return
	}

	host, dialErr := net.DialTCP("tcp", lAddr, rAddr)

	if dialErr != nil {
		_ = file.Close()

		return
	}

	for {
		buf := make([]byte, 512)

		i, err := file.Read(buf)

		if err != nil {
			break
		}

		_, writeErr := host.Write(buf[:i])

		if writeErr != nil {
			break
		}
	}

	_ = host.Close()
	_ = file.Close()
}
