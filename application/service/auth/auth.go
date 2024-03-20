package auth

import (
	"go-training/domain/model"
	"go-training/domain/repository"
)

type AuthService struct {
	repo repository.UserRepository
}

func NewAuthService(repo repository.UserRepository) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (h *AuthService) RegisterUser(user model.CreatedUserData) error {
	userInfo := model.User{
		Name:     user.Username,
		Password: user.Password,
	}
	return h.repo.Create(&userInfo)
}
