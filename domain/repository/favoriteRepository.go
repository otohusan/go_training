package repository

import "go-training/domain/model"

type FavoriteRepository interface {
	AddFavorite(userID, studySetID string) error
	RemoveFavorite(userID, studySetID string) error
	GetFavoritesByUserID(userID string) ([]*model.Favorite, error)
	IsFavorite(userID, studySetID string) (bool, error)
}
