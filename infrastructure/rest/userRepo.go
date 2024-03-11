package userRepo

import (
	"context"
	"go-training/domain/model"
	"go-training/domain/repository"
)

type UserRepositoryImpl struct{}

func NewUserRepositoryImpl() repository.UserRepository {
	return &UserRepositoryImpl{}
}

func (s *UserRepositoryImpl) Create(ctx context.Context, user *model.User) error {
	return nil
}

func (s *UserRepositoryImpl) FindByID(id uint) (*model.User, error) {
	// ここにID検索のロジックを実装します。
	return &model.User{Name: "shin", ID: 228}, nil
}
