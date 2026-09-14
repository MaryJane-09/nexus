package register

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/MaryJane-09/nexus/backend/internal/otp"
	"github.com/MaryJane-09/nexus/backend/internal/pending"
	"github.com/MaryJane-09/nexus/backend/internal/user"
)

type OTPVerification struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

func VerifyOTPHandler(repo *user.UserRepository, otpRepo *otp.OTPRepository, pendingRepo *pending.PendingRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(ErrorResponse{Error: "Method not allowed"})
			return
		}
		defer r.Body.Close()

		var info OTPVerification
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&info)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrorResponse{Error: "Failed to decode request body"})
			return
		}
		storedOTP, err := otpRepo.FindByEmail(info.Email)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(ErrorResponse{Error: "No OTP was found for this email."})
			return
		}
		if time.Now().After(storedOTP.ExpiresAt)  {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrorResponse{Error: "This OTP has expired"})
			otpRepo.Delete(info.Email)
			return
		}
		if info.Code != storedOTP.Code{
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrorResponse{Error: "Incorrect OTP. No match found"})
			return
		}
	}

}
