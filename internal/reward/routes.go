package reward

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// RegisterRoutes adds reward/XP/leaderboard routes under the given router (protected).
func RegisterRoutes(r *mux.Router, svc *Service) {
	r.HandleFunc("/rewards/xp", getXP(svc)).Methods("GET")
	r.HandleFunc("/rewards/badges", listBadges(svc)).Methods("GET")
	r.HandleFunc("/rewards/leaderboard", getLeaderboard(svc)).Methods("GET")
}

func getXP(svc *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{"xp": 0})
	}
}

func listBadges(svc *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{"badges": []interface{}{}})
	}
}

func getLeaderboard(svc *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = mux.Vars(r)
		json.NewEncoder(w).Encode(map[string]interface{}{"entries": []interface{}{}})
	}
}
