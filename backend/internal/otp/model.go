package otp

import (
	"time"
)

type OTP struct {
	Email     string
	OtpCode   string
	ExpiresAt time.Time
	Verified  bool
}
