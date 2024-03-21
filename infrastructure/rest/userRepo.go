package userRepo

import (
	// "context"
	"go-training/domain/model"
	"go-training/domain/repository"
)

type UserRepositoryImpl struct{}

// 仮のデータリスト
type users []model.User

var usersList = users{{Name: "sas", ID: "33", Password: "hey"}, {Name: "you", ID: "44", Password: "sa"}, {Name: "mina", ID: "22", Password: "kin"}}

func NewUserRepositoryImpl() repository.UserRepository {
	return &UserRepositoryImpl{}
}

func (s *UserRepositoryImpl) GetPost(id string) ([]model.Post, error) {

	return []model.Post{}, nil
}

func (s *UserRepositoryImpl) CreatePost(model.Post) error {
	return nil
}

func (s *UserRepositoryImpl) Get() ([]model.User, error) {

	return usersList, nil
}

func (s *UserRepositoryImpl) Create(user *model.User) error {
	return nil
}

func (s *UserRepositoryImpl) FindByID(id string) (*model.User, error) {
	// ここにID検索のロジックを実装します。
	return &model.User{Name: "shin", ID: "33"}, nil
}

func (s *UserRepositoryImpl) Delete(id string) error {
	return nil
}
