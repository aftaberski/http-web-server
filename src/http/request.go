package http

import "bufio"

// Request not used yet
type Request struct {
	Method   string
	URI      string
	Protocol string
	Headers  []string
	Body     []byte
}

// NewRequest used to make new Request structs
func NewRequest(r bufio.Reader) Request {
	headers := make([]string, 0)
	headers = append(headers, "headers")
	body := make([]byte, 0)
	return Request{"method", "uri", "protocol", headers, body}
}
