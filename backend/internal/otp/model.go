package otp

import (
	"time"
)
const ExpiryTime = Time * time.Minute
const Time = 5

type OTP struct {
	Email     string
	Code   string
	ExpiresAt time.Time
	Verified  bool
}
