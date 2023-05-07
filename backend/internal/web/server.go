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
func (s *Server) Serve() {
	s.srv = &http.Server{Addr: ":8080", Handler: s.router}
	go func() {
		log.Println("Server start!")
		if err := s.srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("ListenAndServe error: %v", err)
		} else {
			log.Println("Server stop!")
		}
	}()
}

func (s *Server) Shutdown() {
	if s.srv == nil {
		log.Println("server not started yet.")
		return
	}
	// force shut down after 5s
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.srv.Shutdown(ctx); err != nil {
		log.Println("shutdown server failed, err: ", err.Error())
	}
}
