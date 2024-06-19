package service

import (
	"errors"
	"go-training/domain/model"
	"go-training/domain/repository"
)

type FavoriteService struct {
	favoriteRepo  repository.FavoriteRepository
	flashcardRepo repository.FlashcardRepository
}

func NewFavoriteService(favoriteRepo repository.FavoriteRepository, flashcardRepo repository.FlashcardRepository) *FavoriteService {
	return &FavoriteService{
		favoriteRepo:  favoriteRepo,
		flashcardRepo: flashcardRepo,
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
	studySets, err := s.favoriteRepo.GetFavoriteStudySetsByUserID(userID)

	if err != nil {
		return nil, err
	}

	//NOTICE: クエリをユーザの学習セット分行うから効率悪い
	// もっと良い方法ありそう
	for _, studySet := range studySets {
		flashcards, err := s.flashcardRepo.GetByStudySetID(studySet.ID)
		if err != nil {
			return nil, err
		}
		studySet.Flashcards = flashcards
	}

	return studySets, nil
}
