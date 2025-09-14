package service

import (
	"github.com/Azm1l/rest-api-go/entity"
	"github.com/Azm1l/rest-api-go/repository"
)

type UserService interface {
	CreateUser(user *entity.User) (*entity.User, error)
	ShowAllUsers() ([]*entity.User, error)
	FindUserByID(id int64) (*entity.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) CreateUser(user *entity.User) (*entity.User, error) {
	return s.repo.Create(user)
}

func (s *userService) ShowAllUsers() ([]*entity.User, error) {
	return s.repo.ShowAll()
}

func (s *userService) FindUserByID(id int64) (*entity.User, error) {
	return s.repo.FindOne(id)
}
