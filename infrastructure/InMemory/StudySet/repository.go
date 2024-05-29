package inmemory

import (
	"errors"
	"go-training/domain/model"
	"go-training/domain/repository"
	"sync"
)

type StudySetRepository struct {
	mu        sync.Mutex
	studySets map[string]*model.StudySet
}

func NewStudySetRepository() repository.StudySetRepository {
	return &StudySetRepository{
		studySets: make(map[string]*model.StudySet),
	}
}

func (r *StudySetRepository) Create(studySet *model.StudySet) error {
	r.mu.Lock()
	defer r.mu.Unlock()

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
