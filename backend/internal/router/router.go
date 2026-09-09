package router

import (
	"net/http"

	"github.com/MaryJane-09/nexus/backend/internal/health"
	"github.com/MaryJane-09/nexus/backend/internal/register"
	"github.com/MaryJane-09/nexus/backend/internal/user"

 )

func New() *http.ServeMux {
	repo := user.NewRepository()
	mux := http.NewServeMux()
	mux.HandleFunc("/health", health.HealthHandler)
	mux.HandleFunc("/register", register.RegisterHandler(repo))
	mux.HandleFunc("/users", user.UsersHandler(repo))

	return mux
}
