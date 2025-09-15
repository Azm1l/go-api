package service

import (
	"errors"

	"github.com/Azm1l/rest-api-go/dto"
	"github.com/Azm1l/rest-api-go/entity"
	"github.com/Azm1l/rest-api-go/repository"
	"github.com/Azm1l/rest-api-go/utils"
)

type UserService interface {
	CreateUser(req dto.CreateUserRequest) (*entity.User, error)
	ShowAllUsers() ([]entity.User, error)
	FindUserByID(id int64) (*entity.User, error)
	UpdateUser(id int64, req dto.UpdateUserRequest) (*entity.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) CreateUser(req dto.CreateUserRequest) (*entity.User, error) {

	existingUser, _ := s.repo.FindByEmail(req.Email)

	if existingUser != nil {
		return nil, errors.New("email already in use")
	}

	hashPassword, err := utils.HashPassword(req.Password)

	if err != nil {
		return nil, err
	}

	user := &entity.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashPassword,
	}

	return s.repo.Create(user)
}

func (s *userService) ShowAllUsers() ([]entity.User, error) {
	return s.repo.ShowAll()
}

func (s *userService) FindUserByID(id int64) (*entity.User, error) {
	return s.repo.FindById(id)
}

func (s *userService) UpdateUser(id int64, req dto.UpdateUserRequest) (*entity.User, error) {
	user, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}
	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Email != "" {
		existingUser, _ := s.repo.FindByEmail(req.Email)
		if existingUser != nil && existingUser.ID != id {
			return nil, errors.New("email already in use")
		}
		user.Email = req.Email
	}
	if req.Password != "" {
		hashPassword, err := utils.HashPassword(req.Password)
		if err != nil {
			return nil, err
		}
		user.Password = hashPassword
	}
	return s.repo.Update(user)
}
