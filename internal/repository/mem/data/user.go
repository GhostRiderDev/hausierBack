package repository

import "github.com/ghostriderdev/housierBack/internal/model"

type UserRepositoryMem struct {
	users *[]model.User
}

func NewUserRepository() *UserRepositoryMem {
	return &UserRepositoryMem{&[]model.User{}}
}

func (r UserRepositoryMem) InsertUser(user *model.User) error {
	users := *r.users
	*r.users = append(users, *user)
	return nil
}

func (r UserRepositoryMem) GetAllUsers() *[]model.User {
	return r.users
}
