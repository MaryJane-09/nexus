package otp

import (
	"errors"
	"sync"
)

var ErrNotFound = errors.New("OTP was not found for this email")

type Repository struct {
	mu   sync.RWMutex 
	otps map[string]OTP
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
	return OTP{}, ErrNotFound
}

func (r *Repository) Delete(email string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	
	delete(r.otps, email)
}
