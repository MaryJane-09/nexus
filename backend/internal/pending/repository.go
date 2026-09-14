package pending

import (
	"errors"
	"sync"
	"github.com/MaryJane-09/nexus/backend/internal/user"
)

type PendingRepository struct {
	mu    sync.RWMutex
	users map[string]user.User
}

func NewRepository() *PendingRepository {
	return &PendingRepository{
		users: make(map[string]user.User),
	}
}

func (r *PendingRepository) Create(info user.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.users[info.Email] = info
	return nil
}

func (r *PendingRepository) FindByEmail(email string) (user.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	foundEmail, ok := r.users[email]
	if ok {
		return foundEmail, nil
	}
	return user.User{}, errors.New("Pending user was not found")
}


func (r *PendingRepository) Delete(email string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.users, email)
}