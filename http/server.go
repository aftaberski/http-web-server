package http

import (
	"fmt"
	"log"
	"net"
)

// Server instance
type Server struct {
	AvailableRoutes map[Method]map[string]Handler
}

// Handler used to create handlers
type Handler func(Request) (Response, error)

// NewServer instantiates a new server
func NewServer() *Server {
	server := Server{}
	server.AvailableRoutes = make(map[Method]map[string]Handler)
	return &server
}

// AddRoute adds a route to a server
func (server *Server) AddRoute(method Method, uri string, handler Handler) {
	if _, ok := server.AvailableRoutes[method]; !ok {
		server.AvailableRoutes[method] = make(map[string]Handler)
	}
	server.AvailableRoutes[method][uri] = handler
}

func (server *Server) handleConnection(conn net.Conn) {
	defer conn.Close()

	req, err := NewRequest(conn)
	if err != nil {
		log.Fatal(err)
	}
	if _, ok := server.AvailableRoutes[req.Method][req.URI]; !ok {
		// TODO return default 404 response
		// handleErrors
		log.Fatal("Route not registered: %s", req.URI)
	}
	handler := server.AvailableRoutes[req.Method][req.URI]
	response, err := handler(*req)
	// TODO: Create writable response and write to conn
}

// Run listens on the port its given
// and begins accepting requests
func (server *Server) Run(port int64) {
	portStr := fmt.Sprintf(":%d", port)
	l, err := net.Listen("tcp", portStr)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	fmt.Printf("Listening at %d...", port)
	for {
		// Wait for a connection.
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		server.handleConnection(conn)
	}
}
