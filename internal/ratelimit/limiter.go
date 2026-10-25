// Package ratelimit limita solicitudes por dirección IP.
package ratelimit

import (
	"net"
	"net/http"
	"strings"
	"sync"
	"time"
)

type visitor struct {
	started time.Time
	count   int
}
type Limiter struct {
	mu       sync.Mutex
	visitors map[string]visitor
	limit    int
	window   time.Duration
}

func New(limit int, window time.Duration) *Limiter {
	return &Limiter{visitors: make(map[string]visitor), limit: limit, window: window}
}

func (l *Limiter) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := clientIP(r)
		now := time.Now()
		l.mu.Lock()
		current := l.visitors[key]
		if current.started.IsZero() || now.Sub(current.started) >= l.window {
			current = visitor{started: now}
		}
		current.count++
		l.visitors[key] = current
		allowed := current.count <= l.limit
		l.mu.Unlock()
		if !allowed {
			w.Header().Set("Retry-After", "60")
			http.Error(w, "rate limit exceeded", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func clientIP(r *http.Request) string {
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err == nil {
		return host
	}
	forwarded := strings.TrimSpace(strings.Split(r.Header.Get("X-Forwarded-For"), ",")[0])
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}
