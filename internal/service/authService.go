package service

import (
	"github.com/ghostriderdev/housierBack/internal/dto"
	"github.com/ghostriderdev/housierBack/internal/mapper"
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
	return s.repository.InsertUser(userModel)
}
