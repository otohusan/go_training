package favorite

import (
	"database/sql"
	"go-training/domain/model"
)

type FavoriteRepository struct {
	db *sql.DB
}

func NewFavoriteRepository(db *sql.DB) *FavoriteRepository {
	return &FavoriteRepository{db: db}
}

func (r *FavoriteRepository) AddFavorite(userID, studySetID string) error {
	return nil
}

func (r *FavoriteRepository) RemoveFavorite(userID, studySetID string) error {
	return nil
}

func (r *FavoriteRepository) GetFavoritesByUserID(userID string) ([]*model.Favorite, error) {
	return nil, nil
}

func (r *FavoriteRepository) GetFavoriteStudySetsByUserID(userID string) ([]*model.StudySet, error) {
	return nil, nil
}
