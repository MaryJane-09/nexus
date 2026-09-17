package router

import (
	"net/http"

	"github.com/MaryJane-09/nexus/backend/config"
	"github.com/MaryJane-09/nexus/backend/internal/email"
	"github.com/MaryJane-09/nexus/backend/internal/health"
	"github.com/MaryJane-09/nexus/backend/internal/otp"
	"github.com/MaryJane-09/nexus/backend/internal/pending"
	"github.com/MaryJane-09/nexus/backend/internal/register"
	"github.com/MaryJane-09/nexus/backend/internal/user"
	"github.com/jackc/pgx/v5/pgxpool"
)


func New(pool *pgxpool.Pool) *http.ServeMux {

	sender := email.EmailSender{
		GmailAddress:     config.AppConfig.GmailAddress,
		GmailAppPassword: config.AppConfig.GmailAppPassword,
		SMTPHost:         "smtp.gmail.com",
		SMTPPort:         "587",
	}
	repo := user.NewRepository(pool)
	otpRepo := otp.NewRepository()
	pendingRepo := pending.NewRepository()
	mux := http.NewServeMux()
	mux.HandleFunc("/health", health.HealthHandler)
	mux.HandleFunc("/register", register.RegisterHandler(repo, otpRepo, pendingRepo, &sender))
	mux.HandleFunc("/users", user.UsersHandler(repo))
	mux.HandleFunc("/verify-otp", register.VerifyOTPHandler(repo, otpRepo, pendingRepo))
	return mux
}
