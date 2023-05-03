package server

type Server struct {
	Logger *ServerLogger
}

func NewServer() (*Server, error) {
	// Creating Server example
	server := &Server{}
	// Initializing logger
	server.Logger = initServerLogger()
	return server, nil
}
