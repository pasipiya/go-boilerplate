package service

import (
	"errors"
	"github.com/pasipiya/go-boilerplate/internal/model"
	"github.com/pasipiya/go-boilerplate/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (s *UserService) GetAllUsers() ([]model.User, error) {
	return s.repo.FindAll()
}

func (s *UserService) GetUserByID(id int) (*model.User, error) {
	return s.repo.FindByID(id)
}

func (s *UserService) CreateUser(req CreateUserRequest) (*model.User, error) {
	if req.Name == "" || req.Email == "" {
		return nil, errors.New("invalid input")
	}
	return s.repo.Create(req.Name, req.Email)
}
