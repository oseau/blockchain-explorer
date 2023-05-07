package web

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouting(t *testing.T) {
	srv := httptest.NewServer(NewServer().router)
	defer srv.Close()

	res, err := http.Get(fmt.Sprintf("%s/health", srv.URL))
	if err != nil {
		t.Fatalf("could not send GET request: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", res.Status)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("could not read response: %v", err)
	}
	if string(b) != "" {
		t.Errorf("expected empty content; got %v", string(b))
	}
}
