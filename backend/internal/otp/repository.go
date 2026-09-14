package otp

import (
	"errors"
	"sync"
)

type OTPRepository struct {
	mu   sync.RWMutex
	otps map[string]OTP
}

func (r *OTPRepository) Generate(i int) (any, error) {
	panic("unimplemented")
}

func NewRepository() *OTPRepository {
	return &OTPRepository{
		otps: make(map[string]OTP),
	}
}

func (r *OTPRepository) Create(otp OTP) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.otps[otp.Email] = otp
	return nil
}

func (r *OTPRepository) FindByEmail(email string) (OTP, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	foundEmail, ok := r.otps[email]
	if ok {
		return foundEmail, nil
	}
	return OTP{}, errors.New("OTP was not found for this email")
}

func (r *OTPRepository) Delete(email string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.otps, email)
}
