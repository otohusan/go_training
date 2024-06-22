package repository

import "go-training/domain/model"

type UserRepository interface {
	CreateWithEmail(user *model.User) (string, error)
	CreateWithGoogle(user *model.User) (string, error)
	GetByID(id string) (*model.UserResponse, error)
	GetByUsername(username string) (*model.UserResponse, error)
	Update(user *model.User) error
	Delete(id string) error
	// 名前良くないかも
	// これはログイン時に使うメソッドだからパスワードを含む構造体を使う
	GetByEmail(email string) (*model.User, error)
	IsEmailExist(email string) (bool, error)
	IsUsernameExist(username string) (bool, error)
	GetAll() ([]*model.User, error)
}
