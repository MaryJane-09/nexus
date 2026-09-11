package register

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/MaryJane-09/nexus/backend/internal/email"
	"github.com/MaryJane-09/nexus/backend/internal/otp"
	"github.com/MaryJane-09/nexus/backend/internal/user"
	"github.com/MaryJane-09/nexus/backend/internal/validate"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func RegisterHandler(repo *user.Repository, otpRepo *otp.Repository, sender *email.EmailSender) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(ErrorResponse{Error: "Method not allowed"})
			return
		}
		defer r.Body.Close()

		var info user.User
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&info)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrorResponse{Error: "Failed to decode request body"})
			return
		}

		err = validate.ValidateRegister(info)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
			return
		}

		code, err := otp.Generate(8)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Could not generate OTP"})
			return
		}

		newOTP := otp.OTP{
			Email:     info.Email,
			Code:      code,
			ExpiresAt: time.Now().Add(otp.ExpiryTime),
			Verified:  false,
		}

		err = otpRepo.Create(newOTP)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Could not create new OTP"})
			return
		}

		err = sender.SendVerification(info.Email, code)
		if err != nil {
			otpRepo.Delete(info.Email)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Verification code generated successfully",
		})
	}
}
