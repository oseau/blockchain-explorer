package web

// routes actually mount each handlers
func (s *Server) routes() {
	s.router.HandleFunc("/health", s.handleHealth())
}
