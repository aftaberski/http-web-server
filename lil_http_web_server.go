package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"strings"
)

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", ":8888")
	if err != nil {
		log.Fatal(err)
	}

	l, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Listening at :8888...")
	defer l.Close()

	for {
		fmt.Println("Waiting for connection...............")
		// Wait for a connection.
		conn, err := l.AcceptTCP()
		if err != nil {
			fmt.Println("CONNECTION ERROR:")
			log.Fatal(err)
		}

		var reqBody = make([]byte, 1024)
		bytesRead, err := conn.Read(reqBody)
		if err != nil {
			fmt.Println("READ ERROR:")
			log.Fatal(err)
		}
		strReqBody := string(reqBody[:bytesRead])
		reqStringArr := strings.Split(strReqBody, "\n")

		uri := strings.Split(reqStringArr[0], " ")[1]
		fmt.Println(uri)

		var headers string
		var body []byte

		switch uri {
		case "/hello":
			headers = "HTTP/1.1 200 OK\r\n\r\n"
			body = []byte("Hello, World!")
			go handleConn(conn, headers, body, rand.Intn(1000))
		case "/puppy":
			body, err = ioutil.ReadFile("puppy.jpg")
			if err != nil {
				fmt.Println("READ FILE ERROR:")
				log.Fatal(err)
			}

			contentLength := len(body)
			headers = fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: image/jpeg\r\nContent-Length: %d\r\n\r\n", contentLength)
			go handleConn(conn, headers, body, rand.Intn(1000))
		default:
			headers = "HTTP/1.1 404 Not Found\r\n\r\n"
			body = []byte("404 There's nothing here!")
			go handleConn(conn, headers, body, rand.Intn(1000))
		}

		fmt.Println("HEADERS:", headers)
	}
}

func handleConn(c net.Conn, headers string, body []byte, random int) {
	fmt.Println("HANDLE CONN:", random)
	// Shut down the connection.
	defer c.Close()
	// Write response to connection
	_, err := c.Write([]byte(headers))
	if err != nil {
		log.Fatal(err)
	}

	_, err = c.Write(body)
	if err != nil {
		log.Fatal(err)
	}
}
