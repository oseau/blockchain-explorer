package web

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/oseau/blockchain-explorer/internal/data"
)

type Server struct {
	srv         *http.Server
	router      *http.ServeMux
	hub         *Hub
	data        *data.Server
	taskBalance chan string
}

// NewServer create a Server instance
func NewServer() *Server {
	s := &Server{
		router:      http.NewServeMux(),
		taskBalance: make(chan string, 64),
	}
	s.routes()
	return s
}

// Serve actually start the http server
func (s *Server) Serve() error {
	s.srv = &http.Server{Addr: ":8080", Handler: s.router}
	s.hub = newHub()
	s.data = data.NewServer()
	go s.hub.run()
	go s.runTask()
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
	s.data.Shutdown()
	return s.srv.Shutdown(ctx)
}

func (s *Server) runTask() {
	for account := range s.taskBalance {
		blockNumbers, err := s.data.GetRecentBalanceChangeBlockNumbersRpc(account)
		if err != nil {
			continue
		}
		for _, block := range blockNumbers {
			balance, err := s.data.GetBalanceAtBlockRpc(account, block)
			if err != nil {
				continue
			}
			_ = s.data.UpsertBalance(account, block, balance)
			newBalance, _ := json.Marshal(MsgNewBalances{
				Action:   "new-balances",
				Balances: []Balance{Balance{BlockNumber: block.String(), Balance: balance.String()}}})
			time.Sleep(1 * time.Second)
			s.hub.broadcast <- newBalance
		}
	}
}
