package register

import (
	"encoding/json"
	"net/http"

	"github.com/MaryJane-09/nexus/backend/internal/user"
)

type EmailReg struct {
	Email string `json:"email"`
}

func EmailRegHandler(repo *user.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		defer r.Body.Close()

		var email EmailReg
		decoder := json.NewDecoder(r.Body)
		w.Header().Set("Content-Type", "application/json")
		err := decoder.Decode(&email)
		if err != nil {
			http.Error(w, "Invalid Request", http.StatusBadRequest)
			return
		}

		_, err = repo.FindByEmail(email.Email)
		if err == nil {
			http.Error(w, "Email already exsist", http.StatusConflict)
			return 
		}
		if err != nil {
			//continue
		}
	}
}
