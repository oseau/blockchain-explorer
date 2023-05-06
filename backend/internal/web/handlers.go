package web

import (
	"fmt"
	"net/http"
)

func (s *Server) handleAbout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "about what?")
	}
}
