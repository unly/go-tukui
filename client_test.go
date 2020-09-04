package tukui

import (
	"net/http"
	"net/http/httptest"
	"testing"
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

func testMethod(t *testing.T, r *http.Request, want string) {
	t.Helper()
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

func String(s string) *string {
	return &s
}
