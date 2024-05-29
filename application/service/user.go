package service

import (
	// "context"
	"go-training/domain/model"
	"go-training/domain/repository"
)

type UserService struct {
	repo repository.UserRepository
}

// NewUserService は新しいUserServiceオブジェクトを作成します。
func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

// func (s *UserService) GetUserList() ([]model.User, error) {
// 	return s.repo.Get()
// }

func (s *UserService) ReturnUser(id string) (*model.User, error) {
	return s.repo.GetByID(id)
}

func (s *UserService) Create(user *model.User) error {
	return s.repo.Create(user)
}

func (s *UserService) DeleteUserByID(id string) error {
	return s.repo.Delete(id)
}

// func (s *UserService) GetPosts(id string) ([]model.Post, error) {
// 	return s.repo.GetPost(id)
// }

// func (s *UserService) CreatePost(post model.Post) error {
// 	return s.repo.CreatePost(post)
// }
