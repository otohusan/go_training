package studySet

import (
	"errors"
	"fmt"
	"go-training/domain/model"
	"go-training/domain/repository"
	inmemory "go-training/infrastructure/InMemory"
	"sync"
)

type StudySetRepository struct {
	mu        sync.Mutex
	studySets map[string]*model.StudySet
}

func NewStudySetRepository() repository.StudySetRepository {
	repo := &StudySetRepository{
		studySets: make(map[string]*model.StudySet),
	}
	for _, studySet := range inmemory.InitializeStudySets() {
		repo.studySets[studySet.ID] = studySet
	}
	return repo
}

func (r *StudySetRepository) Create(studySet *model.StudySet) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// 学習セットの数を基に新しいIDを生成 TODO:削除とかあったら、id被る
	studySet.ID = fmt.Sprintf("%d", len(r.studySets)+1)

	if _, exists := r.studySets[studySet.ID]; exists {
		return errors.New("study set already exists")
	}

	r.studySets[studySet.ID] = studySet
	return nil
}

func (r *StudySetRepository) GetByID(id string) (*model.StudySet, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	studySet, exists := r.studySets[id]
	if !exists {
		return nil, errors.New("study set not found")
	}

	return studySet, nil
}

func (r *StudySetRepository) GetByUserID(userID string) ([]*model.StudySet, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var userStudySets []*model.StudySet
	for _, studySet := range r.studySets {
		if studySet.UserID == userID {
			userStudySets = append(userStudySets, studySet)
		}
	}
	return userStudySets, nil
}

func (r *StudySetRepository) Update(studySet *model.StudySet) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.studySets[studySet.ID]; !exists {
		return errors.New("study set not found")
	}

	r.studySets[studySet.ID] = studySet
	return nil
}

func (r *StudySetRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.studySets[id]; !exists {
		return errors.New("study set not found")
	}

	delete(r.studySets, id)
	return nil
}
