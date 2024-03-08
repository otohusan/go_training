package repository

import (
	"context"
	"go-training/domain/model"
	"go-training/domain/repository"
)

type UserRepositoryImpl struct{}

type users []model.User

var usersList = users{{Name: "sas", ID: 33}, {Name: "you", ID: 39}}

func NewUserRepositoryImpl() repository.UserRepository {
	return &UserRepositoryImpl{}
}

func (s *UserRepositoryImpl) Create(ctx context.Context, user *model.User) error {

	return nil
}

func (s *UserRepositoryImpl) FindByID(ctx context.Context, id uint) (*model.User, error) {
	// ここにID検索のロジックを実装します。
	return &usersList[1], nil
}
