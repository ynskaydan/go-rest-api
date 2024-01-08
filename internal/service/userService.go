package service

import (
	"UserManagementProject/internal/domain"
	"UserManagementProject/internal/repository"
)

type UserService interface {
	GetAllUsers() ([]domain.User, error)
	GetUserByID(id int) (domain.User, error)
	CreateUser(user domain.User) error
	UpdateUser(user domain.User) error
	DeleteUser(id int) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

// CreateUser implements UserService.
func (s *userService) CreateUser(user domain.User) error {
	return s.repo.CreateUser(user)
}

func (s *userService) UpdateUser(user domain.User) error {
	return s.repo.UpdateUser(user)
}

// DeleteUser implements UserService.
func (s *userService) DeleteUser(id int) error {
	return s.repo.DeleteUser(id)
}

// GetAllUsers implements UserService.
func (s *userService) GetAllUsers() ([]domain.User, error) {
	return s.repo.GetAllUsers()
}

// GetUserByID implements UserService.
func (s *userService) GetUserByID(id int) (domain.User, error) {
	return s.repo.GetUser(id)
}
