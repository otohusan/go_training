package inmemory

import (
	"errors"
	"go-training/domain/model"
	"go-training/domain/repository"
	inmemory "go-training/infrastructure/InMemory"
	"sync"
	"time"

	"github.com/google/uuid"
)

type FavoriteRepository struct {
	mu        sync.Mutex
	favorites map[string]*model.Favorite
}

func NewFavoriteRepository() repository.FavoriteRepository {
	repo := &FavoriteRepository{
		favorites: make(map[string]*model.Favorite),
	}

	for _, favorite := range inmemory.InitializeFavorites() {
		repo.favorites[favorite.ID] = favorite
	}

	return repo
}

func (r *FavoriteRepository) AddFavorite(userID, studySetID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// 新しいIDをUUIDで生成
	tableID := uuid.New().String()
	newFavorite := &model.Favorite{
		ID:         tableID,
		UserID:     userID,
		StudySetID: studySetID,
		CreatedAt:  time.Now(),
	}
	r.favorites[newFavorite.ID] = newFavorite
	return nil
}

func (r *FavoriteRepository) RemoveFavorite(userID, studySetID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for id, favorite := range r.favorites {
		if favorite.UserID == userID && favorite.StudySetID == studySetID {
			delete(r.favorites, id)
			return nil
		}
	}
	return errors.New("favorite not found")
}

func (r *FavoriteRepository) GetFavoritesByUserID(userID string) ([]*model.Favorite, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var userFavorites []*model.Favorite
	for _, favorite := range r.favorites {
		if favorite.UserID == userID {
			userFavorites = append(userFavorites, favorite)
		}
	}
	return userFavorites, nil
}

func (r *FavoriteRepository) IsFavorite(userID, studySetID string) (bool, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, favorite := range r.favorites {
		if favorite.UserID == userID && favorite.StudySetID == studySetID {
			return true, nil
		}
	}
	return false, nil
}
