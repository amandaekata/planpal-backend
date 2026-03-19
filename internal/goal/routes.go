package goal

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// RegisterRoutes adds goal routes under the given router (protected).
func RegisterRoutes(r *mux.Router, svc *Service) {
	r.HandleFunc("/goals", listGoals(svc)).Methods("GET")
	r.HandleFunc("/goals", createGoal(svc)).Methods("POST")
	r.HandleFunc("/goals/{id}", getGoal(svc)).Methods("GET")
	r.HandleFunc("/goals/{id}", updateGoal(svc)).Methods("PUT", "PATCH")
	r.HandleFunc("/goals/{id}", deleteGoal(svc)).Methods("DELETE")
}

func listGoals(svc *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{"goals": []interface{}{}})
	}
}

func createGoal(svc *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "create goal not implemented"})
	}
}

func getGoal(svc *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = mux.Vars(r)
		json.NewEncoder(w).Encode(map[string]string{"message": "get goal not implemented"})
	}
}

func updateGoal(svc *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "update goal not implemented"})
	}
}

func deleteGoal(svc *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "delete goal not implemented"})
	}
}
