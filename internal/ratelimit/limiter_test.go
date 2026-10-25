package ratelimit

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestMiddlewareLimitsRequests(t *testing.T) {
	limiter := New(1, time.Minute)
	handler := limiter.Middleware(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) { w.WriteHeader(http.StatusNoContent) }))
	first := httptest.NewRequest(http.MethodGet, "/", nil)
	first.RemoteAddr = "192.0.2.1:1234"
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, first)
	if response.Code != http.StatusNoContent {
		t.Fatalf("primer request: %d", response.Code)
	}
	second := httptest.NewRequest(http.MethodGet, "/", nil)
	second.RemoteAddr = first.RemoteAddr
	response = httptest.NewRecorder()
	handler.ServeHTTP(response, second)
	if response.Code != http.StatusTooManyRequests {
		t.Fatalf("segunda request: %d", response.Code)
	}
}
