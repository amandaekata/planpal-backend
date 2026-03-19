package notification

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// RegisterRoutes adds notification routes under the given router (protected).
func RegisterRoutes(r *mux.Router, svc *Service) {
	r.HandleFunc("/notifications", listNotifications(svc)).Methods("GET")
	r.HandleFunc("/notifications/{id}/read", markRead(svc)).Methods("POST", "PATCH")
}

func listNotifications(svc *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{"notifications": []interface{}{}})
	}
}

func markRead(svc *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = mux.Vars(r)
		json.NewEncoder(w).Encode(map[string]string{"message": "mark read not implemented"})
	}
}
