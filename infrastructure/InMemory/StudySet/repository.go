package studySet

import (
	"errors"
	"go-training/domain/model"
	"go-training/domain/repository"
	inmemory "go-training/infrastructure/InMemory"
	"go-training/utils"
	"strings"
	"sync"

	"github.com/google/uuid"
)

type StudySetRepository struct {
	mu        sync.Mutex
	studySets map[string]*model.StudySet
}

func NewStudySetRepository() repository.StudySetRepository {
	repo := &StudySetRepository{
		studySets: make(map[string]*model.StudySet),
	}
	for _, studySet := range inmemory.StudySets {
		repo.studySets[studySet.ID] = studySet
	}
	return repo
}

func (r *StudySetRepository) Create(studySet *model.StudySet) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	// TODO:本人確認が必要

	// 外部キーのチェック: UserIDが存在するか
	isUserExists := false
	for _, user := range inmemory.Users {
		if user.ID == studySet.UserID {
			isUserExists = true
		}
	}
	if !isUserExists {
		return errors.New("user doesn't exist")
	}

	// uuid作成
	studySet.ID = uuid.New().String()

	inmemory.StudySets = append(inmemory.StudySets, studySet)
	return nil
}

func (r *StudySetRepository) GetByID(id string) (*model.StudySet, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, studySet := range inmemory.StudySets {
		if studySet.ID == id {
			return studySet, nil
		}
	}

	return nil, errors.New("study set not found")
}

func (r *StudySetRepository) GetByUserID(userID string) ([]*model.StudySet, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var userStudySets []*model.StudySet

	for _, studySet := range inmemory.StudySets {
		if studySet.UserID == userID {
			userStudySets = append(userStudySets, studySet)
		}
	}

	return userStudySets, nil
}

func (r *StudySetRepository) Update(studySet *model.StudySet) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// TODO:本人確認が必要
	for i, studySetFromDB := range inmemory.StudySets {
		if studySetFromDB.ID == studySet.ID {
			inmemory.StudySets[i] = studySet
			return nil
		}
	}

	return errors.New("study set not found")

}

func (r *StudySetRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	// TODO:本人確認が必要

	for i, studySetFromDB := range inmemory.StudySets {
		if studySetFromDB.ID == id {
			inmemory.StudySets = utils.RemoveElementFromSlice(inmemory.StudySets, i)
			return nil
		}
	}

	return errors.New("study set not found")

}

func (r *StudySetRepository) SearchByTitle(title string) ([]*model.StudySet, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var results []*model.StudySet
	for _, studySet := range inmemory.StudySets {
		if strings.Contains(strings.ToLower(studySet.Title), strings.ToLower(title)) {
			results = append(results, studySet)
		}
	}
	return results, nil
}
