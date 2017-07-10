package http

import "errors"

// -----------------METHODS------------------

// Method for Request
type Method string

// Constants used to create Methods
const (
	GET           Method = "GET"
	POST          Method = "POST"
	InvalidMethod Method = "INVALID"
)

func newMethod(method string) (Method, error) {
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
type Protocol string

// Constants used to create Protocols
const (
	HTTP11          Protocol = "HTTP/1.1"
	InvalidProtocol Protocol = "INVALID"
)

func newProtocol(protocol string) (Protocol, error) {
	switch protocol {
	case "HTTP/1.1":
		return HTTP11, nil
	default:
		return InvalidProtocol, errors.New("Invalid protocol or protocol not supported")
	}
}
