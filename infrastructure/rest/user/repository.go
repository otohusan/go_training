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
	query := `INSERT INTO users (username, password, email) VALUES ($1, $2, $3) RETURNING id, created_at`
	// 作成と、作成されたIDをJWTを生成するためにuserにすぐ割り当てている
	err := r.db.QueryRow(query, user.Name, user.Password, user.Email).Scan(&user.ID, &user.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetByID(id string) (*model.UserResponse, error) {
	query := `SELECT id, username, email, created_at 
			  FROM users 
			  WHERE id = $1`
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
	// 製品で使う予定がまだないから開発しない
	// usernameはuniqueじゃないから複数の値を返すのに、配列の返り値なってない
	return nil, errors.New("this func is still developing")

	// return &model.UserResponse{}, nil
}

func (r *UserRepository) Update(user *model.User) error {
	query := `UPDATE users SET username = $1 WHERE id = $2`
	_, err := r.db.Exec(query, user.Name, user.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) Delete(id string) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetAll() ([]*model.User, error) {
	// 製品で使う予定がまだないから開発しない
	// usernameはuniqueじゃないから複数の値を返すのに、配列の返り値なってない
	return nil, errors.New("this func is still developing")

	// return []*model.User{}, nil
}

// loginの確認で使用するからuserの全情報を返す
func (r *UserRepository) GetByEmail(email string) (*model.User, error) {
	query := `SELECT id, username, password, email, created_at 
			  FROM users 
			  WHERE email = $1`
	row := r.db.QueryRow(query, email)
	user := &model.User{}
	err := row.Scan(&user.ID, &user.Name, &user.Password, &user.Email, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) IsEmailExist(email string) (bool, error) {
	query := `SELECT 1 FROM users WHERE email = $1`
	row := r.db.QueryRow(query, email)

	var exists int
	err := row.Scan(&exists)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
