package user

import (
	"encoding/json"
	"net/http"

	"github.com/amandaekata/planpal-backend/middleware"
	"github.com/gorilla/mux"
)

// RegisterRoutes adds user routes under the given router (protected).
func RegisterRoutes(r *mux.Router, svc *Service) {
	r.HandleFunc("/users/me", getMe(svc)).Methods("GET")
	r.HandleFunc("/users/me", updateMe(svc)).Methods("PUT", "PATCH")
}

func getMe(svc *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		userID := middleware.UserID(r.Context())
		if userID == "" {
			http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
			return
		}
		u, _ := svc.GetByID(r.Context(), userID)
		if u == nil {
			json.NewEncoder(w).Encode(map[string]string{"id": userID})
			return
		}
		json.NewEncoder(w).Encode(u)
	}
}

func updateMe(svc *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = middleware.UserID(r.Context())
		json.NewEncoder(w).Encode(map[string]string{"message": "update me not implemented"})
	}
}
