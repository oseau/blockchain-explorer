package web

import (
	"errors"
	"log"
	"net"
	"net/http"
	"sync"
	"syscall"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	wg := new(sync.WaitGroup)
	wg.Add(1)
	server := NewServer()
	go func() {
		defer wg.Done()
		_ = server.Serve()
	}()

	waitForServer("8080")
	res, err := http.Get("http://localhost:8080/health")
	if err != nil {
		t.Fatalf("could not send GET request: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", res.Status)
	}

	if err := server.Shutdown(); err != nil {
		t.Errorf("unexpected Shutdown err: %v", err)
	}
	wg.Wait()

	if _, err := http.Get("http://localhost:8080/health"); !errors.Is(err, syscall.ECONNREFUSED) {
		t.Fatal("expected server to be terminated after shutdown.")
	}
}

func TestServerServeFail(t *testing.T) {
	if _, err := net.Listen("tcp", ":8080"); err != nil {
		t.Fatal("can not create listener")
	}
	server := NewServer()
	if err := server.Serve(); err.Error() != "listen tcp :8080: bind: address already in use" {
		t.Errorf("got unexpected err: %v", err.Error())
	}
}

func TestServerShutdownBeforeServe(t *testing.T) {
	server := NewServer()
	if err := server.Shutdown(); err != nil {
		t.Errorf("unexpected Shutdown err: %v", err)
	}
}

// steal from https://stackoverflow.com/a/56865986
func waitForServer(port string) {
	backoff := 50 * time.Millisecond

	for i := 0; i < 10; i++ {
		conn, err := net.DialTimeout("tcp", ":"+port, 1*time.Second)
		if err != nil {
			time.Sleep(backoff)
			continue
		}
		err = conn.Close()
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	log.Fatalf("Server on port %s not up after 10 attempts", port)
}
