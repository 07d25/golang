// Package auth emite y valida tokens JWT para la API.
package auth

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const userIDKey contextKey = "user_id"

type Service struct{ secret []byte }

func New(secret string) *Service { return &Service{secret: []byte(secret)} }
func (s *Service) Token(userID int64) (string, error) {
	if len(s.secret) == 0 {
		return "", errors.New("AUTH_SECRET no configurado")
	}
	claims := jwt.MapClaims{"sub": strconv.FormatInt(userID, 10), "exp": time.Now().Add(24 * time.Hour).Unix()}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(s.secret)
}
func (s *Service) UserID(token string) (int64, error) {
	parsed, err := jwt.Parse(token, func(t *jwt.Token) (any, error) {
		if t.Method != jwt.SigningMethodHS256 {
			return nil, errors.New("algoritmo JWT inválido")
		}
		return s.secret, nil
	})
	if err != nil || !parsed.Valid {
		return 0, errors.New("token inválido")
	}
	value, err := parsed.Claims.GetSubject()
	if err != nil {
		return 0, err
	}
	return strconv.ParseInt(value, 10, 64)
}
func (s *Service) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := strings.TrimSpace(r.Header.Get("Authorization"))
		if strings.HasPrefix(header, "Bearer ") {
			id, err := s.UserID(strings.TrimSpace(strings.TrimPrefix(header, "Bearer ")))
			if err == nil {
				r = r.WithContext(context.WithValue(r.Context(), userIDKey, id))
				r.Header.Set("X-User-ID", strconv.FormatInt(id, 10))
			}
		}
		next.ServeHTTP(w, r)
	})
}

// Require rechaza las peticiones que no presenten un JWT válido.
func (s *Service) Require(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, ok := UserID(r.Context()); !ok {
			writeJSONError(w, http.StatusUnauthorized, "autenticación requerida")
			return
		}
		next.ServeHTTP(w, r)
	})
}

func writeJSONError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(map[string]any{"status": status, "error": message})
}
func UserID(ctx context.Context) (int64, bool) {
	value, ok := ctx.Value(userIDKey).(int64)
	return value, ok
}
