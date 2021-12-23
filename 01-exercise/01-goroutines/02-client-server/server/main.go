package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	// TODO: write server program to handle concurrent client connections.
	l, err := net.Listen("tcp", "localhost:8000")

	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := l.Accept()

		if err != nil {
			log.Fatal("conn", err)
			continue
		}

		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, "response from server\n")
		if err != nil {
			return
		}
		time.Sleep(time.Second)
	}
}
