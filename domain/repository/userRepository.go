package repository

import (
	// "context"
	"go-training/domain/model"
)

// UserRepository はユーザーリポジトリのインターフェースです。
type UserRepository interface {
	Get() ([]model.User, error)
	FindByID(id uint) (*model.User, error)
	Create(user *model.User) error
	Delete(id uint) error
}
