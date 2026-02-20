package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	tcpConnObject, err := net.Dial("tcp", ":10055")
	if err != nil {
		log.Fatal(err)
	}

	_, err = tcpConnObject.Write([]byte("GET /marks/login?= HTTP/1.0\r\n\r\n"))
	if err != nil {
		log.Fatal("Error writing to server:", err)
	}

	buffer := make([]byte, 500)
	for {
		n, err := tcpConnObject.Read(buffer)
		if n > 0 {
			fmt.Println(string(buffer[:n]))
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("Error reading response:", err)
		}
	}

	defer tcpConnObject.Close()
}
