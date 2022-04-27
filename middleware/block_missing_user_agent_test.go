package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBlockMissingUserAgentMux(t *testing.T) {
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

	resp := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "http://localhost", nil)
	req.Header.Set("User-Agent", "test")
	blockFloc := BlockMissingUserAgentMux(testHandler)
	blockFloc.ServeHTTP(resp, req)

	if resp.Result().StatusCode != 200 {
		t.Errorf("incorrect status code, got: %d, want: %d.", resp.Result().StatusCode, 200)
	}
}

func TestBlockMissingUserAgentMux_MissingUserAgent(t *testing.T) {
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

	resp := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "http://localhost", nil)
	blockFloc := BlockMissingUserAgentMux(testHandler)
	blockFloc.ServeHTTP(resp, req)

	if resp.Result().StatusCode != 400 {
		t.Errorf("incorrect status code, got: %d, want: %d.", resp.Result().StatusCode, 400)
	}
}
