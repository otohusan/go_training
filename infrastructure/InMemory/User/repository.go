package repository

import (
	"context"
	"errors"
	"go-training/domain/model"
	"go-training/domain/repository"
)

type UserRepositoryImpl struct{}

type users []model.User

var usersList = users{{Name: "sas", ID: 33}, {Name: "you", ID: 39}, {Name: "mina", ID: 21}, {Name: "mina", ID: 31}}

func NewUserRepositoryImpl() repository.UserRepository {
	return &UserRepositoryImpl{}
}

func (s *UserRepositoryImpl) Create(ctx context.Context, user *model.User) error {

	return nil
}

func (s *UserRepositoryImpl) FindByID(id uint) (*model.User, error) {
	userIndex := -1

	for i, v := range usersList {
		if v.ID == id {
			userIndex = i
			break
		}
	}

	if userIndex == -1 {
		return nil, errors.New("userが見つかりません")
	}
	// ここにID検索のロジックを実装します。
	return &usersList[1], nil
}
