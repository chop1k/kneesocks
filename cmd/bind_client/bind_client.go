package main

import (
	"bytes"
	"encoding/binary"
	"io"
	"net"
	"os"
)

func main() {
	client, err := net.Dial("tcp", "localhost:1080")

	if err != nil {
		panic(err)
	}

	sendBind(client)
	receiveFirst(client)
	receiveSecond(client)
	receiveData(client)
}

func sendBind(client net.Conn) {
	buf := bytes.Buffer{}

	buf.WriteByte(4)
	buf.WriteByte(2)

	var hui = uint16(58512)

	err := binary.Write(&buf, binary.BigEndian, hui)

	buf.Write([]byte{0, 0, 0, 1, 0, 108, 111, 99, 97, 108, 104, 111, 115, 116, 0})

	if err != nil {
		panic(err)
	}

	_, err = client.Write(buf.Bytes())

	if err != nil {
		panic(err)
	}
}

func receiveFirst(client net.Conn) {
	buf := make([]byte, 512)

	i, err := client.Read(buf)

	if err != nil {
		panic(err)
	}

	buf = buf[:i]

	if buf[0] != 0 {
		panic("a")
	}

	if buf[1] != 90 {
		panic("b")
	}
}

func receiveSecond(client net.Conn) {
	buf := make([]byte, 512)

	i, err := client.Read(buf)

	if err != nil {
		panic(err)
	}

	buf = buf[:i]

	if buf[0] != 0 {
		panic("a")
	}

	if buf[1] != 90 {
		panic("b")
	}
}

func receiveData(client net.Conn) {
	file, err := os.Create("/home/chop1k/test.jpg")

	if err != nil {
		panic(file)
	}

	_, err = io.Copy(file, client)

	if err != nil {
		panic(err)
	}

	//for {
	//	buf := make([]byte, 512)
	//
	//	i, readErr := client.Read(buf)
	//
	//	if readErr != nil {
	//		panic(readErr)
	//	}
	//
	//	println(buf[:i])
	//
	//	_, err = file.Write(buf[:i])
	//
	//	if err != nil {
	//		panic(err)
	//	}
	//}
}