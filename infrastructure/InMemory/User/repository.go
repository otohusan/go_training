package repository

import (
	// "context"
	"errors"
	"go-training/domain/model"
	"go-training/domain/repository"
	"strconv"
)

type UserRepositoryImpl struct{}

type users []model.User

var usersList = users{{Name: "sas", ID: "33", Password: "hey"}, {Name: "you", ID: "44", Password: "sa"}, {Name: "mina", ID: "22", Password: "kin"}}

func NewUserRepositoryImpl() repository.UserRepository {
	return &UserRepositoryImpl{}
}

func (s *UserRepositoryImpl) Get() ([]model.User, error) {

	return usersList, nil
}

func (s *UserRepositoryImpl) Create(user *model.User) error {
	id := int(len(usersList) + 1)
	CreatedUser := model.User{Name: user.Name, ID: strconv.Itoa(id), Password: user.Password}
	usersList = append(usersList, CreatedUser)
	return nil
}

func (s *UserRepositoryImpl) FindByID(id string) (*model.User, error) {
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

	return &usersList[userIndex], nil
}

func (s *UserRepositoryImpl) Delete(id string) error {
	userIndex := -1

	for i, v := range usersList {
		if v.ID == id {
			userIndex = i
			break
		}
	}

	if userIndex == -1 {
		return errors.New("userが見つかりません")
	}

	usersList = append(usersList[:userIndex], usersList[userIndex+1:]...)
	// ここにID検索のロジックを実装します。
	return nil
}
