package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"strings"
)

func handleConnection(conn net.Conn) {

	defer conn.Close()

	// TODO: pull out parse header code into a function that takes a scanner
	// Avoid panics by checking explicitly if the scan returned
	reader := bufio.NewReader(conn)
	status, err := reader.ReadString('\n')
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(len(status))
	uri := strings.Split(status, " ")[1]

	var headers string
	var body []byte

	switch uri {
	case "/hello":
		headers = "HTTP/1.1 200 OK\r\n\r\n"
		body = []byte("Hello, World!")
	case "/puppy":
		body, _ = ioutil.ReadFile("puppy.jpg")
		// if err != nil {
		// 	log.Fatal(err)
		// }

		contentLength := len(body)
		// TODO make cute;
		headers = fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: image/jpeg\r\nContent-Length: %d\r\nConnection: close\r\n\r\n", contentLength)
	default:
		headers = "HTTP/1.1 404 Not Found\r\n\r\n"
		body = []byte("404 There's nothing here!")
	}

	fmt.Println(headers)
	// Handle the connection in a new goroutine.
	// The loop then returns to accepting, so that
	// multiple connections may be served concurrently.

	// Write response to connection
	conn.Write([]byte(headers))
	conn.Write(body)
	// Shut down the connection.

}

func main() {
	l, err := net.Listen("tcp", ":8888")
	fmt.Printf("Listening at :8888...")
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
		go handleConnection(conn)
	}
}
