package router

import (
	"net/http"

	"github.com/MaryJane-09/nexus/backend/internal/health"
	"github.com/MaryJane-09/nexus/backend/internal/register"
 )

func New() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", health.HealthHandler)
	mux.HandleFunc("/register", register.RegisterHandler)

	return mux
}
