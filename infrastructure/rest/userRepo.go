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

func (repo *UserRepositoryImpl) FindByID(ctx context.Context, id uint) (string, error) {
	// ここにID検索のロジックを実装します。
	return "hehe", nil
}
