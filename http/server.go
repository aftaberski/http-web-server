package http

type server struct {
	availableRoutes map[string]map[string]handler
}

type handler func(Request) (Response, error)
