package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"./http"
)

func helloHandler(request http.Request) (*http.Response, error) {
	headers := make([]string, 0)
	body := []byte("Hello, World!")

	return http.NewResponse("HTTP/1.1", "200", "OK", headers, body)
}

func puppyHandler(request http.Request) (*http.Response, error) {
	body, err := ioutil.ReadFile("puppy.jpg")
	if err != nil {
		log.Fatal(err)
	}

	contentLength := len(body)
	headers := []string{"Content-Type: image/jpeg", fmt.Sprintf("Content-Length: %d", contentLength)}
	return http.NewResponse("HTTP/1.1", "200", "OK", headers, body)
}

func main() {
	server := http.NewServer()
	server.AddRoute(http.GET, "/hello", helloHandler)
	server.AddRoute(http.GET, "/puppy", puppyHandler)
	server.Run(8888)
}
