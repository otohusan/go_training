package service

import (
	"errors"
	"go-training/domain/model"
	"go-training/domain/repository"
)

// userが適切に値を持ってるか確かめる関数
func validateUser(user *model.User) error {
	if user.Name == "" {
		return errors.New("username cannot be empty")
	}
	if user.Password == "" {
		return errors.New("password cannot be empty")
	}
	return nil
}

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) CreateUserWithEmail(user *model.User) error {
	if err := validateUser(user); err != nil {
		return err
	}
	if user.Email == "" {
		return errors.New("email cannot be empty")
	}

	return s.repo.CreateWithEmail(user)
}

func (s *UserService) GetUserByID(id string) (*model.User, error) {
	return s.repo.GetByID(id)
}

func (s *UserService) GetUserByUsername(username string) (*model.User, error) {
	return s.repo.GetByUsername(username)
}

func (s *UserService) UpdateUser(user *model.User) error {
	if err := validateUser(user); err != nil {
		return err
	}

	return s.repo.Update(user)
}

func (s *UserService) DeleteUser(id string) error {
	return s.repo.Delete(id)
}

func (s *UserService) GetUserByEmail(email string) (*model.User, error) {
	return s.repo.GetByEmail(email)
}

func (s *UserService) GetAllUsers() ([]*model.User, error) {
	return s.repo.GetAll()
}
