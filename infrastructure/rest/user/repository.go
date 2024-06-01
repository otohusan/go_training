package user

import (
	"go-training/domain/model"
	"go-training/domain/repository"
)

type UserRepository struct {
}

func NewUserRepository() repository.UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) CreateWithEmail(user *model.User) error {
	return nil
}

func (r *UserRepository) GetByID(id string) (*model.User, error) {
	return &model.User{}, nil
}

func (r *UserRepository) GetByUsername(username string) (*model.User, error) {
	return &model.User{}, nil
}

func (r *UserRepository) Update(user *model.User) error {
	return nil
}

func (r *UserRepository) Delete(id string) error {
	return nil
}

func (r *UserRepository) GetAll() ([]*model.User, error) {
	return []*model.User{}, nil
}

func (r *UserRepository) GetByEmail(email string) (*model.User, error) {
	return &model.User{}, nil
}
