package favorite

import (
	"errors"
	"go-training/domain/model"
	"go-training/domain/repository"
	inmemory "go-training/infrastructure/InMemory"
	"go-training/utils"
	"sync"
	"time"

	"github.com/google/uuid"
)

type FavoriteRepository struct {
	mu sync.Mutex
}

func NewFavoriteRepository() repository.FavoriteRepository {
	return &FavoriteRepository{}
}

func (r *FavoriteRepository) AddFavorite(userID, studySetID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	// TODO:本人確認が必要

	// 外部キーのチェック: UserIDが存在するか
	isUserExists := false
	for _, user := range inmemory.Users {
		if user.ID == userID {
			isUserExists = true
			break
		}
	}
	if !isUserExists {
		return errors.New("user doesn't exist")
	}
	isStudySetExists := false
	for _, studySet := range inmemory.StudySets {
		if studySet.ID == studySetID {
			isStudySetExists = true
			break
		}
	}
	if !isStudySetExists {
		return errors.New("studySet doesn't exists")
	}

	// 新しいIDをUUIDで生成
	tableID := uuid.New().String()
	newFavorite := &model.Favorite{
		ID:         tableID,
		UserID:     userID,
		StudySetID: studySetID,
		CreatedAt:  time.Now(),
	}
	inmemory.Favorites = append(inmemory.Favorites, newFavorite)
	return nil
}

func (r *FavoriteRepository) RemoveFavorite(userID, studySetID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	// TODO:本人確認が必要

	for i, favorite := range inmemory.Favorites {
		if favorite.UserID == userID && favorite.StudySetID == studySetID {
			inmemory.Favorites = utils.RemoveElementFromSlice(inmemory.Favorites, i)
			return nil
		}
	}
	return errors.New("favorite not found")
}

func (r *FavoriteRepository) GetFavoritesByUserID(userID string) ([]*model.Favorite, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var userFavorites []*model.Favorite
	for _, favorite := range inmemory.Favorites {
		if favorite.UserID == userID {
			userFavorites = append(userFavorites, favorite)
		}
	}
	return userFavorites, nil
}

func (r *FavoriteRepository) IsFavorite(userID, studySetID string) (bool, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, favorite := range inmemory.Favorites {
		if favorite.UserID == userID && favorite.StudySetID == studySetID {
			return true, nil
		}
	}
	return false, nil
}

func (r *FavoriteRepository) GetFavoriteStudySetsByUserID(userID string) ([]*model.StudySet, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var studySets []*model.StudySet
	// 2重ループになっていて計算量は良くないけど、テストケースは多くないから問題ない
	for _, favorite := range inmemory.Favorites {
		if favorite.UserID == userID {
			for _, studySet := range inmemory.StudySets {
				if studySet.ID == favorite.StudySetID {
					studySets = append(studySets, studySet)
				}
			}
		}
	}
	return studySets, nil
}
