package user

import (
	"database/sql"
	"go-training/domain/model"
	"go-training/domain/repository"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateWithEmail(user *model.User) error {
	query := `INSERT INTO users (username, password, email) VALUES ($1, $2, $3)`
	_, err := r.db.Exec(query, user.Name, user.Password, user.Email)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetByID(id string) (*model.UserResponse, error) {
	return &model.UserResponse{}, nil
}

func (r *UserRepository) GetByUsername(username string) (*model.UserResponse, error) {
	return &model.UserResponse{}, nil
}

func (r *UserRepository) Update(user *model.User) error {
	return nil
}

func (r *UserRepository) Delete(id string) error {
	return nil
}

func (r *UserRepository) GetAll() ([]*model.User, error) {
	return []*model.User{}, nil
}

func (r *UserRepository) GetByEmail(email string) (*model.User, error) {
	return &model.User{}, nil
}
