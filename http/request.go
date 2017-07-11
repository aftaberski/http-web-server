package http

import (
	"bufio"
	"errors"
	"net"
	"strconv"
	"strings"
)

// getStatusLine parses a status line of an HTTP request
// ex. GET / HTTP/1.1
func getStatusLine(scanner *bufio.Scanner) (method, string, protocol, error) {
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

func getHeaders(scanner *bufio.Scanner) []string {
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

// getContentLength checks headers for "Content-Length" key
// when the method is POST
func getContentLength(headers []string) (int64, error) {
	for _, header := range headers {
		if strings.Contains(header, "Content-Length:") {
			stringArr := strings.Split(header, " ")
			contentLength, err := strconv.ParseInt(stringArr[1], 0, 64)
			if err != nil {
				return -1, errors.New("Could not parse Content-Length header into int")
			}
			return contentLength, nil
		}
	}
	return -1, errors.New("No content-length header")
}

// GetBody gets the body of an HTTP request
// Currently just returns empty byte array
func getBody(scanner *bufio.Scanner, contentLength int64) []byte {
	bytes := make([]byte, contentLength)
	scanner.Split(bufio.ScanBytes)
	var i int64
	for i = 0; i < contentLength; i++ {
		scanner.Scan()
		bytes[i] = scanner.Bytes()[0]
	}
	return bytes
}

// Request holds an HTTP Request
type Request struct {
	Method   method
	URI      string
	Protocol protocol
	Headers  []string
	Body     []byte
}

// NewRequest instantiates a Request
func NewRequest(conn net.Conn) (*Request, error) {
	scanner := bufio.NewScanner(conn)
	method, uri, protocol, err := getStatusLine(scanner)
	if err != nil {
		return nil, err
	}

	headers := getHeaders(scanner)

	var body []byte

	// If POST request, check for 'Content-Length' header
	// and get post body
	if method == POST {
		contentLength, err := getContentLength(headers)
		if err != nil {
			return nil, errors.New("Content-Length header required for POST request")
		}
		body = getBody(scanner, contentLength)
	}

	request := Request{method, uri, protocol, headers, body}
	return &request, nil
}
