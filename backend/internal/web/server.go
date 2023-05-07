package web

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"
)

type Server struct {
	srv    *http.Server
	router *http.ServeMux
}

// NewServer create a Server instance
func NewServer() *Server {
	s := &Server{
		router: http.NewServeMux(),
	}
	s.routes()
	return s
}

// Serve actually start the http server
func (s *Server) Serve() error {
	s.srv = &http.Server{Addr: ":8080", Handler: s.router}
	log.Println("Server start!")
	if err := s.srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		log.Printf("ListenAndServe failed; error: %v\n", err)
		return err
	} else {
		log.Println("Server stop!")
	}
	return nil
}

func (s *Server) Shutdown() error {
	if s.srv == nil {
		log.Println("server not started yet.")
		return nil
	}
	// force shut down after 5s
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return s.srv.Shutdown(ctx)
}
