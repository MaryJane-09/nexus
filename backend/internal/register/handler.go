package register

import (
	"encoding/json"
	"net/http"
	"github.com/MaryJane-09/nexus/backend/internal/user"
	"github.com/MaryJane-09/nexus/backend/internal/validate"
)


func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	var info = user.User{}

	decoder := json.NewDecoder(r.Body)
	w.Header().Set("Content-Type", "application/json")
	err := decoder.Decode(&info)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}
	err = validate.ValidateRegister(info); 
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(info); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
