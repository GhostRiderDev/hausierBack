package controller

import (
	"encoding/json"
	"net/http"

	"github.com/ghostriderdev/housierBack/internal/dto"
	"github.com/ghostriderdev/housierBack/internal/service"
)

type AuthController struct {
	Service *service.AuthService
}

func (c *AuthController) Signup(w http.ResponseWriter, r *http.Request) {
	var user dto.UserSignup
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid request payload"))
		return
	}

	if user.Email == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid Email"))
		return
	}

	if user.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid Password"))
		return
	}

	c.Service.SignupUser(user)

	// Aquí puedes agregar la lógica para manejar el registro del usuario
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User signed up successfully"))
}
