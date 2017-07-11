package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"

	"./http"
)

func handleConnection(conn net.Conn) {

	defer conn.Close()

	// TODO: make HTTPRequest type
	// scanner := bufio.NewScanner(conn)
	req, err := http.NewRequest(conn)
	if err != nil {
		log.Fatal(err)
	}

	var headers string
	var body []byte
	if err != nil {
		log.Fatal(err)
	}

	// get route if exists and return a handler
	// Will need something like:
	// {"GET": {"/route-name": handlerFn,
	//          "/other": otherHandlerFn}}
	switch req.URI {
	case "/hello":
		// this is really the handler ->
		// Takes a Request and returns a Response
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

func helloHandler(request http.Request) (*http.Response, error) {
	headers := make([]string, 0)
	body := []byte("Hello World")

	return http.NewResponse("HTTP/1.1", "200", "OK", headers, body)
}

func main() {
	server := http.NewServer()
	server.AddRoute(http.GET, "/hello", helloHandler)
	server.Run(8888)

	// l, err := net.Listen("tcp", ":8888")
	// fmt.Println("Listening at :8888...")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer l.Close()
	// for {
	// 	// Wait for a connection.
	// 	conn, err := l.Accept()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	// Handle the connection in a new goroutine.
	// 	// The loop then returns to accepting, so that
	// 	// multiple connections may be served concurrently.
	// 	go handleConnection(conn)
	// }
}
