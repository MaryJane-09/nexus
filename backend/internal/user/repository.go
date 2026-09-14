package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	users map[string]User
}

func (r *UserRepository) Create(user User) error {

	_, ok := r.users[user.Email]
	if ok {
		return errors.New("this email already exists")
	}
	hash := []byte(user.Password)
	hashed, err := bcrypt.GenerateFromPassword(hash, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashed)

	r.users[user.Email] = user

	return nil
}

func NewRepository() *UserRepository {
	return &UserRepository{
		users: make(map[string]User),
	}
}

func (r *UserRepository) FindByEmail(email string) (User, error) {
	foundUser, ok := r.users[email]
	if ok {
		return foundUser, nil
	}
	return foundUser, errors.New("User not found")
}

func (r *UserRepository) GetAllUsers() []UserResponse {
	var users []UserResponse

	for _, allUsers := range r.users {
		something := UserResponse{
			Name: allUsers.Name,
			Email: allUsers.Email,
		}
		users = append(users, something)
	}
	return users
}
