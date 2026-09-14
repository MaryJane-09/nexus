package user

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
