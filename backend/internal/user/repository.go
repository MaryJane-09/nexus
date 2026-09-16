package user

import (
	"errors"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	id map[uuid.UUID]User
}

func (r *UserRepository) Create(user User) error {

	for _, value := range r.users {
		if value.Email == user.Email {
			return errors.New("this email already exists")
		}
	}
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

	r.id[user.Id] = user

	return nil
}

func NewRepository() *UserRepository {
	return &UserRepository{
		id: make(map[uuid.UUID]User),
	}
}

func (r *UserRepository) FindByEmail(id uuid.UUID) (User, error) {
	foundUser, ok := r.id[id]
	if ok {
		return foundUser, nil
	}
	return foundUser, errors.New("User not found")
}

func (r *UserRepository) GetAllUsers() []UserResponse {
	var users []UserResponse

	for _, allUsers := range r.id {
		something := UserResponse{
			Name:  allUsers.Name,
			Email: allUsers.Email,
		}
		users = append(users, something)
	}
	return users
}
