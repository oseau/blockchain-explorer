package web

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleHealth(t *testing.T) {
	server := NewServer()
	req, err := http.NewRequest("GET", "localhost:8080/health", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}
	rec := httptest.NewRecorder()
	server.handleHealth()(rec, req)
	res := rec.Result()
	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("could not read response: %v", err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", res.Status)
	}
	if string(b) != "" {
		t.Errorf("expected empty content; got %v", string(b))
	}
}
