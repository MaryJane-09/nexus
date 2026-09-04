package health

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Data struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Version string `json:"version"`
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	var response = Data{
		Status:  "success",
		Message: "Nexus API is running",
		Version: "v1",
	}
	encoder := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")
	err := encoder.Encode(response)
	if err != nil {
		fmt.Println("Encodeing failed", err)
	}
}
