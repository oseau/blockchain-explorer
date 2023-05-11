package web

// routes actually mount each handlers
func (s *Server) routes() {
	s.router.HandleFunc("/health", s.handleHealth())
	s.router.HandleFunc("/api/nonce", s.handleNonce())
	s.router.HandleFunc("/api/login", s.handleLogIn())
}
