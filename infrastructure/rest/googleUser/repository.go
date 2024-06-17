package googleUser

import (
	"database/sql"
	"go-training/domain/model"
)

type GoogleUserRepositoryImpl struct {
	db *sql.DB
}

func NewGoogleUserRepository(db *sql.DB) *GoogleUserRepositoryImpl {
	return &GoogleUserRepositoryImpl{db: db}
}

func (r *GoogleUserRepositoryImpl) Create(googleUser *model.GoogleUser) error {
	query := `INSERT INTO google_users (user_id, google_id) VALUES ($1, $2)`
	_, err := r.db.Exec(query, googleUser.UserID, googleUser.GoogleID)
	return err
}

func (r *GoogleUserRepositoryImpl) GetByGoogleID(googleID string) (*model.GoogleUser, error) {
	query := `SELECT id, user_id, google_id, created_at FROM google_users WHERE google_id = $1`
	row := r.db.QueryRow(query, googleID)
	googleUser := &model.GoogleUser{}
	err := row.Scan(&googleUser.ID, &googleUser.UserID, &googleUser.GoogleID, &googleUser.CreatedAt)
	if err != nil {
		// 見つけれなかった場合は問題ない
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return googleUser, nil
}
