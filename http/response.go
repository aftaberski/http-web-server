package http

import "fmt"

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

func (res *Response) Format() []byte {
	// Write status line ex. HTTP/1.1 500 Internal Server Error
	formattedRes := []byte(fmt.Sprintf("%s %s %s\r\n", res.Protocol, res.StatusCode, res.StatusMessage))

	// Write all headers
	for _, header := range res.Headers {
		formattedRes = append(formattedRes, fmt.Sprintf("%s\r\n", header)...)
	}

	// Blank line separating headers from body as specified for HTTP/1.1
	formattedRes = append(formattedRes, "\r\n"...)

	if len(res.Body) > 0 {
		formattedRes = append(formattedRes, res.Body...)
	}
	fmt.Println(formattedRes)
	return formattedRes
}
