package dto

// UserSignin defines a user data to login
type UserSignin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}