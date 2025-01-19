package mapper

import (
	"github.com/ghostriderdev/housierBack/internal/dto"
	"github.com/ghostriderdev/housierBack/internal/model"
)

type UserMapper struct {
}

func NewUserMapper() *UserMapper {
	return &UserMapper{}
}

func (m UserMapper) ToModel(dto dto.UserSignup, role string) *model.User {
	return &model.User{
		Email:    dto.Email,
		Password: dto.Password,
		Role:     role,
	}
}
