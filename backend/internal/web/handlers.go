package web

import (
	"encoding/json"
	"io/ioutil"
	"math/big"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/gorilla/sessions"
)

var (
	// TODO: move to secret
	sessionStore = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
	COOKIE_NAME  = "BlockchainExplorer"
	nonces       = sync.Map{}
)

func (s *Server) handleHealth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

type RespNonce struct {
	Nonce string `json:"nonce"`
}

type Login struct {
	Signature string `json:"signature"`
}
type RespLogin struct {
	ValidUntil int64 `json:"validUntil"`
}

func (s *Server) handleNonce() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		address := r.URL.Query().Get("address")
		lang := r.URL.Query().Get("lang")
		nonce := getNonce(lang)
		nonces.Store(address, nonce)
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(RespNonce{nonce})
	}
}

func (s *Server) handleLogIn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		address := r.URL.Query().Get("address")
		b, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var msg Login
		if err = json.Unmarshal(b, &msg); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		nonce, ok := nonces.Load(address)
		if !ok {
			return
		}
		if !verify(address, nonce.(string), msg.Signature) {
			return
		}
		nonces.Delete(address)
		session, _ := sessionStore.Get(r, COOKIE_NAME)
		tomorrow := time.Now().Add(24 * time.Hour).Unix()
		session.Values["validUntil"] = tomorrow
		if err := session.Save(r, w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(RespLogin{tomorrow})
	}
}

func (s *Server) handleLogOut() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := sessionStore.Get(r, COOKIE_NAME)
		session.Values["validUntil"] = -1
		if err := session.Save(r, w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) handleWs() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		serveWs(s.hub, w, r)
	}
}

type Balance struct {
	BlockNumber string `json:"blockNumber"`
	Balance     string `json:"balance"`
}
type MsgNewBalances struct {
	Action   string    `json:"action"`
	Balances []Balance `json:"balances"`
}

func (s *Server) handleTest() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		blockNumber := new(big.Int)
		balance := new(big.Int)
		blockNumber.SetString(r.URL.Query().Get("blockNumber"), 10)
		balance.SetString(r.URL.Query().Get("balance"), 10)
		b, _ := json.Marshal(MsgNewBalances{
			Action:   "new-balances",
			Balances: []Balance{Balance{BlockNumber: blockNumber.String(), Balance: balance.String()}}})
		s.hub.broadcast <- b
	}
}

type RespGetBalances struct {
	Balance  string    `json:"balance"`
	Balances []Balance `json:"balances"`
}

func (s *Server) handleGetBalances() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		address := r.URL.Query().Get("address")
		balances, err := s.data.GetBalances(address)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		resp := RespGetBalances{Balances: toBalances(balances)}
		if len(balances) == 0 {
			if balance, err := s.data.GetBalanceRpc(address); err == nil {
				resp.Balance = balance.String()
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(resp)
		s.taskBalance <- address
	}
}
