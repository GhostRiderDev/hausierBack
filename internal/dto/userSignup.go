package dto

// UserSignupDto defines a user data to signup
type UserSignup struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
