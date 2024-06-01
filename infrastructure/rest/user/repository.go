package user

import (
	"database/sql"
	"errors"
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
	// idとcreatedAtは自動で生成される
	query := `INSERT INTO users (username, password, email) VALUES ($1, $2, $3)`
	_, err := r.db.Exec(query, user.Name, user.Password, user.Email)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetByID(id string) (*model.UserResponse, error) {
	query := `SELECT id, username, email, created_at FROM users WHERE id = $1`
	row := r.db.QueryRow(query, id)
	user := &model.UserResponse{}
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return user, nil
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
