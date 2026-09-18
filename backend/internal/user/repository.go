package user

import (
	"context"
	"errors"
	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
	"sync"
)

type UserRepository struct {
	mu    sync.RWMutex
	users map[uuid.UUID]User
	pool  *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		users: make(map[uuid.UUID]User),
		pool:  pool,
	}
}

func (r *UserRepository) Create(user User) error {
	ctx := context.Background()
	hash := []byte(user.Password)
	hashed, err := bcrypt.GenerateFromPassword(hash, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashed)

	id, err := uuid.NewV4()
	if err != nil {
		return err
	}
	user.Id = id

	_, err = r.pool.Exec(ctx, `INSERT INTO users (id, name, email, password)
	VALUES ($1, $2, $3, $4)`, user.Id, user.Name, user.Email, user.Password)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return errors.New("this email already exists")
			}
		}
		return err
	}
	return nil
}

func (r *UserRepository) FindByID(id uuid.UUID) (User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	foundUser, ok := r.users[id]
	if ok {
		return foundUser, nil
	}
	return foundUser, errors.New("User not found")
}

func (r *UserRepository) FindByEmail(email string) (User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, value := range r.users {
		if value.Email == email {
			return value, nil
		}
	}
	return User{}, errors.New("User not found")
}

func (r *UserRepository) GetAllUsers() []UserResponse {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var users []UserResponse

	for _, allUsers := range r.users {
		something := UserResponse{
			Name:  allUsers.Name,
			Email: allUsers.Email,
		}
		users = append(users, something)
	}
	return users
}
