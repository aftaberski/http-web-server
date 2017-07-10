package http

// Response holds an HTTP response
type Response struct {
	Protocol Protocol
	Status   string
	Headers  []string
	Body     []byte
}

// NewResponse creates a Response
func NewResponse(proto string, status string, headers []string, body []byte) (*Response, error) {
	protocol, newProtocolErr := newProtocol(proto)
	if newProtocolErr != nil {
		return nil, newProtocolErr
	}

	return &Response{protocol, status, headers, body}, nil
}
