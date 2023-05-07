package web

import (
	"errors"
	"net/http"
	"sync"
	"syscall"
	"testing"
)

func TestServer(t *testing.T) {
	wg := new(sync.WaitGroup)
	wg.Add(1)
	server := NewServer()
	go func() {
		defer wg.Done()
		server.Serve()
	}()

	res, err := http.Get("http://localhost:8080/health")
	if err != nil {
		t.Fatalf("could not send GET request: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", res.Status)
	}

	server.Shutdown()
	wg.Wait()

	if _, err := http.Get("http://localhost:8080/health"); !errors.Is(err, syscall.ECONNREFUSED) {
		t.Fatal("expected server to be terminated after shutdown.")
	}
}
