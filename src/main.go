package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"strings"

	"./http"
)

func handleConnection(conn net.Conn) {

	defer conn.Close()

	// TODO: make HTTPRequest type
	reader := bufio.NewReader(conn)
	req := http.NewRequest(*reader)
	fmt.Println(req.Method)
	status, err := reader.ReadString('\n')
	if err != nil {
		log.Println(err)
	}

	// This is a fix for a bug occuring with a Chrome browser
	// TODO RCA this bug!
	// But for now...its killed ~4 hrs of time so I'm choosing to
	// use this workaround
	if status != "" {
		uri := strings.Split(status, " ")[1]

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
			// TODO make HTTPResponse type
			headers = fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: image/jpeg\r\nContent-Length: %d\r\nConnection: close\r\n\r\n", contentLength)
		default:
			headers = "HTTP/1.1 404 Not Found\r\n\r\n"
			body = []byte("404 There's nothing here!")
		}

		fmt.Println(headers)

		// Write response to connection
		conn.Write([]byte(headers))
		conn.Write(body)
	}
}

func main() {
	l, err := net.Listen("tcp", ":8888")
	fmt.Println("Listening at :8888...")
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
		// Handle the connection in a new goroutine.
		// The loop then returns to accepting, so that
		// multiple connections may be served concurrently.
		go handleConnection(conn)
	}
}
