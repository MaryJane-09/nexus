package register

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Register struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if  r.Method != http.MethodPost{
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	var info = Register{}

	decoder := json.NewDecoder(r.Body)
	w.Header().Set("Content-Type", "application/json")
	err := decoder.Decode(&info)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) 
	
	if err := json.NewEncoder(w).Encode(info); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
