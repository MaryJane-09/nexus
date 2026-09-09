package otp

import (
	"errors"
	"time"
)

type OTP struct {
	Email     string
	OtpCode   string
	ExpiresAt time.Time
	Verified  bool
}

type Repository struct {
	otps map[string]OTP
}

func NewRepository() *Repository {
	return &Repository{
		otps: make(map[string]OTP),
	}
}

func (r *Repository) Create(otp OTP) {
	r.otps[otp.Email] = otp
}

func (r *Repository) FindByEmail(email string) (OTP, error) {
	foundEmail, ok := r.otps[email]
	if ok {
		return foundEmail, nil
	}
	return foundEmail, errors.New("email not found")
}
