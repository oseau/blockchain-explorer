package web

// routes actually mount each handlers
func (s *Server) routes() {
	s.router.HandleFunc("/health", s.handleHealth())
	s.router.HandleFunc("/api/nonce", s.handleNonce())
	s.router.HandleFunc("/api/login", s.handleLogIn())
	s.router.HandleFunc("/api/logout", s.handleLogOut())
	s.router.HandleFunc("/api/get-balances", s.handleGetBalances())
	s.router.HandleFunc("/ws/", s.handleWs())
	s.router.HandleFunc("/ws/test", s.handleTest())
}
