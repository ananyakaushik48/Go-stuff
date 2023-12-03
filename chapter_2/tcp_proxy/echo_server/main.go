package main

import (
	"io"
	"log"
	"net"
)

// This is the main handler fucntion that simply echoes any Data it receives
func echo(conn net.Conn) {
	defer conn.Close()
	// Create that stores the received data
	b := make([]byte, 512)
	// an infinite loop to listen any incoming data
	for {
		size, err := conn.Read(b[0:])
		if err != io.EOF {
			log.Println("Client disconnected")
		}
	}
}
