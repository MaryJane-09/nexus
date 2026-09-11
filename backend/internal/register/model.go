package register

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/MaryJane-09/nexus/backend/internal/email"
	"github.com/MaryJane-09/nexus/backend/internal/otp"
	"github.com/MaryJane-09/nexus/backend/internal/user"
)

type EmailReg struct {
	Email string `json:"email"`
}
var ErrNotFound = errors.New("OTP was not found for this email")

func EmailRegHandler(userRepo *user.Repository, otpRepo *otp.Repository, sender email.EmailSender) http.HandlerFunc {
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
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err = userRepo.FindByEmail(email.Email)
		if err == nil {
			http.Error(w, "Email already exsist", http.StatusConflict)
			return
		}

		if !errors.Is(err, ErrNotFound) {
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
			Code:      code,
			ExpiresAt: time.Now().Add(otp.ExpiryTime),
			Verified:  false,
		}
		err = otpRepo.Create(newOTP)
		if err != nil {
			http.Error(w, "Could not create new OTP", http.StatusInternalServerError)
			return
		}

		err = sender.SendVerification(email.Email, code)
		if err != nil{
			otpRepo.Delete(email.Email)
			http.Error(w, "Email fail to send", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Verification code generated successfully",
		})
	}
}
