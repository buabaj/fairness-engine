package server

import (
	"fmt"
	"net/http"
)

type Server struct {
	Port int
}

func NewServer(port int) *Server {
	return &Server{
		Port: port,
	}
}

// Run starts the server and listens for incoming connections.
func (s *Server) Run() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/key-generation", handleKeyGeneration)
	mux.HandleFunc("/share-submission", handleShareSubmission)
	mux.HandleFunc("/computation", handleComputation)

	addr := fmt.Sprintf(":%d", s.Port)
	fmt.Printf("Server listening on %s\n", addr)
	return http.ListenAndServe(addr, mux)
}
