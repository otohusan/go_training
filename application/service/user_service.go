package service

import (
	"errors"
	"go-training/domain/model"
	"go-training/domain/repository"
)

// userが適切に値を持ってるか確かめる関数
func validateUser(user *model.User) error {
	if user.Name == "" {
		return errors.New("username cannot be empty")
	}
	if user.Password == "" {
		return errors.New("password cannot be empty")
	}
	return nil
}

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) CreateUserWithEmail(user *model.User) (string, error) {
	if err := validateUser(user); err != nil {
		return "", err
	}
	if !user.Email.Valid || user.Email.String == "" {
		return "", errors.New("email cannot be empty")
	}

	return s.repo.CreateWithEmail(user)
}

func (s *UserService) GetUserByID(id string) (*model.UserResponse, error) {
	return s.repo.GetByID(id)
}

func (s *UserService) GetUserByUsername(username string) (*model.UserResponse, error) {
	return s.repo.GetByUsername(username)
}

func (s *UserService) GetPublicUserInfo(userID string) (*model.PublicUser, error) {
	user, err := s.repo.GetByID(userID)
	if err != nil {
		return nil, err
	}

	// 一般公開用の情報をフィルタリング
	publicUser := &model.PublicUser{
		ID:   user.ID,
		Name: user.Name,
	}

	return publicUser, nil
}

func (s *UserService) UpdateUser(authUserID string, user *model.User) error {
	if user.Name == "" {
		return errors.New("username cannot be empty")
	}

	// 認可できるか
	if user.ID != authUserID {
		return errors.New("not authorized")
	}

	return s.repo.Update(user)
}

func (s *UserService) DeleteUser(authUserID string, id string) error {
	// 認可できるか
	if id != authUserID {
		return errors.New("not authorized")
	}

	return s.repo.Delete(id)
}

func (s *UserService) GetUserByEmail(email string) (*model.User, error) {
	return s.repo.GetByEmail(email)
}

func (s *UserService) GetAllUsers() ([]*model.User, error) {
	return s.repo.GetAll()
}

func (s *UserService) IsEmailExist(email string) (bool, error) {
	return s.repo.IsEmailExist(email)
}

func (s *UserService) IsUsernameExist(username string) (bool, error) {
	return s.repo.IsUsernameExist(username)
}
