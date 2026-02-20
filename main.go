package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) < 4 {
		log.Fatal("needs more args : <address:port> <method> <route> ")
	}

	address := os.Args[1]
	meth := os.Args[2]
	route := os.Args[3]

	tcpConnObject, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal(err)
	}

	_, err = fmt.Fprintf(tcpConnObject, "%s %s HTTP/1.0\r\n\r\n", meth, route)
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
