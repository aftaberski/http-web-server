package main

import "./http"

func helloHandler(request http.Request) (*http.Response, error) {
	headers := make([]string, 0)
	body := []byte("Hello, World!")

	return http.NewResponse("HTTP/1.1", "200", "OK", headers, body)
}

func main() {
	server := http.NewServer()
	server.AddRoute(http.GET, "/hello", helloHandler)
	server.Run(8888)
}
