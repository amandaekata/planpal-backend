package streak

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// RegisterRoutes adds streak routes under the given router (protected).
func RegisterRoutes(r *mux.Router, svc *Service) {
	r.HandleFunc("/streaks/me", getMyStreak(svc)).Methods("GET")
	r.HandleFunc("/streaks/record", recordCompletion(svc)).Methods("POST")
}

func getMyStreak(svc *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{"current": 0, "longest": 0})
	}
}

func recordCompletion(svc *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = mux.Vars(r)
		json.NewEncoder(w).Encode(map[string]string{"message": "record completion not implemented"})
	}
}
