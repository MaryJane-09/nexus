package otp

import (
	"time"
)
const ExpiryTime = 5 * time.Minute

type OTP struct {
	Email     string
	Code   string
	ExpiresAt time.Time
	Verified  bool
}
