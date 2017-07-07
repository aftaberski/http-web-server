package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":8888")
	fmt.Println("Listening...")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	for {
		// Wait for a connection.
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		response := `HTTP/1.1 200 OK

    Hello, World!
    `

		fmt.Println(response)
		// Handle the connection in a new goroutine.
		// The loop then returns to accepting, so that
		// multiple connections may be served concurrently.
		go func(c net.Conn) {
			// Write response to connection
			c.Write([]byte(response))
			// Shut down the connection.
			c.Close()
		}(conn)
	}
}
