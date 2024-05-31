package service

import (
	"errors"
	"go-training/domain/model"
	"go-training/domain/repository"
)

type FavoriteService struct {
	favoriteRepo repository.FavoriteRepository
}

func NewFavoriteService(favoriteRepo repository.FavoriteRepository) *FavoriteService {
	return &FavoriteService{
		favoriteRepo: favoriteRepo,
	}
}

func (s *FavoriteService) AddFavorite(authUserID, userID, studySetID string) error {
	// 認可できるか
	if authUserID != userID {
		return errors.New("not authorized")
	}
	return s.favoriteRepo.AddFavorite(userID, studySetID)
}

func (s *FavoriteService) RemoveFavorite(authUserID, userID, studySetID string) error {
	// 認可できるか
	if authUserID != userID {
		return errors.New("not authorized")
	}
	return s.favoriteRepo.RemoveFavorite(userID, studySetID)
}

func (s *FavoriteService) GetFavoritesByUserID(userID string) ([]*model.Favorite, error) {
	return s.favoriteRepo.GetFavoritesByUserID(userID)
}

func (s *FavoriteService) IsFavorite(userID, studySetID string) (bool, error) {
	return s.favoriteRepo.IsFavorite(userID, studySetID)
}

func (s *FavoriteService) GetFavoriteStudySetsByUserID(userID string) ([]*model.StudySet, error) {
	return s.favoriteRepo.GetFavoriteStudySetsByUserID(userID)
}
