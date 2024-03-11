package repository

import (
	// "context"
	"go-training/domain/model"
)

// UserRepository はユーザーリポジトリのインターフェースです。
type UserRepository interface {
	FindByID(id uint) (*model.User, error)
	Create(user *model.User) error
	Get() ([]model.User, error)
}
