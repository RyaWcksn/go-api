package services

import (
	"github.com/RyaWcksn/go-api/entities"
	"github.com/RyaWcksn/go-api/repositories"
)

type UserInterface interface {
    Create(user entities.User) (entities.User, error)
    GetUser(id int) (entities.User, error)
    GetUsers() ([]entities.User, error)
}

type UserService struct {
    repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) *UserService {
    return &UserService{repo}
}

func (s UserService) Create(user entities.User) (entities.User, error) {
    user, err := s.Create(user)
    if err != nil {
        panic(err.Error())
    }
    return user, nil
}
