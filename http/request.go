package http

import (
	"bufio"
	"errors"
	"strconv"
	"strings"
)

// Request holds an HTTP Request
type Request struct {
	Method   Method
	URI      string
	Protocol Protocol
	Headers  []string
	Body     []byte
}

// // NewRequest used to make new Request structs
// func NewRequest(scanner *bufio.Scanner) Request {
// 	headers := make([]string, 0)
// 	headers = append(headers, "headers")
// 	body := make([]byte, 0)
// 	return Request{"method", "uri", "protocol", headers, body}
// }

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

// Protocol for Request
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

// -----------------REQUEST PARSING------------------

// GetStatusLine parses a status line of an HTTP request
// ex. GET / HTTP/1.1
func GetStatusLine(scanner *bufio.Scanner) (Method, string, Protocol, error) {
	scanner.Scan()
	status := scanner.Text()

	splitStatus := strings.Split(status, " ")
	method, newMethodErr := newMethod(splitStatus[0])
	if newMethodErr != nil {
		return InvalidMethod, "", InvalidProtocol, newMethodErr
	}

	uri := splitStatus[1]

	protocol, newProtocolErr := newProtocol(splitStatus[2])
	if newProtocolErr != nil {
		return InvalidMethod, "", InvalidProtocol, newProtocolErr
	}

	return method, uri, protocol, nil

}

// GetHeaders gets the headers of an HTTP request
func GetHeaders(scanner *bufio.Scanner) []string {
	headers := make([]string, 0)

	for scanner.Scan() {
		header := scanner.Text()

		if len(header) == 0 {
			break
		}

		headers = append(headers, header)
	}
	return headers
}

// GetContentLength checks headers for "Content-Length" key
// when the method is POST
func GetContentLength(headers []string) (int64, error) {
	contentLengthStr := "Content-Length:"
	for _, header := range headers {
		if strings.Contains(header, contentLengthStr) {
			sArr := strings.Split(header, " ")
			contentLength, _ := strconv.ParseInt(sArr[1], 0, 64)
			return contentLength, nil
		}
	}
	return -1, errors.New("No content-length header")
}

// GetBody gets the body of an HTTP request
// Currently just returns empty byte array
func GetBody(scanner *bufio.Scanner, contentLength int64) []byte {
	bytes := make([]byte, contentLength)
	scanner.Split(bufio.ScanBytes)
	var i int64
	for i = 0; i < contentLength; i++ {
		scanner.Scan()
		bytes[i] = scanner.Bytes()[0]
	}
	return bytes
}

// NewRequest instantiates a Request
func NewRequest(scanner *bufio.Scanner) (*Request, error) {

	method, uri, protocol, err := GetStatusLine(scanner)
	if err != nil {
		return nil, err
	}

	headers := GetHeaders(scanner)

	var body []byte

	// If POST request, check for 'Content-Length' header
	// and get post body
	if method == POST {
		contentLength, err := GetContentLength(headers)
		if err != nil {
			return nil, errors.New("Content-Length header required for POST request")
		}
		body = GetBody(scanner, contentLength)
	}

	request := Request{method, uri, protocol, headers, body}
	return &request, nil
}
