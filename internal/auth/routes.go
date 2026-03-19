package auth

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// RegisterRoutes adds auth routes to the given router (e.g. /api/v1).
func RegisterRoutes(r *mux.Router, svc *Service) {
	r.HandleFunc("/auth/register", handleRegister(svc)).Methods("POST")
	r.HandleFunc("/auth/login", handleLogin(svc)).Methods("POST")
	r.HandleFunc("/auth/refresh", handleRefresh(svc)).Methods("POST")
}

func handleRegister(svc *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		// TODO: parse body, create user, issue tokens
		json.NewEncoder(w).Encode(map[string]string{"message": "register not implemented"})
	}
}

func handleLogin(svc *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		// TODO: parse body, validate user, issue tokens
		json.NewEncoder(w).Encode(map[string]string{"message": "login not implemented"})
	}
}

func handleRefresh(svc *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		// TODO: parse refresh token, validate, issue new pair
		json.NewEncoder(w).Encode(map[string]string{"message": "refresh not implemented"})
	}
}
