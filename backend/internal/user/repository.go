package user

import "errors"

type Repository struct{
	users map[string]User
}

func (r *Repository) Create(user User) error {

	_, ok := r.users[user.Email]
	if ok{
		return  errors.New("This email already exists")
	}

    r.users[user.Email] = user

	return nil
}

func NewRepository() *Repository{
	return &Repository{
		users: make(map[string]User),
	}
}

func (r *Repository) FindByEmail(email string) (User, error){
	foundUser, ok := r.users[email]
	if ok{
		return  foundUser, nil
	}
	return foundUser, errors.New("User not found")
}
