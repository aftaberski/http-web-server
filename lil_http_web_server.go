package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"strings"
)

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

		// Get the request
		// TODO switch to bufio.Scanner
		scanner := bufio.NewScanner(conn)
		scanner.Scan()
		requestStr := scanner.Text()
		uri := strings.Split(requestStr, " ")[1]

		var headers string
		var body []byte

		switch uri {
		case "/hello":
			headers = "HTTP/1.1 200 OK\r\n\r\n"
			body = []byte("Hello, World!")
		case "/puppy":
			body, err = ioutil.ReadFile("puppy.jpg")
			if err != nil {
				log.Fatal(err)
			}

			contentLength := len(body)
			// TODO make cute;
			headers = fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: image/jpeg\r\nContent-Length: %d\r\n\r\n", contentLength)
		default:
			headers = "HTTP/1.1 404 Not Found\r\n\r\n"
			body = []byte("404 There's nothing here!")
		}

		fmt.Println("lol")

		fmt.Println(headers)
		// Handle the connection in a new goroutine.
		// The loop then returns to accepting, so that
		// multiple connections may be served concurrently.
		go func(c net.Conn) {
			// Write response to connection
			c.Write([]byte(headers))
			c.Write(body)
			// Shut down the connection.
			c.Close()
		}(conn)
	}
}
