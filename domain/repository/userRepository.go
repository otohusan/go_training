package repository

import "go-training/domain/model"

type UserRepository interface {
	CreateWithEmail(user *model.User) error
	GetByID(id string) (*model.UserResponse, error)
	GetByUsername(username string) (*model.User, error)
	Update(user *model.User) error
	Delete(id string) error
	GetByEmail(email string) (*model.User, error)
	GetAll() ([]*model.User, error)
}
