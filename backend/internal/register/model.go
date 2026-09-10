package register

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"
	"github.com/MaryJane-09/nexus/backend/internal/otp"
	"github.com/MaryJane-09/nexus/backend/internal/user"
)

type EmailReg struct {
	Email string `json:"email"`
}

func EmailRegHandler(userRepo *user.Repository, otpRepo *otp.Repository) http.HandlerFunc {
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

		_, err = userRepo.FindByEmail(email.Email)
		if err == nil {
			http.Error(w, "Email already exsist", http.StatusConflict)
			return
		}

		if !errors.Is(err, otp.ErrNotFound) {
			http.Error(w, `{"error": "Database lookup failed"}`, http.StatusInternalServerError)
			return
		}

		code, err := otp.Generate(8)
		if err != nil {
			http.Error(w, "Could not generate OTP", http.StatusInternalServerError)
			return
		}
		newOTP := otp.OTP{
			Email:     email.Email,
			Code:   code,
			ExpiresAt: time.Now().Add(otp.ExpiryTime),
			Verified:  false,
		}
		otpRepo.Create(newOTP)

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Verification code generated successfully",
		})
	}
}
