package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"noniserver/m/v2/internal/auth"
	"noniserver/m/v2/internal/chat"
	"noniserver/m/v2/internal/coach"
	"noniserver/m/v2/internal/community"
	"noniserver/m/v2/internal/database"
	"noniserver/m/v2/internal/ratelimit"
)

type response struct {
	Status int `json:"status"`
	Data   any `json:"data,omitempty"`
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	if err := godotenv.Load(); err != nil && !errors.Is(err, os.ErrNotExist) {
		log.Printf("cargar .env: %v", err)
	}
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL no está configurada; copia .env.example como .env")
	}
	db, err := database.Open(ctx, dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if err = database.Migrate(ctx, db); err != nil {
		log.Fatal(err)
	}

	store := community.NewStore(db)
	hub := chat.NewHub(store.SaveEvent, store.AuthorizeSocket)
	coachService := coach.NewRulesCoach()
	authService := auth.New(env("AUTH_SECRET", "development-only-change-me"))
	router := chi.NewRouter()
	router.Use(ratelimit.New(120, time.Minute).Middleware)
	router.Use(authService.Middleware)
	router.Post("/auth/login", func(w http.ResponseWriter, r *http.Request) {
		var request struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			http.Error(w, "JSON inválido", http.StatusBadRequest)
			return
		}
		userID, err := store.Authenticate(r.Context(), request.Username, request.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		token, err := authService.Token(userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.SetCookie(w, &http.Cookie{Name: "checkpoint_token", Value: token, HttpOnly: true, SameSite: http.SameSiteLaxMode, Path: "/", MaxAge: 86400})
		write(w, http.StatusOK, map[string]string{"token": token})
	})
	router.Get("/health", func(w http.ResponseWriter, _ *http.Request) {
		write(w, http.StatusOK, map[string]string{"status": "ok"})
	})
	frontend := http.FileServer(http.Dir("web"))
	router.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))
	router.Get("/", func(w http.ResponseWriter, r *http.Request) { frontend.ServeHTTP(w, r) })
	router.Get("/login", func(w http.ResponseWriter, r *http.Request) { servePage(w, r, "login.html") })
	router.Get("/register", func(w http.ResponseWriter, r *http.Request) { servePage(w, r, "register.html") })
	router.Get("/app", func(w http.ResponseWriter, r *http.Request) { servePage(w, r, "app.html") })
	router.Get("/chat/{roomID}", func(w http.ResponseWriter, r *http.Request) { servePage(w, r, "chat.html") })
	router.Get("/ui/lesson", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = w.Write([]byte(`<p><strong>HTMX</strong> pidió este fragmento al servidor y lo insertó sin recargar la página. El navegador conserva la pantalla y solo cambia esta pieza.</p>`))
	})
	router.Group(func(protected chi.Router) {
		protected.Use(authService.Require)
		protected.Post("/api/coach/next", func(w http.ResponseWriter, r *http.Request) {
			var request struct {
				Goal     coach.Goal     `json:"goal"`
				Progress coach.Progress `json:"progress"`
			}
			if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
				http.Error(w, "JSON inválido", http.StatusBadRequest)
				return
			}
			mission, err := coachService.NextMission(r.Context(), request.Goal, request.Progress)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			write(w, http.StatusOK, mission)
		})
		protected.Post("/api/check-ins", func(w http.ResponseWriter, r *http.Request) {
			var request struct {
				Goal    coach.Goal `json:"goal"`
				Content string     `json:"content"`
			}
			if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
				http.Error(w, "JSON inválido", http.StatusBadRequest)
				return
			}
			feedback, err := coachService.ReviewCheckIn(r.Context(), request.Goal, request.Content)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			write(w, http.StatusOK, feedback)
		})
	})
	router.Mount("/", community.Routes(store, authService.Require))
	router.Handle("/ws/rooms/{roomID}", authService.Require(hub))

	server := &http.Server{Addr: env("PORT", ":3000"), Handler: router, ReadHeaderTimeout: 5 * time.Second, ReadTimeout: 10 * time.Second, WriteTimeout: 10 * time.Second, IdleTimeout: 60 * time.Second}
	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("servidor: %v", err)
		}
	}()
	log.Printf("API escuchando en %s", server.Addr)
	<-ctx.Done()
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Printf("cerrar API: %v", err)
	}
}

func env(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
func write(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(response{Status: status, Data: data})
}

func servePage(w http.ResponseWriter, r *http.Request, name string) {
	http.ServeFile(w, r, "web/"+name)
}
