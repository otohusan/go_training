package userRepo

import (
	// "context"
	"go-training/domain/model"
	"go-training/domain/repository"
)

type UserRepositoryImpl struct{}

// 仮のデータリスト
type users []model.User

var usersList = users{{Name: "sas", ID: 33}, {Name: "you", ID: 39}, {Name: "mina", ID: 21}, {Name: "mina", ID: 31}}

func NewUserRepositoryImpl() repository.UserRepository {
	return &UserRepositoryImpl{}
}

func (s *UserRepositoryImpl) Get() ([]model.User, error) {

	return usersList, nil
}

func (s *UserRepositoryImpl) Create(user *model.User) error {
	return nil
}

func (s *UserRepositoryImpl) FindByID(id uint) (*model.User, error) {
	// ここにID検索のロジックを実装します。
	return &model.User{Name: "shin", ID: 228}, nil
}
