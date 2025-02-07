package service

import (
	"encoding/json"
	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/ghostriderdev/housierBack/internal/dto"
	"github.com/ghostriderdev/housierBack/internal/mapper"
	"github.com/ghostriderdev/housierBack/internal/model"
	repository "github.com/ghostriderdev/housierBack/internal/repository/mem/data"
)

type AuthService struct {
	repository *repository.UserRepositoryMem
	mapper     *mapper.UserMapper
}

func NewUserService(repository *repository.UserRepositoryMem, mapper *mapper.UserMapper) *AuthService {
	return &AuthService{repository, mapper}
}

func (s AuthService) SignupUser(user dto.UserSignup) error {
	userModel := s.mapper.ToModel(user, "USER")

	userDB := s.repository.InsertUser(userModel)
	usersSaved := s.repository.GetAllUsers()
	usersSavedJSON, err := json.Marshal(usersSaved)
	if err != nil {
		return err
	}
	log.Println(string(usersSavedJSON))
	return userDB
}

func (s AuthService) SigninUser(user dto.UserSignin) *dto.SigninResponse {
	users := *s.repository.GetAllUsers()
	var userFound model.User
	for _, u := range users {
		if u.Email == user.Email && u.Password == user.Password {
			userFound = u
			break
		}
	}

	if (userFound == model.User{}) {
		return nil
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"email": userFound.Email,
	})

	secretKey := []byte("secret")
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return nil
	}

	return &dto.SigninResponse{
		Token: tokenString,
	}
}
