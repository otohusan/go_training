package inmemory

import (
	"errors"
	"fmt"
	"go-training/domain/model"
	"go-training/domain/repository"
	inmemory "go-training/infrastructure/InMemory"
	"sync"
)

type UserRepository struct {
	mu    sync.Mutex
	users map[string]*model.User
}

func NewUserRepository() repository.UserRepository {
	repo := &UserRepository{
		users: make(map[string]*model.User),
	}
	for _, user := range inmemory.InitializeUsers() {
		repo.users[user.ID] = user
	}
	return repo
}

func (r *UserRepository) Create(user *model.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// ユーザーの数を基に新しいIDを生成 TODO:削除とかあったら、id被る
	user.ID = fmt.Sprintf("%d", len(r.users)+1)

	if _, exists := r.users[user.ID]; exists {
		return errors.New("user already exists")
	}

	r.users[user.ID] = user
	return nil
}

func (r *UserRepository) GetByID(id string) (*model.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	user, exists := r.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (r *UserRepository) GetByUsername(username string) (*model.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, user := range r.users {
		if user.Name == username {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (r *UserRepository) Update(user *model.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.users[user.ID]; !exists {
		return errors.New("user not found")
	}

	r.users[user.ID] = user
	return nil
}

func (r *UserRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.users[id]; !exists {
		return errors.New("user not found")
	}

	delete(r.users, id)
	return nil
}

func (r *UserRepository) GetAll() ([]*model.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var users []*model.User
	for _, user := range r.users {
		users = append(users, user)
	}
	return users, nil
}
