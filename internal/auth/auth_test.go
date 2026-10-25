package auth

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func TestRequireRejectsAnonymousRequest(t *testing.T) {
	handler := New("secret").Require(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) { w.WriteHeader(http.StatusNoContent) }))
	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, httptest.NewRequest(http.MethodGet, "/private", nil))
	if recorder.Code != http.StatusUnauthorized {
		t.Fatalf("status = %d, want %d", recorder.Code, http.StatusUnauthorized)
	}
}

func TestTokenRoundTrip(t *testing.T) {
	service := New("test-secret")
	token, err := service.Token(42)
	if err != nil {
		t.Fatal(err)
	}
	id, err := service.UserID(token)
	if err != nil {
		t.Fatal(err)
	}
	if id != 42 {
		t.Fatalf("id inesperado: %d", id)
	}
}
func TestUserIDRejectsTamperedToken(t *testing.T) {
	service := New("test-secret")
	token, err := service.Token(42)
	if err != nil {
		t.Fatal(err)
	}
	token += "tampered"
	if _, err := service.UserID(token); err == nil {
		t.Fatal("se esperaba rechazar token alterado")
	}
}
func TestMiddlewareSetsIdentityHeader(t *testing.T) {
	service := New("test-secret")
	token, err := service.Token(7)
	if err != nil {
		t.Fatal(err)
	}
	handler := service.Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got := r.Header.Get("X-User-ID"); got != "7" {
			t.Errorf("identity header: %q", got)
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	request.Header.Set("Authorization", "Bearer "+token)
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, request)
	if response.Code != http.StatusNoContent {
		t.Fatalf("status: %d", response.Code)
	}
}
func TestUserIDRejectsExpiredToken(t *testing.T) {
	service := New("test-secret")
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"sub": "42", "exp": time.Now().Add(-time.Hour).Unix()}).SignedString([]byte("test-secret"))
	if err != nil {
		t.Fatal(err)
	}
	if _, err := service.UserID(token); err == nil || !strings.Contains(err.Error(), "token") {
		t.Fatalf("se esperaba error de expiración: %v", err)
	}
}
