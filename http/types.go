package http

import "errors"

// -----------------METHODS------------------

// Method for Request
type method string

// Constants used to create Methods
const (
	GET           method = "GET"
	POST          method = "POST"
	InvalidMethod method = "INVALID"
)

func newMethod(method string) (method, error) {
	switch method {
	case "GET":
		return GET, nil
	case "POST":
		return POST, nil
	default:
		return InvalidMethod, errors.New("Invalid method or method not supported")
	}
}

// --------------------PROTOCOLS---------------------

// Protocol for Request/Response
type protocol string

// Constants used to create protocols
const (
	HTTP11          protocol = "HTTP/1.1"
	InvalidProtocol protocol = "INVALID"
)

func newProtocol(protocol string) (protocol, error) {
	switch protocol {
	case "HTTP/1.1":
		return HTTP11, nil
	default:
		return InvalidProtocol, errors.New("Invalid protocol or protocol not supported")
	}
}
