package http

import (
	"bytes"
	"fmt"
)

// Response holds an HTTP response
type Response struct {
	Protocol      Protocol
	StatusCode    string
	StatusMessage string
	Headers       []string
	Body          []byte
}

// NewResponse creates a Response
func NewResponse(proto string, statusCode string, statusMessage string, headers []string, body []byte) (*Response, error) {
	protocol, newProtocolErr := newProtocol(proto)
	if newProtocolErr != nil {
		return nil, newProtocolErr
	}

	return &Response{protocol, statusCode, statusMessage, headers, body}, nil
}

// Bytes formats a Response to be written to a conn
// Uses bytes package for efficient string concatenation
// http://herman.asia/efficient-string-concatenation-in-go
func (res *Response) Bytes() []byte {

	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("%s %s %s\r\n", res.Protocol, res.StatusCode, res.StatusMessage))

	for _, header := range res.Headers {
		buffer.WriteString(fmt.Sprintf("%s\r\n", header))
	}

	buffer.WriteString("\r\n") // blank line between headers and body

	if len(res.Body) > 0 {
		buffer.Write(res.Body)
	}

	return buffer.Bytes()
}
