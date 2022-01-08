package main

import (
	"io/ioutil"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:1080")

	if err != nil {
		panic(err)
	}

	file, fErr := ioutil.ReadFile("/home/chop1k/Pictures/Memes/memes_37/xuBY1ksbsUw.jpg")

	if fErr != nil {
		panic(fErr)
	}

	_, err = conn.Write(file)

	if err != nil {
		panic(err)
	}

	//for {
	//	buf := make([]byte, 512)
	//
	//	i, readErr := file.Read(buf)
	//
	//	if readErr != nil && i <= 0 {
	//		panic(readErr)
	//	}
	//
	//	_, writeErr := conn.Write(buf[:i])
	//
	//	if writeErr != nil {
	//		panic(writeErr)
	//	}
	//}
}
