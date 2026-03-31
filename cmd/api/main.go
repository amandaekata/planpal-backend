package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/amandaekata/planpal-backend/config"
	"github.com/amandaekata/planpal-backend/internal/auth"
	"github.com/amandaekata/planpal-backend/internal/goal"
	"github.com/amandaekata/planpal-backend/internal/notification"
	"github.com/amandaekata/planpal-backend/internal/reward"
	"github.com/amandaekata/planpal-backend/internal/streak"
	"github.com/amandaekata/planpal-backend/internal/user"
	"github.com/amandaekata/planpal-backend/internal/ws"
	"github.com/amandaekata/planpal-backend/middleware"
	"github.com/gorilla/mux"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/rs/cors"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config load: %v", err)
	}

	db, err := sql.Open("pgx", cfg.Database.URL)
	if err != nil {
		log.Fatalf("db open: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("db ping: %v", err)
	}
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)

	// Hub for WebSocket connections (streaks, notifications)
	hub := ws.NewHub()
	go hub.Run()

	// Services (repos can be wired here when DB layer is ready)
	userSvc := user.NewService(nil)
	authSvc := auth.NewService(cfg.JWT.Secret, cfg.JWT.AccessExpiry, cfg.JWT.RefreshExpiry, userSvc)
	goalSvc := goal.NewService(nil)
	streakSvc := streak.NewService(nil)
	rewardSvc := reward.NewService(nil)
	notifSvc := notification.NewService(nil, hub)

	// Router
	r := mux.NewRouter()
	r.Use(middleware.RequestID, middleware.Logging)

	// Health
	r.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"ok"}`))
	}).Methods(http.MethodGet)

	// WebSocket
	r.HandleFunc("/ws", ws.Handler(hub)).Methods(http.MethodGet)

	// Public API
	api := r.PathPrefix("/api/v1").Subrouter()

	// Auth
	auth.RegisterRoutes(api, authSvc)

	// Protected API (require JWT)
	protected := api.PathPrefix("").Subrouter()
	protected.Use(middleware.Auth(cfg.JWT.Secret))
	user.RegisterRoutes(protected, userSvc)
	goal.RegisterRoutes(protected, goalSvc)
	streak.RegisterRoutes(protected, streakSvc)
	reward.RegisterRoutes(protected, rewardSvc)
	notification.RegisterRoutes(protected, notifSvc)

	// CORS
	handler := cors.New(cors.Options{
		AllowedOrigins:   cfg.CORS.Origins,
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodOptions},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	}).Handler(r)

	addr := fmt.Sprintf(":%d", cfg.Port)
	srv := &http.Server{Addr: addr, Handler: handler, ReadHeaderTimeout: 10 * time.Second}

	go func() {
		log.Printf("planpal-api listening on %s (env=%s)", addr, cfg.Env)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("server error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("shutdown: %v", err)
	}
	log.Println("server stopped")
}
