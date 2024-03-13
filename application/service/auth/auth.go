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

func (h *AuthService) RegisterUser(user model.User) error {
	return h.repo.Create(&user)
}
