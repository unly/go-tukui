package tukui

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func setupTestEnv() (client *Client, mux *http.ServeMux, teardown func()) {
	mux = http.NewServeMux()

	api := http.NewServeMux()
	api.Handle("/", mux)

	server := httptest.NewServer(api)

	client = NewClient(nil)
	client.url = server.URL + "/"

	return client, mux, server.Close
}

func testHTTPMethod(t *testing.T, r *http.Request, want string) {
	t.Helper()
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

func testHTTPQuery(t *testing.T, r *http.Request, want url.Values) {
	t.Helper()
	if got := r.URL.Query(); !cmp.Equal(got, want) {
		t.Errorf("Request query: %v, want %v", got, want)
	}
}

func String(s string) *string {
	return &s
}
