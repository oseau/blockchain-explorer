package web

// routes actually mount each handlers
func (s *Server) routes() {
	s.router.HandleFunc("/about", s.handleAbout())
}
