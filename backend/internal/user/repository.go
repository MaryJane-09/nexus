package user

import (
	"errors"
)

type UserRepository struct {
	users map[string]User
}

func (r *UserRepository) Create(user User) error {

	_, ok := r.users[user.Email]
	if ok {
		return errors.New("this email already exists")
	}

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

func (r *UserRepository) GetAllUsers() []User {
	var users []User

	for _, allUsers := range r.users {
		users = append(users, allUsers)
	}
	return users
}
