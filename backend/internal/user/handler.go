package user

import (
	"encoding/json"
	"net/http"
)

func UsersHandler(repo *Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		allUsers := repo.GetAllUsers()

		encoder := json.NewEncoder(w)
		w.Header().Set("Content-Type", "application/json")
		err := encoder.Encode(allUsers)
		if err != nil {
			http.Error(w, "Encoding failed", http.StatusInternalServerError)
			return
		}

	}
}
