package otp

import (
	"errors"
	"sync"
)

type Repository struct {
	mu   sync.RWMutex
	otps map[string]OTP
}

func (r *Repository) Generate(i int) (any, error) {
	panic("unimplemented")
}

func NewRepository() *Repository {
	return &Repository{
		otps: make(map[string]OTP),
	}
}

func (r *Repository) Create(otp OTP) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.otps[otp.Email] = otp
	return nil
}

func (r *Repository) FindByEmail(email string) (OTP, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	foundEmail, ok := r.otps[email]
	if ok {
		return foundEmail, nil
	}
	return OTP{}, errors.New("OTP was not found for this email")
}

func (r *Repository) Delete(email string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.otps, email)
}
