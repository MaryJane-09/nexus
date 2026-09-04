package router

import (
	"net/http"

	"github.com/MaryJane-09/nexus/backend/internal/health"
)

func New() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", health.HealthHandler)

	return mux
}
