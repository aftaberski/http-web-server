# HTTP Server in Go

An HTTP Server written in Go using TCP sockets

## Usage

```golang
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
```

`NewServer()`: creates a new `Server` instance 

`AddRoutes()`: adds routes to your server instance. The method takes an HTTP `method` type (currently either `GET` or `POST`), a route, and a `Handler`.

`Handler`: a function type that takes a `Request` and returns a reference to the `Response` and an error.

`NewResponse()`: creates a new `Response` instance. Takes the HTTP protocol (currently only "HTTP/1.1"), status code, status message, array of headers, and byte array for the body.

## Future Improvements
* Allow chunked transfer encoding
* Implement connection pooling
* Allow middleware
* Write tests