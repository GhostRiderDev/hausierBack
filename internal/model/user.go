package model

const (
	ROLE_USER  = "USER"
	ROLE_ADMIN = "ADMIN"
)

// User defines a user basic
type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
