package repository

import (
	"context"
	"go-training/domain/model"
)

// UserRepository はユーザーリポジトリのインターフェースです。
type UserRepository interface {
	FindByID(id uint) (*model.User, error)
	Create(ctx context.Context, user *model.User) error
}
