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
	mu    sync.Mutex
	users map[string]*model.User
}

func NewUserRepository() repository.UserRepository {
	repo := &UserRepository{
		users: make(map[string]*model.User),
	}
	for _, user := range inmemory.Users {
		repo.users[user.ID] = user
	}
	return repo
}

func (r *UserRepository) CreateWithEmail(user *model.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, userSet := range inmemory.Users {
		if userSet.Email == user.Email {
			return errors.New("the email can't use")
		}
	}

	// uuid作成
	user.ID = uuid.New().String()

	inmemory.Users = append(inmemory.Users, user)
	return nil
}

func (r *UserRepository) GetByID(id string) (*model.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, user := range inmemory.Users {
		if user.ID == id {
			return user, nil
		}
	}

	return nil, errors.New("user not found")

}

func (r *UserRepository) GetByUsername(username string) (*model.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, user := range inmemory.Users {
		if user.Name == username {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (r *UserRepository) Update(user *model.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// TODO:本人確認が必要
	for i, userFromDB := range inmemory.Users {
		if userFromDB.ID == user.ID {
			inmemory.Users[i] = user
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
		if user.Email == email {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}
