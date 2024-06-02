package favorite

import (
	"database/sql"
	"errors"
	"go-training/domain/model"
)

type FavoriteRepository struct {
	db *sql.DB
}

func NewFavoriteRepository(db *sql.DB) *FavoriteRepository {
	return &FavoriteRepository{db: db}
}

func (r *FavoriteRepository) AddFavorite(userID, studySetID string) error {
	query := `INSERT INTO favorites (user_id, study_set_id) VALUES ($1, $2)`
	_, err := r.db.Exec(query, userID, studySetID)
	if err != nil {
		return err
	}

	return nil
}

func (r *FavoriteRepository) RemoveFavorite(userID, studySetID string) error {
	query := `DELETE FROM favorites WHERE user_id = $1 AND study_set_id = $2`
	result, err := r.db.Exec(query, userID, studySetID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("favorite not found")
	}

	return nil
}

func (r *FavoriteRepository) IsFavorite(userID, studySetID string) (bool, error) {
	return false, nil
}

func (r *FavoriteRepository) GetFavoritesByUserID(userID string) ([]*model.Favorite, error) {
	return nil, nil
}

func (r *FavoriteRepository) GetFavoriteStudySetsByUserID(userID string) ([]*model.StudySet, error) {
	return nil, nil
}
