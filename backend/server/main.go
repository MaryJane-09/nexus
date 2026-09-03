package main

import (
	"fmt"
	"net/http"
	"github.com/MaryJane-09/nexus/backend/internal/health"
)



func main() {
	fmt.Println("Nexus server running on :8080")
	http.HandleFunc("/health", health.HealthHandler)
	err := http.ListenAndServe(":8080", nil); if err != nil {
		fmt.Println("Failed to start HTTP server: ", err)
	}
}