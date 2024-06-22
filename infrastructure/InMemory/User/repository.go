package user

import (
	"errors"
	"go-training/domain/model"
	"go-training/domain/repository"
	inmemory "go-training/infrastructure/InMemory"
	"go-training/utils"
	"sync"

	"github.com/google/uuid"
)

type UserRepository struct {
	mu sync.Mutex
}

func NewUserRepository() repository.UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) CreateWithEmail(user *model.User) (string, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, userSet := range inmemory.Users {
		if userSet.Email == user.Email {
			return "", errors.New("the email can't use")
		}
	}

	// uuid作成
	user.ID = uuid.New().String()

	inmemory.Users = append(inmemory.Users, user)
	return user.ID, nil
}

func (r *UserRepository) CreateWithGoogle(user *model.User) (string, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, userSet := range inmemory.Users {
		if userSet.Email == user.Email {
			return "", errors.New("the email can't use")
		}
	}

	// uuid作成
	user.ID = uuid.New().String()

	inmemory.Users = append(inmemory.Users, user)
	return user.ID, nil
}

func (r *UserRepository) GetByID(id string) (*model.UserResponse, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, user := range inmemory.Users {
		if user.ID == id {
			// UserからUserResponseへの変換
			userResponse := &model.UserResponse{
				ID:        user.ID,
				Name:      user.Name,
				Email:     user.Email,
				CreatedAt: user.CreatedAt,
			}
			return userResponse, nil
		}
	}

	return nil, errors.New("user not found")

}

func (r *UserRepository) GetByUsername(username string) (*model.UserResponse, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, user := range inmemory.Users {
		if user.Name == username {
			// UserからUserResponseへの変換
			userResponse := &model.UserResponse{
				ID:        user.ID,
				Name:      user.Name,
				Email:     user.Email,
				CreatedAt: user.CreatedAt,
			}
			return userResponse, nil
		}
	}
	return nil, errors.New("user not found")
}

func (r *UserRepository) Update(user *model.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, userFromDB := range inmemory.Users {
		if userFromDB.ID == user.ID {
			// 変更可能な場所のみを変更する
			inmemory.Users[i].Name = user.Name
			return nil
		}
	}

	return errors.New("user not found")

}

func (r *UserRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	// TODO:本人確認が必要

	for i, userSet := range inmemory.Users {
		if userSet.ID == id {
			inmemory.Users = utils.RemoveElementFromSlice(inmemory.Users, i)
			return nil
		}
	}

	return errors.New("user not found")
}

func (r *UserRepository) GetAll() ([]*model.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	return inmemory.Users, nil
}

func (r *UserRepository) GetByEmail(email string) (*model.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	for _, user := range inmemory.Users {
		if user.Email.String == email {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (r *UserRepository) IsEmailExist(email string) (bool, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, user := range inmemory.Users {
		if user.Email.String == email {
			return true, nil
		}
	}
	return false, nil
}

func (r *UserRepository) IsUsernameExist(username string) (bool, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, user := range inmemory.Users {
		if user.Name == username {
			return true, nil
		}
	}
	return false, nil
}
