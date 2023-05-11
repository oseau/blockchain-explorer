package web

import (
	"encoding/json"
	"io/ioutil"
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

type NonceResp struct {
	Nonce string `json:"nonce"`
}

type Login struct {
	Signature string `json:"signature"`
}
type LoginResp struct {
	ValidUntil int64 `json:"validUntil"`
}

func (s *Server) handleNonce() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		address := r.URL.Query().Get("address")
		lang := r.URL.Query().Get("lang")
		nonce := getNonce(lang)
		nonces.Store(address, nonce)
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(NonceResp{nonce})
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
		_ = json.NewEncoder(w).Encode(LoginResp{tomorrow})
	}
}
